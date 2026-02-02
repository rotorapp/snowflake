package main

import (
	"fmt"
	"time"

	"github.com/rotorapp/snowflake"
)

func main() {
	for i := range 5 {
		sf := snowflake.New()
		fmt.Println(sf)

		if i == 4 {
			break
		}

		time.Sleep(1 * time.Second)
	}
}
