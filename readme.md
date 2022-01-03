# Working With This Code

### building
Building this code is as simple running `make build`.
This will build both `neuBi` which is the new byte code assembler, as well as neuVM which is the neu byte code interpreter

### testing
Testing this code is as simple as running `make test`.
This will run all the test files in this repo

### benchmarks
Benchmarking the code can be done by running `make bench`.
Currently, the only benchmarks are in [core/util_bench_test.go](https://github.com/bjatkin/NEU/blob/main/core/util_bench_test.go).
These benchmarks were written to look at the performance differences between I64tob and I64asb (and similar methods for the other int types).
Looking at these benchmarks you can see the huge performance gain we get by using unsafe pointers to convert from byte slices into integers and vice versa.

# Neu Assembler

The neu text code assembler is called neuBi (pronounced newbie).
You can build neuBi by running `make build`.
Once you have neuBi built you can run `./neuBi -h` to get more info about how it works.
The most common way to use neuBi is to assemble .nb files.
Simply run `./neuBi my_file.nb` to assemble `my_file.nb` into `my_file.n`.

# Neu Bytecode and Neu Text Code

Neu is a stackbased bytecode that runs on a small vm which uses a single stack and 1 byte opcodes.
Neu text code is the textual representation of neu byte code.
Neu text code was designed to be human readable and can be assembled into neu byte code by [neuBi](#neu-assembler)
Neu text code is stored in .nb files.
Neu byte code is stored in .n files in a formt which can be run by the neuVM.

# Neu VM

Neu VM is the virtual machine that runs neu byte code.
You can build the neuVM by running `make build`.
Once the new VM is built you can run `neuVM my_file.n` to run your compiled neu byte code.

# Writing Neu Text Code
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
Hex and binary literals must include their respective prefixes to be processed correctly.
All numbers are considers signed twos compliment nubers with 2 exceptions
  1) addresses are always considered unsigned 64 bit integers
  2) bytes are always considered unsigned

### jump statments

Jump statments come in two flavors, normal jumps or conditional jumps
Normal jumps will pop the top 8 bytes off the stack and move to that line in the program and then continue execution from that point.
Conditional jumps will first pop the first 2 bytes (or sets of bytes) off the stack and test them agains each other.
If the condition returns true it behaves the same as a normal jump.
If it returns false the jump address is popped off the stack and discarded and execution continues on to the next instruction.

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
| byte push             | <.  #   | 0x10 | push a new byte onto the stack, # must be a literal                                                                          |
| int16 push            | <o  #   | 0x11 | push 2 new bytes onto the stack as an int16, # must be a literal                                                             |
| int32 push            | <O  #   | 0x12 | push 4 new bytes onto the stack as an int32, # must be a literal                                                            |
| int64 push            | <   #   | 0x13 | push 8 new bytes onto the stack as an int64, # must be a literal                                                            |
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
| bitwise left shift    | <<.     | 0x24 | pop the top byte off the stack as count, shift the next byte count times to the left                                         |
| int16 bit left shift  | <<o     | 0x25 | pop the top byte off the stack as count, shift the next int16(2 bytes) count times to the left                               |
| int32 bit left shift  | <<O     | 0x26 | pop the top byte off the stack as count, shift the next int32(4 bytes) count times to the left                               |
| int64 bit left shift  | <<      | 0x27 | pop the top byte off the stack as count, shift the next int64(8 bytes) count times to the left                               |
| bitwise right shift   | >>.     | 0x28 | pop the top byte off the stack as count, shift the next byte count times to the right                                        |
| int16 bit right shift | >>o     | 0x29 | pop the top byte off the stack as count, shift the next nt16(2 bytes) count times to the right                               |
| int32 bit right shift | >>O     | 0x2a | pop the top byte off the stack as count, shift the next nt32(4 bytes) count times to the right                               |
| int64 bit right shift | >>      | 0x2b | pop the top byte off the stack as count, shift the next nt64(8 bytes) count times to the right                               |
| jump if greater       | ?>.     | 0x2c | jump the execution pointer to the memory address on the stack if the top byte on the stack is larger than the second byte    |
| int16 jump if greater | ?>o     | 0x2d | jump the execution pointer to the memory address on the stack if the top int16 on the stack is larger than the second int16  |
| int32 jump if greater | ?>O     | 0x2e | jump the execution pointer to the memory address on the stack if the top int32 on the stack is larger than the second int32  |
| int64 jump if greater | ?>      | 0x2f | jump the execution pointer to the memory address on the stack if the top int64 on the stack is larger than the second int32  |
| jump if less          | ?<.     | 0x30 | jump the execution pointer to the memory address on the stack if the top byte on the stack is smaller than the second byte   |
| int16 jump if less    | ?<o     | 0x31 | jump the execution pointer to the memory address on the stack if the top int16 on the stack is smaller than the second int16 |
| int32 jump if less    | ?<O     | 0x32 | jump the execution pointer to the memory address on the stack if the top int32 on the stack is smaller than the second int32 |
| int64 jump if less    | ?<      | 0x33 | jump the execution pointer to the memory address on the stack if the top int64 on the stack is smaller than the second int32 |
| jump                  | |>      | 0x34 | jump the execution pointer to the memory address on the stack                                                                |
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
| byte push             | <#.     | 0x45 | pop the top uint64 off the stack as an address, push the byte at that address onto the stack                                 |
| int16 push            | <#o     | 0x46 | pop the top uint64 off the stack as an address, push the int16 at that address onto the stack                                |
| int32 push            | <#O     | 0x47 | pop the top uint64 off the stack as an address, push the int32 at that address onto the stack                                |
| int64 push            | <#      | 0x48 | pop the top uint64 off the stack as an address, push the int64 at that address onto thte stack                               |
| break point           | (/)     | 0x49 | break point for debugging                                                                                                    |
| byte duplicate stack  | X2.     | 0x4a | push the top byte onto the stack again                                                                                       |
| int16 duplicate stack | X2o     | 0x4b | push the top int16 onto the stack again                                                                                      |
| int32 duplicate stack | X2O     | 0x4c | push the top int32 onto the stack again                                                                                      |
| int64 duplicate stack | X2      | 0x4d | push the top int64 onto the stack again                                                                                      |
| label                 | [L]     | ---- | label marks a section of the code.                                                                                           |
| name address          | _ = #   | ---- | specifiy a name for a numerical constant that can be used later in your code                                                 |
| memory address        | #       | ---- | converts a numerical literal into a memory address                                                                           |

# Debugger
The neu interpreter comes with a debugger built in to make it easier to write your code.
The breakpoint symbole is `(/)`, and will halt execution of you program and start the debugger.
To get more information on the debugger simply type `?` or `help` and you will be given a list of debugger commands.
Remember that your code will run every frame so you code may halt every frame depending on where you put your break point.
Generally the debugger presents its data as hex numbers but occationally it shows them as decimal numbers as well.
When this is the case the decimal number is printed first followed by the `|` symbole and then the hex number with an explicit hex prefix `0x`
For example the `spt` or stack pointer address command will return `[ 100 | 0x00000064 ]`.

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
00000010 | 10 | 
00000011 | 15 | < stack pointer
```

### drawing to the screen
drawing to the screen is as simple as writing to the correct location in memory. Pixels are 2 bits each and there are 4 pixels packed into each byte
```
       pxl1 pxl2 pxl3 pxl4
byte: [ 00   11   01   10 ]
```

You can see an example of how to draw to the screen [here](https://github.com/bjatkin/NEU/blob/main/examples/draw.nb)

### hello world
Neu byte code is very low level and bare bones.
This mean writing a hello world example is not as simple as it might be in a higher level language like python or even c.
You can see an example of a hello world program [here](https://github.com/bjatkin/NEU/blob/main/examples/hello_world.n).
This program starts by loading a custom font into memory.
Once this is done the code uses 3 "functions" to write the text 'HELLO WORLD!' to the screen.
While these are not true functions in the strict sense of the word the work in a somewhat similar manner so I use the terminology moving forward.
The first is a function to draw strings to the screen.
The second is a function that draws individual characters to the screen which the PrintString function calls on a loop.
The final function draws a row of pixels from a character, this is called in a loop from the PrintChar function.
Together these 3 functions write the text to the screen.