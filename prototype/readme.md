# Neu ASM

When neu code is assembled it is placed in a .n file.
This file can be run by the neu interpreter vm.

# Neu Bytecode

Neu is a stackbased bytecode that runs on a small vm.
It uses a single stack and 1 byte opcodes.
Code files should be stored in .jy files

# Op Codes Reference

op codes are 1 byet a pieces and can have 0, 1, or 2 arguments.
the stack consists of 64 bit cells
(Note: that this is really wastefull of memory now and will likely change).
you must specify the max height of the stack when your program starts
(Note: this may change as the language grows so you can request a larger stack if needed).
Numerical literals can be either in decimal notation (1, 10, 50), hexidecimal notation (0xff, 0x0a), or binary notation (0b00001100).
Memory addresses can be in either hexidecimal notation (#0x5a, #0x7b) or binary notation (#0b00001110).

This table lists all the available op codes, the # indicates an argument which is either a memory address or a literal value.

| Name       | Usage   | Hex  | Description                                                                                          |
| ---------- | ------- | ---- | ---------------------------------------------------------------------------------------------------- |
| push       | < #     | 0x01 | push a new value(#) onto the stack                                                                   |
| add        | +       | 0x02 | pop the top two values off the stack, add them together and push the result onto the stack           |
| minus      | -       | 0x03 | pop the top two values off the stack, subtrack them from one another, push the result onto the stack |
| pop        | > #     | 0x04 | pop the top value off the stack and write it to memory address(#)                                    |
| stack size | []      | 0x05 | set the size of the stack so the vm can size the stack                                               |

# Example Programs
### example 1
this is a simple program for adding the numbers 5 and 10 together

```
0: [] 2
1: < 5
2: < 10
3: +
```

the line numbers are only here for reference and are not present in an actual program.
The first line defines the size of the stack. The stack is not allowed to grow past this size.
Here the max size is defined as two.
Line 1 pushes the literal value 5 onto the stack with the push opcode(<).
Line 2 pushes the literal value 10 onto the stack in the same way.
Line 3 addes the top pops the top two numbers off of the stack, adds them together and then pushes the result back onto the stack.
Note the poping a value off the stack does not clear the value, it simply moves the stack pointer.
Thus the final state of the stack after this progam runs is:

```
  ------
  | 10 |
  ------
> | 15 |
  ------
```