package create

import "testing"

func TestCreateEmptyFile(t *testing.T) {
  want := true
  got := CreateEmptyFile()
  if got != want { t.Errorf("CreateEmptyFile failed") }
}
