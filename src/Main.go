package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		now := time.Now()
		isoTime := now.Format(time.RFC3339)

		fmt.Printf("[%s] Heartbeat\n", isoTime)

		time.Sleep(10 * time.Second)
	}
}
