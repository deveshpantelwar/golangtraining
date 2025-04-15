package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func test_printSomething(t *testing.T) {

	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go printSomething("gamma", &wg)
	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "gamma") {
		t.Errorf("Expected to finf gamma but not present ")
	}
}
