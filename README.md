### Brainfuck implementation in golang
#### About
- [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck)
- [esolang](https://esolangs.org/wiki/Brainfuck)

#### Build
```sh
git clone https://github.com/ubermint/bf-go
cd bf-go
go build bf.go
```

#### Usage
Input from output:
```sh
./bf --utf --file=path/to/file
```

Input from stdin:
```sh
./bf < path/to/file
```
```sh
./bf
<here paste your bf program>
<Ctrl+D to end input>
<Ctrl+C to abort>
```

Test samples:
```sh
go test
```