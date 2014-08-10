package connpass

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	file := "sample.json"
	jsonBlob, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("Failed to open %v. Reason:%s.", file, err)
	}
	result, err := parse(jsonBlob)
	if err != nil {
		t.Fatalf("Failed to parse the JSON. Reason:%s.", err)
	}

	if result.Returned != 10 {
		t.Errorf("Unexpected number of returned events: %v.", result.Returned)
	}
}
