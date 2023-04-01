package biz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	fmt.Println(hashPassword("abc"))
}

func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)

	a.True(verifyPassword("a", "a"))
}