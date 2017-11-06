package memory

import (
    "fmt"
)

const MEMORY_SIZE int = 256

type Memory [MEMORY_SIZE]uint8

func (m *Memory) ReInitialize() {
    for i:= 0; i < len(m); i++ {
        m[i] = 0
    }
}

func (m *Memory) Write(address uint8, value uint8) {
    m[address] = value
}

func (m Memory) Read(address uint8) uint8 {
    return m[address]
}

func (m Memory) Dump() {
    fmt.Println("Memory Contents:")
    fmt.Println("----------------")

    j := 0

    addressCounter := 0;

    for i := 0; i < len(m); i++ {
        if j % 16 == 0 {
            fmt.Printf("\nAddress(%02X) ", addressCounter)
            addressCounter += 16
        }
        fmt.Printf("%02X ", m[i])
        j++
    }

    fmt.Println()
}
