# Calculator Program

This is a simple calculator program written in Go. It supports both Arabic and Roman numerals.

## Getting Started

To run the program, use the following command:

```bash
go run . [operand1] [operator] [operand2]
```

Replace [operand1],[operator], and [operand2] with your desired values. For example, to add 5 and 3, you would use:

``` bash
go run . 5 + 3
```

## Operators
The program supports the following operators:

* "+" for addition
* "-" for subtraction
* "*" for multiplication
* "/" for division

Note: The multiplication operator (*) must be escaped with a backslash (\) or enclosed in quotes in the command line. For example:
``` bash
go run . 5 "*" 3
```
or

``` bash
go run . 5 \* 3
```