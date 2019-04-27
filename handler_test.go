package main

import (
	"context"
	"testing"
)

func TestAbs(t *testing.T) {
	_, err := HandleRequest(context.TODO())
	if err != nil {
		t.Error(err)
	}
}
