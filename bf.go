package main

import (
    . "github.com/ubermint/bf-go/machine"
    "io"
    "os"
    "bufio"
    "flag"
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

func run(buf []byte) {
    const size int = 4096
    var mem[size]byte
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)

    vm := Machine{0, size, mem, *reader, *writer}
    vm.Compute(buf)

    vm.Writer.Flush()
}

func main() {
    path := flag.String("file", "", "a string")
    // unicode := flag.Bool("utf", false, "a bool")
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

    run(buf)
}