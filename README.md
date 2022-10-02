### Brainfuck implementation in golang
#### About
- [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck)
- [esolang](https://esolangs.org/wiki/Brainfuck)

#### Build
```sh
git clone https://github.com/ubermint/bf-go
cd bf-go
go build
```

#### Usage
Input from file:
```sh
./bf-go --file=path/to/file.bf
```

Test samples:
```sh
go test
```

Input from stdin:
```sh
./bf-go <here paste your bf program>
<Ctrl+D to end input>
<Ctrl+C to abort>
```