package heap

import (
    "fmt"
)

const HEAP_SIZE int = 256

type Heap [HEAP_SIZE]uint8

func (h *Heap) ReInitialize() {
    for i:= 0; i < len(h); i++ {
        h[i] = 0
    }
}

func (h *Heap) Write(address uint8, value uint8) {
    h[address] = value
}

func (h Heap) Read(address uint8) uint8 {
    return h[address]
}

func (h Heap) Dump() {
    fmt.Println("Heap Contents:")
    fmt.Println("--------------")

    j := 0

    addressCounter := 0;

    for i := 0; i < len(h); i++ {
        if j % 16 == 0 {
            fmt.Printf("\nAddress(%02X) ", addressCounter)
            addressCounter += 16
        }
        fmt.Printf("%02X ", h[i])
        j++
    }

    fmt.Println()
}
