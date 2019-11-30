package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pauloavelar/binarydist"
)

type customWriter struct {
	io.Writer
	total int64
}

func (w *customWriter) Write(b []byte) (int, error) {
	w.total += int64(len(b))
	return w.Writer.Write(b)
}

func main() {
	args := os.Args[1:]

	fmt.Printf("Calculating diff...\n- OLD: %s\n- NEW: %s\n", args[0], args[1])
	old, _ := os.Open(args[0])
	new, _ := os.Open(args[1])
	tmp, _ := os.Create(args[2])
	out := &customWriter{Writer: tmp}

	defer old.Close()
	defer new.Close()
	defer tmp.Close()

	binarydist.Diff(old, new, out)

	fileSize := float64(out.total) / 1024
	fmt.Printf("Diff size: %.1f KB\n", fileSize)
}
