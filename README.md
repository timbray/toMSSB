# toMSSB

Processes text, converting alphabetic ASCII letters
and digits (**a** through **z**, 
**A** through **Z**, and **0** through **9**) into their 
equivalents in Unicode's MATHEMATICAL SANS-SERIF BOLD 
range, starting at 0x1D5D4 and 0x1D5EE for capital and 
lower-case letters, and 0x1D7EC for numeric digits.

All other characters are left as-is.

Output is in UTF-8.

If there are command-line arguments, they are converted and
written to Standard Output. Failing that, lines are read from
standard input, converted, and written to standard output.

To build:

```shell
go build -o bin/tomssb
```