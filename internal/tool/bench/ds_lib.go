// Copyright 2015, Joe Tsai. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

// +build !no_ds_lib

package bench

import (
	"io"

	"github.com/dsnet/compress/brotli"
	"github.com/dsnet/compress/bzip2"
	"github.com/dsnet/compress/flate"
)

func init() {
	RegisterDecoder(FormatBrotli, "ds",
		func(r io.Reader) io.ReadCloser {
			rd, err := brotli.NewReader(r, nil)
			if err != nil {
				panic(err)
			}
			return rd
		})
	RegisterDecoder(FormatFlate, "ds",
		func(r io.Reader) io.ReadCloser {
			rd, err := flate.NewReader(r, nil)
			if err != nil {
				panic(err)
			}
			return rd
		})
	RegisterEncoder(FormatBZ2, "ds",
		func(w io.Writer, lvl int) io.WriteCloser {
			wr, err := bzip2.NewWriter(w, &bzip2.WriterConfig{Level: lvl})
			if err != nil {
				panic(err)
			}
			return wr
		})
	RegisterDecoder(FormatBZ2, "ds",
		func(r io.Reader) io.ReadCloser {
			rd, err := bzip2.NewReader(r, nil)
			if err != nil {
				panic(err)
			}
			return rd
		})
}
