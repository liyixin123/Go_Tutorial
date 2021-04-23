package greetings

import (
	"regexp"
	"testing"
)

//TestHelloName calls greetings.Hello with a name, checking for a valid return value
func TestHelloName(t *testing.T) {
	name := "Liyixin"
	msg, err := Hello(name)
	want := regexp.MustCompile(`\b` + name + `\b`)
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

//TestHelloEmpty calls greetings.Hello with an empty string,
//cheking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "",error  `, msg, err)
	}

}
