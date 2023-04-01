package server

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"kratos/internal/errors"
	"testing"
)

func TestHTTPStruct(t *testing.T) {
	a := &errors.HTTPError{Errors: map[string][]string{
		"body": {"hehe"},
	}}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	fmt.Println(string(b))
}
