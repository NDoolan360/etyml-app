package main

import (
	"bytes"
	"context"
	"fmt"
	"regexp"

	"github.com/NDoolan360/etyml-app/web/templates"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type OutboundPayload struct {
	Guesses []string `json:"guesses"`
	Tree    Tree     `json:"tree"`
}

func main() {
	lambda.Start(handler(etymologyTrees))
}

// curried handler to inject trees used in validation and response
func handler(trees map[string]Tree) func(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
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

		puzzle, ok := trees[puzzleId]
		if !ok {
			return &events.APIGatewayProxyResponse{
				StatusCode: 404,
				Body:       fmt.Sprintf("Error 404: Invalid puzzle: %s not found", puzzleId),
			}, nil
		}

		obscuredPuzzle := puzzle.obscure(guesses, hints, '_')
		html := obscuredPuzzle.html(hints)
		complete := obscuredPuzzle.isComplete('_')

		buf := bytes.NewBuffer([]byte{})
		var renderErr error

		if request.Headers["hx-request"] == "true" {
			renderErr = templates.PuzzleUpdate(html, complete).Render(ctx, buf)
			if request.Headers["etyml-hint"] == "true" {
				renderErr = templates.HintUpdate().Render(ctx, buf)
			} else if request.Headers["etyml-guess"] == "true" {
				renderErr = templates.GuessUpdate().Render(ctx, buf)
			}
		} else {
			renderErr = templates.BaseLayout(
				templates.Puzzle(html, complete),
				map[string]any{"hx-boost": "true", "hx-swap": "none", "hx-replace-url": "true"},
			).Render(ctx, buf)
		}

		if renderErr != nil {
			return nil, renderErr
		}

		return &events.APIGatewayProxyResponse{
			StatusCode:      200,
			Headers:         map[string]string{"Content-Type": "text/html"},
			Body:            buf.String(),
			IsBase64Encoded: false,
		}, nil
	}
}
