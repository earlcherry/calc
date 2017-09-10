# Calc

Simple bash calculator written on Go

## Usage:

```shell
$ calc <expression>
```

*expression* - any arithmetic expression supported by golang

## Examples

```shell
$ calc 1+1
2
$ calc 1 + 1
2
$ calc 3 + 3 * 3
12
$ calc "(3 + 3) * 3"
18
```

Supported operands: +, -, *, /, %

## Installation

```shell
$ go get
$ go build calc.go
$ sudo mv ./calc /usr/bin/local/=
```

Now you can evaluate expressions simply by typing `=` in your shell:

```shell
$ = 1 + 1
2
```
