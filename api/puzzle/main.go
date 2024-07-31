package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/NDoolan360/etyml-app/web/templates"
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
	puzzleId, ok := request.QueryStringParameters["id"]
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Error 422: No Puzzle id provided",
		}, nil
	}

	guesses, ok := request.MultiValueQueryStringParameters["guess"]
	if !ok {
		guesses = []string{}
	}

	puzzle, ok := etymologyTrees[puzzleId]
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       fmt.Sprintf("Error 404: Puzzle %s not found", puzzleId),
		}, nil
	}

	obscuredPuzzle := puzzle.obscure(guesses)

	buf := bytes.NewBuffer([]byte{})
	renderErr := templates.BaseLayout(
		templates.Puzzle(
			puzzleId,
			guesses,
			obscuredPuzzle.html(),
			obscuredPuzzle.isComplete(),
		),
	).Render(context.Background(), buf)
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
