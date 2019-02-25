package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var sample1 string
var sample2 string

func BenchmarkSprintf(b *testing.B) {
	s := ""
	sample1 = "Hello"
	sample2 = "World"
	o := sampleStruct{
		Str:    sample1,
		StrPtr: &sample2,
		N:      0,
	}
	for n := 0; n < b.N; n++ {
		o.N = n
		j, _ := json.Marshal(&o)
		s = fmt.Sprintf("%s %s this is iteration %d and the struct is %s", sample1, sample2, n, string(j))
	}
	s += ""
}

func BenchmarkPlusConcat(b *testing.B) {
	s := ""
	sample1 = "Hello"
	sample2 = "World"
	o := sampleStruct{
		Str:    sample1,
		StrPtr: &sample2,
		N:      0,
	}
	for n := 0; n < b.N; n++ {
		o.N = n
		j, _ := json.Marshal(&o)
		s = sample1 + " " + sample2 + " this is iteration " + strconv.Itoa(n) + " and the struct is " + string(j)
	}
	s += ""
}

func BenchmarkStringBuilder(b *testing.B) {
	s := ""
	sample1 = "Hello"
	sample2 = "World"
	o := sampleStruct{
		Str:    sample1,
		StrPtr: &sample2,
		N:      0,
	}
	for n := 0; n < b.N; n++ {
		o.N = n
		j, _ := json.Marshal(&o)
		sb := strings.Builder{}
		sb.WriteString(sample1)
		sb.WriteString(" ")
		sb.WriteString(sample2)
		sb.WriteString(" this is iteration ")
		sb.WriteString(strconv.Itoa(n))
		sb.WriteString(" and the struct is ")
		sb.Write(j)
		s = sb.String()
	}
	s += ""
}

func BenchmarkBytesBuffer(b *testing.B) {
	s := ""
	sample1 = "Hello"
	sample2 = "World"
	o := sampleStruct{
		Str:    sample1,
		StrPtr: &sample2,
		N:      0,
	}
	for n := 0; n < b.N; n++ {
		o.N = n
		j, _ := json.Marshal(&o)
		buf := bytes.Buffer{}
		buf.WriteString(sample1)
		buf.WriteString(" ")
		buf.WriteString(sample2)
		buf.WriteString(" this is iteration ")
		buf.WriteString(strconv.Itoa(n))
		buf.WriteString(" and the struct is ")
		buf.Write(j)
		s = buf.String()
	}
	s += ""
}
