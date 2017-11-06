package cpu

import (
    "fmt"
)

type Cpu struct {
    register0 uint8
    register1 uint8
    programCounter uint8
    heapPointer uint8
    flags uint8
}

// flags:
// 0 = greater than
// 1 = less than
// 2 = zero
// 3 = non zero
// 4 = overflow
// 5 = underflow
// 6 = signing
// 7 = halt

const greaterThanSet uint8 = 0x80
const lessThanSet    uint8 = 0x40
const zeroSet        uint8 = 0x20
const nonZeroSet     uint8 = 0x10
const overflowSet    uint8 = 0x08
const underflowSet   uint8 = 0x04
const signingSet     uint8 = 0x02
const haltSet        uint8 = 0x01

const greaterThanReset uint8 = 0x7F
const lessThanReset    uint8 = 0xBF
const zeroReset        uint8 = 0xDF
const nonZeroReset     uint8 = 0xEF
const overflowReset    uint8 = 0xF7
const underflowReset   uint8 = 0xFB
const signingReset     uint8 = 0xFD
const haltReset        uint8 = 0xFE

func (c *Cpu) ReInitialize() {
    c.register0 = 0
    c.register1 = 0
    c.programCounter = 0
    c.heapPointer = 0
    c.flags = 0
}

func (c *Cpu) SetRegister0(value uint8) {
    c.register0 = value
}

func (c *Cpu) SetRegister1(value uint8) {
    c.register1 = value
}

func (c *Cpu) SetProgramCounter(value uint8) {
    c.programCounter = value
}

func (c *Cpu) SetHeapPointer(value uint8) {
    c.heapPointer = value
}

func (c Cpu) Register0() uint8 {
    return c.register0
}

func (c Cpu) Register1() uint8 {
    return c.register1
}

func (c Cpu) ProgramCounter() uint8 {
    return c.programCounter
}

func (c Cpu) HeapPointer() uint8 {
    return c.heapPointer
}

func (c *Cpu) IncrementProgramCounter() {
    c.programCounter++
}

func (c Cpu) printRegister0() {
    fmt.Printf("Register 0     : %02X\n", c.register0)
}

func (c Cpu) printRegister1() {
    fmt.Printf("Register 1     : %02X\n", c.register1)
}

func (c Cpu) printProgramCounter() {
    fmt.Printf("Program Counter: %02X\n", c.programCounter)
}

func (c Cpu) printHeapPointer() {
    fmt.Printf("Heap Pointer   : %02X\n", c.heapPointer)
}

func (c Cpu) printFlags() {
    fmt.Printf("Flags          : ")

    if c.IsGreaterThan() {
        fmt.Print("GreaterThan ")
    }
    if c.IsLessThan() {
        fmt.Print("LessThan ")
    }
    if c.IsZero() {
        fmt.Print("Zero ")
    }
    if c.IsNonZero() {
        fmt.Print("NonZero ")
    }
    if c.IsOverflow() {
        fmt.Print("Overflow ")
    }
    if c.IsUnderflow() {
        fmt.Print("Underflow ")
    }
    if c.IsSigning() {
        fmt.Print("Signing ")
    }
    if c.IsHalt() {
        fmt.Print("Halt")
    }

    fmt.Println()
}

func (c Cpu) Dump() {
    fmt.Println("CPU Contents:")
    fmt.Println("-------------")
    c.printRegister0()
    c.printRegister1()
    c.printProgramCounter()
    c.printHeapPointer()
    c.printFlags()
}

func (c *Cpu) SetGreaterThan() {
    c.flags |= greaterThanSet
}

func (c *Cpu) SetLessThan() {
    c.flags |= lessThanSet
}

func (c *Cpu) SetZero() {
    c.flags |= zeroSet
}

func (c *Cpu) SetNonZero() {
    c.flags |= nonZeroSet
}

func (c *Cpu) SetOverflow() {
    c.flags |= overflowSet
}

func (c *Cpu) SetUnderflow() {
    c.flags |= underflowSet
}

func (c *Cpu) SetSigning() {
    c.flags |= signingSet
}

func (c *Cpu) SetHalt() {
    c.flags |= haltSet
}

func (c *Cpu) ResetGreaterThan() {
    c.flags &= greaterThanReset
}

func (c *Cpu) ResetLessThan() {
    c.flags &= lessThanReset
}

func (c *Cpu) ResetZero() {
    c.flags &= zeroReset
}

func (c *Cpu) ResetNonZero() {
    c.flags &= nonZeroReset
}

func (c *Cpu) ResetOverflow() {
    c.flags &= overflowReset
}

func (c *Cpu) ResetUnderflow() {
    c.flags &= underflowReset
}

func (c *Cpu) ResetSigning() {
    c.flags &= signingReset
}

func (c *Cpu) ResetHalt() {
    c.flags &= haltReset
}

func (c Cpu) IsGreaterThan() bool {
    return c.flags & greaterThanSet != 0x00
}

func (c Cpu) IsLessThan() bool {
    return c.flags & lessThanSet != 0x00
}

func (c Cpu) IsZero() bool {
    return c.flags & zeroSet != 0x00
}

func (c Cpu) IsNonZero() bool {
    return c.flags & nonZeroSet != 0x00
}

func (c Cpu) IsOverflow() bool {
    return c.flags & overflowSet != 0x00
}

func (c Cpu) IsUnderflow() bool {
    return c.flags & underflowSet != 0x00
}

func (c Cpu) IsSigning() bool {
    return c.flags & signingSet != 0x00
}

func (c Cpu) IsHalt() bool {
    return c.flags & haltSet != 0x00
}

/*
void Cpu::pushState(Stack &stack)
{
   stack.push(getRegister0());
   stack.push(getRegister1());
   stack.push(getProgramCounter());
   stack.push(getHeapPointer());
//   stack.push(flags);
}

void Cpu::popState(Stack &stack)
{
//   flags = stack.pop();
   setHeapPointer(stack.pop());
   setProgramCounter(stack.pop());
   setRegister1(stack.pop());
   setRegister0(stack.pop());
}
*/
