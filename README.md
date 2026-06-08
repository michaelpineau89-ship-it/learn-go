# learn-go

Learning workspace for Go examples, TDD practice, and experimentation.

## Layout

- `cmd/examples/<name>/main.go`: runnable examples (Go by Example style)
- `tdd/`: workspace for Learn Go with Tests exercises
- `libtest/`: small library packages with unit tests
- `playground/`: scratchpad for quick experiments

## Run examples

```bash
go run ./cmd/examples/channels
go run ./cmd/examples/atomic-counter
```

## Run playground

```bash
go run ./playground
```

## Run tests

```bash
go test ./...
```