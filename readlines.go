package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var s = `hello` + "\r" + `
abcdeghijklmnopqrstuvwxyz
0123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789`

func main() {
	b := bufio.NewReaderSize(strings.NewReader(s), 16)
	err := ReadLines(b, 60, func(line []byte) error {
		fmt.Printf("%q\n", line)
		return nil
	})
	fmt.Printf("error: %v\n", err)
}

// ReadLines reads lines from r and calls fn with each read line, not
// including the line terminator.  If a line exceeds the given maximum
// size, it will be truncated and the rest of the line discarded.  If fn
// returns a non-nil error, reading ends and the error is returned from
// ReadLines.  When EOF is encountered, ReadLines returns nil.
func ReadLines(r io.Reader, maxSize int, fn func(line []byte) error) error {
	b := bufio.NewReader(r)
	for {
		line, isPrefix, err := b.ReadLine()
		if err != nil {
			return eofNilError(err)
		}
		if !isPrefix {
			// Simple line that fits within the bufio buffer size.
			if len(line) > maxSize {
				line = line[0:maxSize]
			}
			if err := fn(line); err != nil {
				return err
			}
			continue
		}
		buf := make([]byte, len(line), len(line)*2)
		copy(buf, line)
		for isPrefix {
			line, isPrefix, err = b.ReadLine()
			if err != nil {
				if err := fn(buf); err != nil {
					return err
				}
				return eofNilError(err)
			}
			buf = append(buf, line...)
			if maxSize > 0 && len(buf) >= maxSize {
				break
			}
		}
		if len(line) > maxSize {
			line = line[0:maxSize]
		}
		if err := fn(line); err != nil {
			return err
		}
		// Discard any of the line that exceeds the maximum size
		for isPrefix {
			_, isPrefix, err = b.ReadLine()
			if err != nil {
				return eofNilError(err)
			}
		}
	}
	panic("unreachable")
}

func eofNilError(err error) error {
	if err == io.EOF {
		return nil
	}
	return err
}
