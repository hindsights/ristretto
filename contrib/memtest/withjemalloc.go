//go:build jemalloc
// +build jemalloc

package main

import (
	"github.com/dgraph-io/ristretto/z"
	"github.com/hindsights/gslog"
)

func Calloc(size int) []byte { return z.Calloc(size, "memtest") }
func Free(bs []byte)         { z.Free(bs) }
func NumAllocBytes() int64   { return z.NumAllocBytes() }

func check() {
	if buf := z.CallocNoRef(1, "memtest"); len(buf) == 0 {
		gslog.Fatalf("Not using manual memory management. Compile with jemalloc.")
	} else {
		z.Free(buf)
	}

	z.StatsPrint()
}

func init() {
	gslog.Infof("USING JEMALLOC")
}
