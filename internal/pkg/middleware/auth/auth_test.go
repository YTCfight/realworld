package auth

import (
	"fmt"
	"testing"
)


func TestGenerateToken(t *testing.T) {
	tk := GenerateToken("hello", "2233")
	fmt.Println(tk)
}

