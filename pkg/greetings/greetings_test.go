package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Marty"
	expected := regexp.MustCompile(`\b` + name + `\b`)
	actual, err := Hello(name)
	if !expected.MatchString(actual) || err != nil {
		t.Fatalf(`Hello("Marty") = %q, %v, want match for %#q, nil`, actual, err, expected)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
