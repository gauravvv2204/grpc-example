package main

import (
	"regexp"
	"testing"
)

func TestSendRequest(t *testing.T) {
	key := "hello"
	val := "world"
	want := regexp.MustCompile(`\b` + val + `\b`)
	msg := sendRequest(key, val)
	if !want.MatchString(msg.GetVal()) {
		t.Fatalf(`sendRequest("%q") = %q, want match for %#q, nil`, key, msg, want)
	}
}
