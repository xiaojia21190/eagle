package base

import "testing"

func TestModuleVersion(t *testing.T) {
	v, err := ModuleVersion("golang.org/x/mod")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestSnakeMod(t *testing.T) {
	out := SnakeMod()
	t.Log(out)
}
