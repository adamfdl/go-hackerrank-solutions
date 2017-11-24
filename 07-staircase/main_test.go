package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestPrintStars(t *testing.T) {
	data_test, err := ioutil.ReadFile("sample.dat")
	if err != nil {
		log.Fatal(err)
	}

	result := PrintStairs(3)

	if string(data_test) != result {
		t.Fatalf("Test failed")
	}
}
