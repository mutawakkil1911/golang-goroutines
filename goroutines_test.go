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

func TampilkanNomer(nomer int) {
	fmt.Println("Nomer ke ", nomer)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go TampilkanNomer(i)
	}

	time.Sleep(5 * time.Second)
}
