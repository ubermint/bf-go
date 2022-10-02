package machine

import (
    "bufio"
)

type Machine struct {
  Index int
  Size int
  Mem [4096]byte
  Reader bufio.Reader
  Writer bufio.Writer
}

func (m *Machine) Read() byte {
  b, err := m.Reader.ReadByte()
  if err != nil {
    panic(err)
  }
  return b
}

func (m *Machine) Write(b *byte) {
  var err error
  err = m.Writer.WriteByte(*b)
  if err != nil {
    panic(err)
  }
}

func (m *Machine) Compute(code []byte) {
  for i := 0; i < len(code); i++ {
      char := code[i]
      switch char {
        case '>': m.Index = (m.Index + m.Size + 1) % m.Size
        case '<': m.Index = (m.Index + m.Size - 1) % m.Size
        case '+': m.Mem[m.Index] += 1
        case '-': m.Mem[m.Index] -= 1
        case '.':
          output := m.Mem[m.Index]
          m.Write(&output)
        case ',':
          input := m.Read()
          m.Mem[m.Index] += input
        case '[':
          if m.Mem[m.Index] == 0 {
            for loop := 1; loop > 0; {
              i++
              if code[i] == '[' { loop++ }
              if code[i] == ']' { loop-- }
            }
          }
        case ']':
          if m.Mem[m.Index] != 0 {
            for loop := 1; loop > 0; {
              i--
              if code[i] == ']' { loop++ }
              if code[i] == '[' { loop-- }
            }
          }
      }
    }
}