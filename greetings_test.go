package greetings

import (
	"testing"
	"regexp"
)

//Test the HelloName calls greetings.Hello with a name,
//checking for a valid return value

func TestHelloName(t *testing.T) {
	name := "Dyes"
	want := regexp.MustCompile(`\b` +name+ `\b`)
	msg, err := Hello("Dyes")
	if !want.MatchString(msg)  || err != nil {
		t.Fatalf(`Hello("Dyes") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

//TestHeloEmpty calls greetings.Hello with empty string,
//checking for an error .

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
if msg != "" || err == nil {
	t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
}
}