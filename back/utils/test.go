package utils

import (
	"strings"
	"testing"
)

func Assert(name string, eval func(string) string, t *testing.T) {
	ansFile := CreatePath(name + "_ans.txt")
	expectFile := CreatePath(name + ".txt")

	ans, err := Load(ansFile)
	if err != nil {
		t.Fatal(err)
	}

	exp := eval(expectFile)
	if strings.Trim(ans, "\n") != strings.Trim(exp, "\n") {
		t.Fatal("Not Equal")
	}

}
