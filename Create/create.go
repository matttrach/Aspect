package create

import "os"

func CreateEmptyFile() bool {
  f, err := os.Create("/mnt/c/Users/mattt/Documents/empty.txt")
  check(err)
  f.Close()
  return true
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
