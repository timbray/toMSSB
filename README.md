# toMSSB

Processes text, converting alphabetic ASCII letters (**a** 
through **z** 
and **A** through **Z**) into their equivalents in Unicode's 
MATHEMATICAL SANS-SERIF BOLD range, starting at 0x1D5D4 and
0x1D5EE for capital and lower-case respectively.  

All other characters are left as-is.

Output is in UTF-8.

If there are command-line arguments, they are converted and
written to Standard Output. Failing that, lines are read from
standard input, converted, and written to standard output.

To build:

```shell
go build -o bin/tomssb
```