# Neu ASM

When neu code is assembled it is placed in a .n file.
This file can be run by the neu interpreter vm.

# Neu Bytecode

Neu is a stackbased bytecode that runs on a small vm.
It uses a single stack and 1 byte opcodes.
Neu byte code is be stored in .nb files

### comments
  
Comments in neu are specified by two forward slashes //.
Anything that comes after these 2 symboles on a line will be ignored by the assembler.

### pointers/ memory addresses

Memory addresses can be in either hexidecimal notation (#0x5a, #0x7b) or binary notation (#0b00001110).
Pointers are 64 bits numbers.

### labels

Labels are specified by a name sourounded by brackets for example [label name].
Label names can include numbers, letters and symboles and can be used intercangably with memory addresses.
Labels are converted to explicit memory addresses by the assembler.

### number literals

Numerical literals can be either in decimal notation (1, 10, 50), hexidecimal notation (0xff, 0x0a), or binary notation (0b00001100).

### jump statments

Jump statments come in two flavors, normal jumps or conditional jumps
Normal jumps will pop the top 8 bytes off the stack and move to that line in the program and then continue execution from that point.
Conditional jumps will first pop the first 2 bytes (or sets of bytes) off the stack and test them agains each other.
If the condition returns true it behaves the same as a normal jump.
If it returns false it pushes the popped values back onto the stack and exxecution continues on to the next instruction.

# Op Codes Reference

op codes are 1 byet a pieces and can have 0, 1, or 2 arguments.
the stack consists of 1 byte cells and live in the same place as main memory.
This table lists all the available op codes, the # indicates an argument which is either a memory address or a literal value.
the [L] indicates an argument which is a label to a location in the code.

| Name                  | Usage   | Hex  | Description                                                                                                                  |
| --------------------- | ------- | ---- | ---------------------------------------------------------------------------------------------------------------------------- |
| byte add              | +.      | 0x00 | pop the top two bytes off the stack, add them together and push the result onto the stack                                    |
| int16 add             | +o      | 0x01 | pop the top 4 bytes off the stack, convert them to 2 int16s, add them, and push the result onto the stack                    |
| int32 add             | +O      | 0x02 | pop the top 8 bytes off the stack, convert them to 2 int32s, add them, and push the result onto the stack                    |
| int64 add             | +       | 0x03 | pop the top 16 bytes off the stack, convert them to 2 int64s, add them, and push the result onto the stack                   |
| byte minus            | -.      | 0x04 | pop the top two bytes off the stack, subtract them from one another, push the result onto the stack                          |
| int16 minus           | -o      | 0x05 | pop the top 4 bytes off the stack, convert them to 2 int16s, subtract them, and push the result onto the stack               |
| int32 minus           | -O      | 0x06 | pop the top 8 bytes off the stack, convert them to 2 int32s, subtract them, and push the result onto the stack               |
| int64 minus           | -       | 0x07 | pop the top 16 bytes off the stack, convert them to 2 int64s, subtract them, and push the result onto the stack              |
| byte multiply         | *.      | 0x08 | pop the top two bytes off the stack, multipl the first value by the second, push the result onto the stack                   |
| int16 multiply        | *o      | 0x09 | pop the top 4 bytes off the stack, convert them to 2 int16s, multiply them, and push the result onto the stack               |
| int32 multiply        | *O      | 0x0a | pop the top 8 bytes off the stack, convert them to 2 int32s, multiply them, and push the result onto the stack               |
| int64 multiply        | *       | 0x0b | pop the top 16 bytes off the stack, convert them to 2 int64s, multiply them, and push the result onto the stack              |
| byte divide           | /.      | 0x0c | pop the top two bytes off the stack, divide the first value by the second, push the result onto the stack                    |
| int16 divide          | /o      | 0x0d | pop the top 4 bytes off the stack, convert them to 2 int16s, divide them, and push the result onto the stack                 |
| int32 divide          | /O      | 0x0e | pop the top 8 bytes off the stack, convert them to 2 int32s, divide them, and push the result onto the stack                 |
| int64 divide          | /       | 0x0f | pop the top 16 bytes off the stack, convert them to 2 int64s, divide them, and push the result onto the stack                |
| byte push             | <. #    | literal(0x10) addr(0x45) | push a new byte onto the stack (# can be an address or a literal)                                        |
| int16 push            | <o #    | literal(0x11) addr(0x46) | push 2 new bytes onto the stack as an int16 (# can be an address or a literal)                           |
| int32 push            | <O #    | literal(0x12) addr(0x47) | push 2 new bytes onto the stack as an int32 (# can be an address or a literal)                           |
| int64 push            | < #     | literal(0x13) addr(0x48) | push 2 new bytes onto the stack as an int64 (# can be an address or a literal)                           |
| byte pop              | >.      | 0x14 | pop the top 9 byte off the stack and writes the last byte to the memory address in the first 8 bytes                         |
| int16 pop             | >o      | 0x15 | pop the top 10 byte off the stack and writes the last 2 bytes to the memory address in the first 8 bytes                     |
| int32 pop             | >O      | 0x16 | pop the top 12 byte off the stack and writes the last 4 bytes to the memory address in the first 8 bytes                     |
| int64 pop             | >       | 0x17 | pop the top 16 byte off the stack and writes the last 8 bytes to the memory address in the first 8 bytes                     |
| bitwise or            | \|.     | 0x18 | pop the top two bytes off the stack, bitwise or's them together and push the result onto the stack                           |
| int16 bit or          | \|o     | 0x19 | pop the top 4 bytes off the stack, convert them to 2 int16s, or them together, and push the result onto the stack            |
| int32 bit or          | \|O     | 0x1a | pop the top 8 bytes off the stack, convert them to 2 int32s, or them together, and push the result onto the stack            |
| int64 bit or          | \|      | 0x1b | pop the top 16 bytes off the stack, convert them to 2 int64s, or them together, and push the result onto the stack           |
| bitwise and           | &.      | 0x1c | pop the top two bytes off the stack, bitwise and's them together and push the result onto the stack                          |
| int16 bit and         | &o      | 0x1d | pop the top 4 bytes off the stack, convert them to 2 int16s, and them together, and push the result onto the stack           |
| int32 bit and         | &O      | 0x1e | pop the top 8 bytes off the stack, convert them to 2 int32s, and them together, and push the result onto the stack           |
| int64 bit and         | &       | 0x1f | pop the top 16 bytes off the stack, convert them to 2 int64s, and them together, and push the result onto the stack          |
| bitwise xor           | ^.      | 0x20 | pop the top two bytes off the stack, bitwise xor's them together and push the result onto the stack                          |
| int16 bit xor         | ^o      | 0x21 | pop the top 4 bytes off the stack, convert them to 2 int16s, xor them together, and push the result onto the stack           |
| int32 bit xor         | ^O      | 0x22 | pop the top 8 bytes off the stack, convert them to 2 int32s, xor them together, and push the result onto the stack           |
| int64 bit xor         | ^       | 0x23 | pop the top 16 bytes off the stack, convert them to 2 int64s, xor them together, and push the result onto the stack          |
| bitwise left shift    | <<.     | 0x24 | pop the top value off the stack, pop another byte off the stack and shift it left that many places                           |
| int16 bit left shift  | <<o     | 0x25 | pop the top 2 bytes off the stack, convert it to an int16, pop another byte off the stack and shift it left that many places |
| int32 bit left shift  | <<O     | 0x26 | pop the top 4 bytes off the stack, convert it to an int32, pop another byte off the stack and shift it left that many places |
| int64 bit left shift  | <<      | 0x27 | pop the top 8 bytes off the stack, convert it to an int64, pop another byte off the stack and shift it left that many places |
| bitwise right shift   | >>.     | 0x28 | pop the top value off the stack, shift it right # places and push the result onto the stack                                  |
| int16 bit right shift | >>o     | 0x29 | pop the top 2 bytes off the stack, convert it to an int16, shift it right # places and push the result onto the stack        |
| int32 bit right shift | >>O     | 0x2a | pop the top 4 bytes off the stack, convert it to an int32, shift it right # places and push the result onto the stack        |
| int64 bit right shift | >>      | 0x2b | pop the top 8 bytes off the stack, convert it to an int64, shift it right # places and push the result onto the stack        |
| jump if greater       | ?>.     | 0x2c | jump the execution pointer to the specified memory address if the top byte on the stack is larger than the second byte       |
| int16 jump if greater | ?>o     | 0x2d | jump the execution pointer to the specified memory address if the top int16 on the stack is larger than the second int16     |
| int32 jump if greater | ?>O     | 0x2e | jump the execution pointer to the specified memory address if the top int32 on the stack is larger than the second int32     |
| int64 jump if greater | ?>      | 0x2f | jump the execution pointer to the specified memory address if the top int64 on the stack is larger than the second int32     |
| jump if less          | ?<.     | 0x30 | jump the execution pointer to the specified memory address if the top byte on the stack is smaller than the second byte      |
| int16 jump if less    | ?<o     | 0x31 | jump the execution pointer to the specified memory address if the top int16 on the stack is smaller than the second int16    |
| int32 jump if less    | ?<O     | 0x32 | jump the execution pointer to the specified memory address if the top int32 on the stack is smaller than the second int32    |
| int64 jump if less    | ?<      | 0x33 | jump the execution pointer to the specified memory address if the top int64 on the stack is smaller than the second int32    |
| jump                  | |>      | 0x34 | jump the execution pointer to the specified memory address                                                                   |
| byte mod              | %.      | 0x35 | pop the top two bytes off the stack, mod the first value by the second, push the result onto the stack                       |
| int16 mod             | %o      | 0x36 | pop the top 4 bytes off the stack, convert them to 2 int16s, mod them, and push the result onto the stack                    |
| int32 mod             | %O      | 0x37 | pop the top 8 bytes off the stack, convert them to 2 int32s, mod them, and push the result onto the stack                    |
| int64 mod             | %       | 0x38 | pop the top 16 bytes off the stack, convert them to 2 int64s, mod them, and push the result onto the stack                   |
| push byte 0           | <0.     | 0x39 | push a zero byte onto the stack                                                                                              |
| push int16 0          | <0o     | 0x3a | push a zero int16 onto the stack                                                                                             |
| push int32 0          | <0O     | 0x3b | push a zero int32 onto the stack                                                                                             |
| push int64 0          | <0      | 0x3c | push a zero int64 onto the stack                                                                                             |
| dec byte              | --.     | 0x3d | pop the top byte off the stack subtract one and push the result back onto the stack                                          |
| dec int16             | --o     | 0x3e | pop the top 2 bytes off the stack, convert them to an int16, subtract one and push the result onto the stack                 |
| dec int32             | --O     | 0x3f | pop the top 4 bytes off the stack, convert them to an int32, subtract one and push the result onto the stack                 |
| dec int64             | --      | 0x40 | pop the top 8 bytes off the stack, convert them to an int64, subtract one and push the result onto the stack                 |
| inc byte              | ++.     | 0x41 | pop the top byte off the stack, add one and push the result onto the stack                                                   |
| inc int16             | ++o     | 0x42 | pop the top 2 bytes off the stack, convert them to an int16, subtract one and push the result onto the stack                 |
| inc int32             | ++O     | 0x43 | pop the top 4 bytes off the stack, convert them to an int32, subtract one and push the result onto the stack                 |
| inc int64             | ++      | 0x44 | pop the top 8 bytes off the stack, convert them to an int64, subtract one and push the result onto the stack                 |
| label                 | [L]     | ---- | label marks a section of the code.                                                                                           |
| name address          | _ = #   | ---- | specifiy a name for a numerical constant that can be used later in your code                                                 |
| memory address        | #       | ---- | converts a numerical literal into a memory address                                                                           |

# Example Programs
### example 1
this is a simple program for adding the numbers 5 and 10 together

```
0: <. 5
1: <. 10
2: +.
```

the line numbers are only here for reference and are not present in an actual program.
The first line defines the size of the stack. The stack is not allowed to grow past this size.
Here the max size is defined as two.
Line 0 pushes the literal value 5 onto the stack with the push opcode(<).
Line 1 pushes the literal value 10 onto the stack in the same way.
Line 2 addes the top pops the top two numbers off of the stack, adds them together and then pushes the result back onto the stack.
Note the poping a value off the stack does not clear the value, it simply moves the stack pointer.
Thus the final state of the stack after this progam runs is:

```
  ------
  | 10 |
  ------
> | 15 |
  ------
```

# Development Thoughts
* Should we use 32 or maybe even 64 bits as the default size for the buffer? There are some challenges that come along with that but some benefits as well.
  - real programs will be working with 32/ 64 bit numbers much more often than small 8 or 16 bit numbers.
  - wastes space and given these need to be sent over the network that's important (although it does so in a way that may be simple to compress?).
  - how do we get smaller values (like byte or int16) to work correctly/ overflow in the right way?
  - we are gonna waste cpu cycles combining bytes into int64 or int32's all the time.
  - what's the point of having a 1 bit op code if it has to be stuck into a 32 bit int anyway?

* fmt nb code to have consistent spacing (neu fmt?)
* [long term] does this stack base setup allow for SIMD? should I change the bytecode setup to make SIMD easier to achieve?
* add a magic string as a header?
* add a version for future proofness?
* seems like I'm working at the resolution of my pointers a lot so 64 bits
* web asm and java byte code both have the concept of 'locals' should I adopt something like this?
* should their be support for floats? signed ints?