// main_test.go
package main

import (
	"testing"
	flatbuffers "github.com/google/flatbuffers/go"
)

func BenchmarkWrite(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	b.ReportAllocs()
	for i := 0; i < b.N ; i++ {
		builder.Reset()
		MakeUser(builder, []byte("Arthur Dent"), 42)
	}
}

func BenchmarkRead(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	buf := MakeUser(builder, []byte("Arthur Dent"), 42)
	b.ReportAllocs()
	for i := 0; i < b.N ; i++ {
		ReadUser(buf)
	}
}

func BenchmarkRoundtrip(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	b.ReportAllocs()
	for i := 0; i < b.N ; i++ {
		builder.Reset()
		ReadUser(MakeUser(builder, []byte("Arthur Dent"), 42))
	}
}

