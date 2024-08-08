package main

import (
	"bytes"
	"context"
	"fmt"
	"regexp"

	"github.com/NDoolan360/etyml-app/web/templates"
	"github.com/a-h/templ"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type OutboundPayload struct {
	Guesses []string `json:"guesses"`
	Tree    Tree     `json:"tree"`
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	re := regexp.MustCompile(`\/puzzle\/(?P<Id>.+)`)
	matches := re.FindStringSubmatch(request.Path)
	idIndex := re.SubexpIndex("Id")
	if idIndex > len(matches) {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error 400: No Puzzle id provided",
		}, nil
	}
	puzzleId := matches[idIndex]

	guesses, ok := request.MultiValueQueryStringParameters["guess"]
	if !ok {
		guesses = []string{}
	}
	ctx = context.WithValue(ctx, "guesses", guesses)

	hints, ok := request.MultiValueQueryStringParameters["hint"]
	if !ok {
		hints = []string{}
	}

	ctx = context.WithValue(ctx, "hints", hints)

	puzzle, ok := etymologyTrees[puzzleId]
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       fmt.Sprintf("Error 404: Invalid puzzle: %s not found", puzzleId),
		}, nil
	}

	obscurer := '_'
	obscuredPuzzle := puzzle.obscure(guesses, hints, obscurer)

	var template templ.Component
	if request.Headers["hx-request"] == "true" {
		fmt.Println(request)
		template = templates.UpdatePuzzle(
			obscuredPuzzle.html(hints),
			obscuredPuzzle.isComplete(obscurer),
			request.Headers["etyml-hint"] == "true",
		)
	} else {
		template = templates.BaseLayout(
			templates.Puzzle(
				obscuredPuzzle.html(hints),
				obscuredPuzzle.isComplete(obscurer),
			),
		)
	}

	buf := bytes.NewBuffer([]byte{})
	renderErr := template.Render(ctx, buf)
	if renderErr != nil {
		return nil, renderErr
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":  "text/html",
			"Cache-Control": "no-cache, no-store, must-revalidate",
		},
		Body:            buf.String(),
		IsBase64Encoded: false,
	}, nil
}
