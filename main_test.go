// main_test.go
package main

import (
	"bytes"
	"testing"
	flatbuffers "github.com/google/flatbuffers/go"
)

func BenchmarkWrite(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	b.ReportAllocs()
	for i := 0; i < b.N ; i++ {
		builder.Reset()
		buf := MakeUser(builder, []byte("Arthur Dent"), 42)
		if i == 0 {
			b.SetBytes(int64(len(buf)))
		}
	}
}

func BenchmarkRead(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	name := []byte("Arthur Dent")
	buf := MakeUser(builder, name, 42)
	b.SetBytes(int64(len(buf)))
	b.ReportAllocs()
	for i := 0; i < b.N ; i++ {
		got_name, _ := ReadUser(buf)
		// do some work to prevent cheating the benchmark:
		bytes.Equal(got_name, name)
	}
}

func BenchmarkRoundtrip(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	b.ReportAllocs()
	for i := 0; i < b.N ; i++ {
		builder.Reset()
		buf := MakeUser(builder, []byte("Arthur Dent"), 42)
		got_name, _ := ReadUser(buf)
		if i == 0 {
			b.SetBytes(int64(len(buf)))
		}
		// do some work to prevent cheating the benchmark:
		bytes.Equal(got_name, []byte("Arthur Dent"))
	}
}
