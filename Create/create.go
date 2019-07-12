package create

import "os"

func CreateEmptyFile() bool {
  f, err := os.Create("../../empty.txt")
  check(err)
  f.Close()
  return true
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
