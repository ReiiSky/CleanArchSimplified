package application_test

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

var client = resty.New()
var endpointURL = "http://localhost:6007"

func TestInsertUser(t *testing.T) {
	client.
		R()
}
