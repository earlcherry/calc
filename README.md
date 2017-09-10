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
