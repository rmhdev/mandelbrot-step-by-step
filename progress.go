package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Progress  responsible for writing the progress of the app
type Progress struct {
	out     io.Writer
	maxBars int
}

// CreateProgress new progress object
func CreateProgress(verbosity string) (p Progress, err error) {
	switch verbosity {
	case "v":
		return Progress{os.Stdout, 20}, nil
	case "quiet":
		return Progress{&bytes.Buffer{}, 20}, nil
	}
	return Progress{&bytes.Buffer{}, 20}, fmt.Errorf("Undefined verbosity '%s'", verbosity)
}

func (p Progress) write(a interface{}) {
	fmt.Fprint(p.out, a)
}

func (p Progress) writeln(a interface{}) {
	fmt.Fprintln(p.out, a)
}

func (p Progress) bar(a interface{}, total int, portion int) {
	p.write("\r")
	p.write(a)
	p.write("  [")
	bars := 0
	if portion > 0 {
		bars = int(float32(p.maxBars) / (float32(total) / float32(portion)))
	}
	spaces := p.maxBars - bars - 1
	for i := 0; i < bars; i++ {
		p.write("=")
	}
	if bars < p.maxBars {
		p.write(">")
	}
	for i := 0; i < spaces; i++ {
		p.write(" ")
	}
	p.write("] ")
}
