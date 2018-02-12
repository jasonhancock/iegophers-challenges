package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

var (
	applicationVersion = "1.0"
)

func main() {

	var (
		countsByte      bool
		countsCharacter bool
		countsWords     bool
		countsLines     bool
		maxLineLength   bool
		version         bool
	)

	flag.BoolVar(&countsByte, "bytes", false, "print the byte counts")
	// flag.BoolVar(&countsCharacter, "chars", false, "print the character count") // Ran out of time to implement
	flag.BoolVar(&countsLines, "lines", false, "print the newline counts")
	// flag.BoolVar(&maxLineLength, "max-line-length", false, "print the length of the longest line") // Ran out of time to implement
	flag.BoolVar(&countsWords, "words", false, "print the word counts")
	flag.BoolVar(&version, "version", false, "outupt version information and exit")
	flag.Parse()

	if version {
		fmt.Printf("version %s\n", applicationVersion)
		return
	}

	var reader io.Reader

	if flag.NArg() == 0 {
		reader = os.Stdin
	} else if flag.NArg() > 1 {
		log.Fatal("more than one file passed")
	} else {
		var err error
		reader, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	}

	var cmd command

	if countsByte {
		cmd = countBytesCmd{}
	}
	if cmd == nil && countsLines {
		cmd = countLinesCmd{}
	}
	if cmd == nil && countsWords {
		cmd = countWordsCmd{}
	}
	if cmd == nil {
		log.Fatal("no command")
	}

	count, err := cmd.Run(reader)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", count)
}

type command interface {
	Run(r io.Reader) (int64, error)
}

type countBytesCmd struct{}

func (c countBytesCmd) Run(r io.Reader) (int64, error) {
	var count int64
	b := make([]byte, 5)
	for {
		n1, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return count, err
		}
		count += int64(n1)
	}

	return count, nil
}

type countLinesCmd struct{}

func (c countLinesCmd) Run(r io.Reader) (int64, error) {
	var count int64
	b := make([]byte, 5)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return count, err
		}
		for i := range b {
			if b[i] == '\n' {
				count++
			}
		}
	}

	return count, nil
}

type countWordsCmd struct{}

func (c countWordsCmd) Run(r io.Reader) (int64, error) {
	var count int64
	b := make([]byte, 5)
	var lastChar = true
	for {
		n, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return count, err
		}

		for i := 0; i < n; i++ {
			thisChar := unicode.IsSpace(rune(b[i]))
			if lastChar && !thisChar {
				count++
			}
			lastChar = thisChar
		}
	}

	return count, nil
}
