package main

import s "strings"
import "fmt"
import "os"
import "math/rand"
import "time"

func main() {
  seed := rand.NewSource(time.Now().UnixNano())
  var r1 = rand.New(seed)
  var p = fmt.Println

  text := os.Args[1]

  var length = len(text)
  var indexToCapitalize []int
  var halfLength = length / 2

  for i := 0; i <= halfLength; i++ {
    indexToCapitalize = append(indexToCapitalize, r1.Intn(length))
  }

  var newString []string

  for i := 0; i < length; i++ {
    var currentChar = string(text[i])
    if (Contains(indexToCapitalize, i)) {
      currentChar = s.ToUpper(currentChar)
    }
    newString = append(newString, currentChar)
  }

  p(s.Join(newString, ""))
}

func Contains(a []int, x int) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}
