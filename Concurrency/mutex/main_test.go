package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(2)

	go updateMsg("x")

	go updateMsg("Goodbye, cruel world!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {

		t.Error("incorrect value in msg")

	}
}
