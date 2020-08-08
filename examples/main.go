package main

import (
	"fmt"
	"time"

	"github.com/usk81/anytime"
)

func main() {
	loc, err := time.LoadLocation("PDT")
	if err != nil {
		loc = time.FixedZone("PDT", -7*60*60)
	}
	time.Local = loc

	now := anytime.Now()
	fmt.Println(now)
}
