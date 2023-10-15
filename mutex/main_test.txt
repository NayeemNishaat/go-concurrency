package main

import "testing"

func TestUnsafeUpdateMessage(t *testing.T) {
	msg = "Sweet Dream"

	wg.Add(2)
	go UnsafeUpdateMessage("Good bye!")
	go UnsafeUpdateMessage("Bye!")
	wg.Wait()

	if msg != "Good bye!" {
		t.Errorf("Incorrect value in msg.")
	}
}

// go test -race . // Remark: Always test for race conditions
