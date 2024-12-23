package beetlejuice

import (
  "testing"
)

func TestLookup(t *testing.T) {
  bj := NewBeetleJuice().WithFilePath("data/bj.yaml")

  name := "foo"
  p, err := bj.Lookup(name)
  if err != nil {
    t.Log("name", name, "not found")
    t.Fail()
  }

  if p != "bar" {
    t.Log("name", name, "p", p, "mismatch")
    t.Fail()
  }

  name = "nonexist"
  p, err = bj.Lookup(name)
  if err == nil {
    t.Log("name", name, "found")
    t.Fail()
  }
}
