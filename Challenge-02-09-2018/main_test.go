package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/cheekybits/is"
)

func TestStuff(t *testing.T) {

	var tests = []struct {
		file  string
		bytes int64
		lines int64
		words int64
	}{
		{"testdata/foo.txt", 32, 6, 7},
	}

	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			isf := is.New(t)
			data, err := ioutil.ReadFile(tt.file)
			isf.NoErr(err)
			t.Run("bytes", func(t *testing.T) {
				is := is.New(t)
				cmd := countBytesCmd{}
				n, err := cmd.Run(bytes.NewReader(data))

				is.NoErr(err)
				is.Equal(n, tt.bytes)
			})

			t.Run("lines", func(t *testing.T) {
				is := is.New(t)
				cmd := countLinesCmd{}
				n, err := cmd.Run(bytes.NewReader(data))

				is.NoErr(err)
				is.Equal(n, tt.lines)
			})

			t.Run("words", func(t *testing.T) {
				is := is.New(t)
				cmd := countWordsCmd{}
				n, err := cmd.Run(bytes.NewReader(data))

				is.NoErr(err)
				is.Equal(n, tt.words)
			})
		})
	}

}
