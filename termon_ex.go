package main

import (
	"fmt"
	"github.com/kless/Go-Inline/inline"
	"log"
	"os"
	"path"
	"syscall"
)

var tempHistory = path.Join(os.TempDir(), "go_inline")

func init() {
	inline.Input = os.Stderr
	inline.InputFd = syscall.Stderr
}

func main() {
	fmt.Println("\n== Reading line")
	fmt.Printf("Press ^D to exit\n\n")

	hist, err := inline.NewHistory(tempHistory)
	if err != nil {
		log.Print(err)
	}
	hist.Load()

	ln, err := inline.NewLineByDefault(hist)
	if err != nil {
		log.Print(err)
	}
	defer ln.Restore()

	for {
		if _, err = ln.Read(); err != nil {
			if err == inline.ErrCtrlD {
				hist.Save()
			} else {
				fmt.Fprintf(os.Stderr, "%s", err)
			}

			break
		}
	}

	os.Remove(tempHistory)
}
