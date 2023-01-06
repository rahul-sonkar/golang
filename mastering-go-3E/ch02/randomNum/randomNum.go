package main

import (
	"fmt"
	"math/rand"

	"time"
)

func main() {
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	buf := make([]byte, 5)
	for i := 0; i < 5; i++ {
		tmp := (rand.Intn(91-65)) + 65
		buf[i] = byte(tmp)
	}

	fmt.Println(string(buf))
}
