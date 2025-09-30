package mergo_test

import (
	"testing"

	"github.com/monetr/mergo"
)

func TestMergePointer(t *testing.T) {
	type Foo struct {
		Bar *string `json:"bar"`
	}

	dst := Foo{}
	src := map[string]interface{}{
		"bar": "this is a string",
	}

	err := mergo.Map(&dst, src, mergo.WithOverride)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if *dst.Bar != "this is a string" {
		t.Errorf("failed to set nil ptr")
	}
}
