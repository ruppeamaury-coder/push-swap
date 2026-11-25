# push-swap

Two Go programs:

- `push-swap` prints the instructions that sort the integers it receives.
- `checker` reads instructions on stdin, applies them and prints `OK` or `KO`.

## Build

```console
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker
```

## Usage

```console
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | wc -l
6
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```

## Tests

```console
go test ./...
```

## Project layout

```
cmd/push-swap/      entry point of push-swap
cmd/checker/        entry point of checker
internal/stack/     the two stacks, the 11 instructions, argument and stdin parsing
internal/solver/    sorting algorithms
```

## Algorithm

Values are first replaced by their rank (0 = smallest).

- **6 elements or fewer**: breadth-first search over every reachable state of the
  two stacks. The first sorted state found is reached with the smallest possible
  number of instructions.
- **More than 6 elements** (greedy "Turk" algorithm):
  1. push two elements to `b`;
  2. while `a` has more than 3 elements, compute for each element of `a` the cost
     of moving it on top of `a` while its target in `b` (the largest smaller value)
     is moved on top of `b`, counting `rr`/`rrr` as one instruction. Push the
     cheapest one. `b` stays sorted in descending order;
  3. sort the 3 elements left in `a`;
  4. push `b` back to `a`, rotating `a` so each element lands at its place;
  5. rotate `a` until the smallest element is on top.

Typical results: 5 numbers → at most 8 instructions, 100 numbers → about 560–610,
500 numbers → about 5200–5300.
