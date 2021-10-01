package main

import (
    "io"
    "os"
    "bufio"
    "flag"
    "fmt"
)

type Machine struct {
  index int
  size int
  mem [1024]byte
  reader bufio.Reader
  writer bufio.Writer
}

func (m *Machine) read() byte {
  b, err := m.reader.ReadByte()
  if err != nil {
    panic(err)
  }
  return b
}

func (m *Machine) write(b *byte) {
  var err error
  err = m.writer.WriteByte(*b)
  if err != nil {
    panic(err)
  }
}

func (m *Machine) compute(code []byte) {
  for i := 0; i < len(code); i++ {
      char := code[i]
      switch char {
        case '>': m.index = (m.index + m.size + 1) % m.size
        case '<': m.index = (m.index + m.size - 1) % m.size
        case '+': m.mem[m.index] += 1
        case '-': m.mem[m.index] -= 1
        case '.':
          output := m.mem[m.index]
          m.write(&output)
        case ',':
          input := m.read()
          m.mem[m.index] += input
        case '[':
          if m.mem[m.index] == 0 {
            for loop := 1; loop > 0; {
              i++
              if code[i] == '[' { loop++ }
              if code[i] == ']' { loop-- }
            }
          }
        case ']':
          if m.mem[m.index] != 0 {
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

  vm.compute(buf)

  vm.writer.Flush()
}