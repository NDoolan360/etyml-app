package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

var testTrees = map[string]Tree{
	"term": testTree,
}

// returns HTTP status code 200 in the response body for a valid id
func TestHandler200(t *testing.T) {
	ctx := context.Background()
	response, handlerErr := handler(testTrees)(ctx, events.APIGatewayProxyRequest{
		Path: "/puzzle/term",
	})
	if handlerErr != nil {
		t.Errorf("expected no error but got %s", handlerErr.Error())
	}

	if response.StatusCode != 200 {
		t.Errorf("expected status code 200 but got %d", response.StatusCode)
	}

	expected := `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link rel="icon" href="/favicon.ico"><link rel="stylesheet" type="text/css" href="/styles.css"><script src="/scripts/htmx@2.0.1-custom-min.js"></script></head><body hx-boost="true" hx-replace-url="true" hx-swap="none"><div id="tree"><ul><li id="jnvoiuhef"><span><h1>Lang</h1><h2>Term</h2><p>Definition</p></span> <ul><li id="noiusdnml"><span><h1>Lang</h1><h2>*_____</h2></span> </li><li id="lkjdafohc"><span><h1>Lang</h1><h2>c_______</h2></span> </li></ul></li></ul></div><div id="controls"><form id="input" hx-headers="{&#34;etyml-guess&#34;: &#34;true&#34;}"><input id="textbox" required autofocus name="guess"></form></div><div id="previous"><h2>Guesses</h2><ul id="guesses"></ul></div></body></html>`
	if response.Body != expected {
		t.Errorf("expected '%s' but got '%s'", expected, response.Body)
	}
}

// returns HTTP status code 400 in the response body when no id is given
func TestHandler400(t *testing.T) {
	ctx := context.Background()
	response, handlerErr := handler(testTrees)(ctx, events.APIGatewayProxyRequest{
		Path: "/puzzle",
	})
	if handlerErr != nil {
		t.Errorf("expected no error but got %s", handlerErr.Error())
	}

	if response.StatusCode != 400 {
		t.Errorf("expected status code 400 but got %d", response.StatusCode)
	}
}

// returns HTTP status code 404 in the response body when the id is not in the trees
func TestHandler404(t *testing.T) {
	ctx := context.Background()
	response, handlerErr := handler(testTrees)(ctx, events.APIGatewayProxyRequest{
		Path: "/puzzle/not_there",
	})
	if handlerErr != nil {
		t.Errorf("expected no error but got %s", handlerErr.Error())
	}

	if response.StatusCode != 404 {
		t.Errorf("expected status code 404 but got %d", response.StatusCode)
	}
}
