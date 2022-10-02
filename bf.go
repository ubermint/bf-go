package main

import (
    . "github.com/ubermint/bf-go/machine"
    "io"
    "os"
    "bufio"
    "flag"
    "fmt"
)



func reader(path string) io.Reader {
    var err error
    r := os.Stdin
    if len(path) > 0 {
        r, err = os.Open(path)
        if err != nil {
            panic(err)
        }
    }
    return r
}

func main() {
  path := flag.String("file", "", "a string")
  // unicode := flag.Bool("utf", false, "a bool")
  flag.Parse()

  input := reader(*path)
  fmt.Print("Brainfuck say:\n")

  buf := make([]byte, 0)
  scanner := bufio.NewScanner(input)
  scanner.Split(bufio.ScanBytes)
  for scanner.Scan() {
    b := scanner.Bytes()[0]
    buf = append(buf, b)
  }

  if err := scanner.Err(); err != nil {
          panic(err)
  }

  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)

  const size int = 1024
  var mem[size]byte
  vm := Machine{0, size, mem, *reader, *writer}

  vm.Compute(buf)

  vm.Writer.Flush()
}