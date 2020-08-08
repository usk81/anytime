package anytime

import (
	"fmt"
	"time"
)

func ExampleNow() {
	now := Now()
	fmt.Println(now)
	// 2020-08-08 16:35:51.934459 +0900 JST m=+0.000123825
}

func ExampleNow_set_local() {
	loc, err := time.LoadLocation("PDT")
	if err != nil {
		loc = time.FixedZone("PDT", -7*60*60)
	}
	time.Local = loc

	now := Now()
	fmt.Println(now)
	// 2020-08-08 00:42:09.267238 -0700 PDT m=+0.002352520
}
