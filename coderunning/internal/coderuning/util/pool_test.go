package util

import (
	"context"
	"testing"
	"time"
)

func TestAddBot(t *testing.T) {
	ctx := context.Background()
	go Run(ctx)

	for i := 0; i < 2; i++ {
		AddBot(int32(i), `package main

import "fmt"

func main() {
	fmt.Println("1")
}`, "sdsadas")
	}

	time.Sleep(4 * time.Second)
}
