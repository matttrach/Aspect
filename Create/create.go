package create

import "os"

func CreateFile(name, contents string) bool {
  f, err := os.Create("../../" + name)
  check(err)
  f.Close()
  return true
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
