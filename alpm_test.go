package alpm

import "testing"

func TestRelease(t *testing.T) {
  Init()
  err := Release()
  if err != nil {
    t.Fail()
  }
}

func TestInit(t *testing.T) {
  err := Init()
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

func TestRun(t *testing.T) {
  test()
}
