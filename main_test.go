package main_test

import (
	"net/http"
	"testing"
)

func TestIndex(t *testing.T) {
	http.Get("http://localhost:8080/test/test")
}
