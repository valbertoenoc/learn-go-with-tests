package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(time.Second)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
