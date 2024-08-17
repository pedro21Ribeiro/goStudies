package greetings

import (
	"testing"
	"regexp"
)

func testeHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello(name)
	if!want.MatchString(name) || err!=nil{
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}

func testeHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg!= "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}