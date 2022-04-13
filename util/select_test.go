package util

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// TestSimpleSelection aa
func TestSimpleSelection(t *testing.T) {
	rescue := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	title := "Question"
	options := []string{"Option A", "Option B"}
	SimpleSelection(title, options)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescue

	output := string(out)

	// Fail the test if any option wasn't present
	for _, v := range options {
		if !strings.Contains(output, v) {
			t.Fail()
		}
	}

	// Check title
	if !strings.Contains(output, title) {
		t.Fail()
	}
}
