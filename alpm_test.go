package alpm

import "testing"

func TestRelease(t *testing.T) {
  Initialize()
  err := Release()
  if err != nil {
    t.Fail()
  }
}

func TestInitialize(t *testing.T) {
  err := Initialize()
  defer Release()
  if err != nil {
    t.Fail()
  }
}

func TestVersion(t *testing.T) {
  if Version() != "5.0.3" {
    t.Fail()
  }
}
