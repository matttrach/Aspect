package create

import "testing"

func TestCreateEmptyFile(t *testing.T) {
  want := true
  got := CreateFile("empty.txt","")
  if got != want { t.Errorf("CreateEmptyFile failed") }
}
