package stack

import (
    "fmt"
)

const STACK_SIZE int = 256

type Stack struct {
    stack [STACK_SIZE]uint8
    stackPointer uint8
}

func (s *Stack) ReInitialize() {
    for i:= 0; i < len(s.stack); i++ {
        s.stack[i] = 0
    }
    s.stackPointer = 0;
}

func (s *Stack) Push(value uint8) {
    s.stack[s.stackPointer] = value
    s.stackPointer++
}

func (s *Stack) Pop() uint8 {
    s.stackPointer--
    return s.stack[s.stackPointer]
}

func (s Stack) Dump() {
    fmt.Println("Stack Contents:")
    fmt.Println("---------------")
    fmt.Printf("   Stack Pointer: %02X\n", s.stackPointer)

    j := 0

    addressCounter := 0;

    for i := 0; i < len(s.stack); i++ {
        if j % 16 == 0 {
            fmt.Printf("\nAddress(%02X) ", addressCounter)
            addressCounter += 16
        }
        fmt.Printf("%02X ", s.stack[i])
        j++
    }

    fmt.Println()
}
