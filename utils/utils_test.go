package utils

import (
	"reflect"
	"testing"
)

func TestFromJson(t *testing.T) {

	type testStruct struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	j := `
{
	"a": 1,
	"b": 2
}
`
	got, err := FromJson[testStruct](j)

	if err != nil {
		t.Error(err.Error())
	}

	want := testStruct{}
	want.A = 1
	want.B = 2

	if !reflect.DeepEqual(*got, want) {
		t.Error("received value differs from value which was expected to get")
	}
}

func TestDecodeFromJson(t *testing.T) {

	type testStruct struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	j := `
{
	"a": 1,
	"b": 2
}
`
	got, err := DecodeFromJson[testStruct](j)

	if err != nil {
		t.Error(err.Error())
	}

	want := testStruct{}
	want.A = 1
	want.B = 2

	if !reflect.DeepEqual(*got, want) {
		t.Error("received value differs from value which was expected to get")
	}
}
