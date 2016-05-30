package main

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func main() {
	tests := []struct {
		in           string
		impPath, pkg string
		ok           bool
	}{
		{"", "", "", false},
		{"foo", "", "foo", true},
		{"github.com/golang/bar", "github.com/golang/bar", "bar", true},
		{"github.com/golang/bar;baz", "github.com/golang/bar", "baz", true},
	}
	for _, tc := range tests {
		d := &generator.FileDescriptor{
			FileDescriptorProto: &descriptor.FileDescriptorProto{
				Options: &descriptor.FileOptions{
					GoPackage: &tc.in,
				},
			},
		}
		fmt.Println("d.PackageName()")
		fmt.Println(d.PackageName())
	}
}
