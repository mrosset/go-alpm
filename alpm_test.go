package alpm

import "testing"

func TestVersion(t *testing.T) {
  if Version() != "5.0.0" {
    t.Fail()
  }
}
