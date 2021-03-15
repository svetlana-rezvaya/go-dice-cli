# go-dice-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-dice-cli)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-dice-cli)

The utility for generating dice throws.

## Installation

```
$ go get github.com/svetlana-rezvaya/go-dice-cli
```

## Usage

```
$ go-dice-cli -h | -help | --help
$ go-dice-cli [-unicode] <dice>
$ go-dice-cli [-unicode] -dice STRING
$ go-dice-cli [-unicode] -throws INTEGER -faces INTEGER
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-dice STRING` &mdash; number of throws and dice faces in the [dice notation](https://en.wikipedia.org/wiki/Dice_notation);
- `-throws INTEGER` &mdash; number of throws;
- `-faces INTEGER` &mdash; number of dice faces;
- `-unicode` &mdash; use Unicode to show dices.

Arguments:

- `<dice>` &mdash; number of throws and dice faces in the [dice notation](https://en.wikipedia.org/wiki/Dice_notation).

## Output Example

Without Unicode:

```
2, 8
minimum: 2
maximum: 8
sum: 10
```

With Unicode:

```
⚁, ⚅⚁
minimum: ⚁
maximum: ⚅⚁
sum: ⚄⚄
```

## License

The MIT License (MIT)

Copyright &copy; 2021 svetlana-rezvaya
