# Overview

```
Write a command line program in the language of your choice that will take operations on fractions as an input and produce a fractional result.

Legal operators shall be *, /, +, - (multiply, divide, add, subtract).

Operands and operators shall be separated by one or more spaces.

Mixed numbers will be represented by whole_numerator/denominator. e.g. "3_1/4".

Improper fractions and whole numbers are also allowed as operands.

Example run
? 1/2 * 3_3/4
= 1_7/8

? 2_3/8 + 9/8
= 3_1/2
```

## Building and Executing

```bash
cd fractionator
go run .
```

## Running Tests

```bash
cd fraction
go test
```

## Initial Setup (just for information)

You don't have to do this.

```bash
go mod init bcr/fraction
```
Is how I made the `go.mod`