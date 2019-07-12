package create

import "testing"

func CreateEmptyFileTest(t *testing.T) {
  CompleteFlag := CreateEmptyFile()
  if !CompleteFlag { t.Error("CreateEmptyFile() failed") }
}
