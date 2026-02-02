package main

import (
	"fmt"
	"time"

	"github.com/rotorapp/snowflake"
)

func main() {
	snowflake.Init(1)

	for i := range 5 {
		sf := snowflake.New()
		fmt.Println(sf, sf.Time(), sf.GeneratorID(), sf.Increment())

		if i == 4 {
			break
		}

		time.Sleep(1 * time.Second)
	}
}
