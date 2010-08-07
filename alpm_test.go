package alpm

import "testing"
import "fmt"

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
  if Version() != "5.0.0" {
    t.Fail()
  }
}
