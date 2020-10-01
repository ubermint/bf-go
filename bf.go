package main

import (
    "io"
    "os"
    "bufio"
    "flag"
    "fmt"
)

type VM struct {
  index int
  size int
  arr [1024]byte
  reader bufio.Reader
  writer bufio.Writer
}

func (x *VM) read() byte {
  b, err := x.reader.ReadByte()
  if err != nil {
    panic(err)
  }
  return b
}

func (x *VM) write(b *byte) {
  var err error
  err = x.writer.WriteByte(*b)
  if err != nil {
    panic(err)
  }
}

func (x *VM) compute(code []byte) {
  for i := 0; i < len(code); i++{
      char := code[i]
      switch char {
        case '>': x.index = (x.index + x.size + 1) % x.size
        case '<': x.index = (x.index + x.size - 1) % x.size
        case '+': x.arr[x.index] += 1
        case '-': x.arr[x.index] -= 1
        case '.':
          output := x.arr[x.index]
          x.write(&output)
        case ',':
          input := x.read()
          x.arr[x.index] += input
        case '[':
          if x.arr[x.index] == 0 {
            for loop := 1; loop > 0; {
              i++
              if code[i] == '[' { loop++ }
              if code[i] == ']' { loop-- }
            }
          }
        case ']':
          if x.arr[x.index] != 0 {
            for loop := 1; loop > 0; {
              i--
              if code[i] == ']' { loop++ }
              if code[i] == '[' { loop-- }
            }
          }
      }
    }
}

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
  unicode := flag.Bool("utf", false, "a bool")
  flag.Parse()


  input := reader(*path)

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
        var arr [size]byte
  vm := VM{0, size, arr, *reader, *writer}

  vm.compute(buf)

  vm.writer.Flush()
}