package main

import (
	"clockface/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
