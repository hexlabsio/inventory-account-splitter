package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestGet(t *testing.T) {
	response, err := HandleRequest(events.APIGatewayProxyRequest{HTTPMethod: "GET"})
	if err != nil {
		t.Error(err)
	}
	assertEqual(t, response.StatusCode, 200)
}

func TestPost(t *testing.T) {
	response, err := HandleRequest(events.APIGatewayProxyRequest{HTTPMethod: "POST"})
	if err != nil {
		t.Error(err)
	}
	assertEqual(t, response.StatusCode, 202)
}
