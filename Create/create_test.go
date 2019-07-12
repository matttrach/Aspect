package create

import "testing"

func CreateEmptyFileTest(t *testing.T) {
  if !CreateEmptyFile() { t.Error("CreateEmptyFile() failed") }
}
