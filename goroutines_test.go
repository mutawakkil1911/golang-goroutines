package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunSalam() {
	fmt.Println("Assalamualaikum")
}

func TestCreateGoroutines(t *testing.T) {
	go RunSalam()
	fmt.Println("Ops")

	time.Sleep(1 * time.Second)
}
