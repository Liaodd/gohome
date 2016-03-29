package http

import (
	"testing"
)

func TestRequest(t *testing.T) {
	//var params map[string] string
	params := make(map[string] string)
	params["test"] = "google"
	request := NewRequest("http://www.google.com", "get", params)

	t.Logf(request.url)
}

