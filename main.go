package main

import (
    "fmt"
    "github.com/jmacIndy/cpugo/cpu"
    "github.com/jmacIndy/cpugo/memory"
    "github.com/jmacIndy/cpugo/heap"
    "github.com/jmacIndy/cpugo/stack"
    "github.com/jmacIndy/cpugo/ops"
)

/*
#define MAX_FILENAME_SIZE 50
#define MAX_INPUT_SIZE 500
*/

func displayMenu() int {

    fmt.Println("Menu (my CPU)")
    fmt.Println("-------------")
    fmt.Println("1. Reset CPU")
    fmt.Println("2. Clear Memory")
    fmt.Println("3. Dump the CPU")
    fmt.Println("4. Dump the Memory")
    fmt.Println("5. Dump the Heap")
    fmt.Println("6. Dump the Stack")
    fmt.Println("7. Run the CPU")
    fmt.Println("8. Load program from file")
    fmt.Println("9. Exit the CPU")
    fmt.Print("Your choice ===> ")

    var choice int
    _, err := fmt.Scanf("%d", &choice)

    if err != nil {
        choice = 0
    }

    return choice
}

/*
void loadProgram(Memory &memory)
{
   char inName[MAX_FILENAME_SIZE];

   std::cout << std::endl
             << "Enter input file name ===> ";
   std::cin >> inName;

   std::ifstream inFile(inName);
   std::string inputLine((std::istreambuf_iterator<char>(inFile)),
      (std::istreambuf_iterator<char>() ));

   int filePointer = 3; // skip over CPU text
   byte memoryPointer = 0x00;
   for (std::string::size_type i = filePointer; i < inputLine.size(); i += 2)
   {
      std::string inData = "0x";
      inData.push_back(inputLine[i]);
      inData.push_back(inputLine[i+1]);
      memory.write(memoryPointer, std::stoul(inData, nullptr, 16));
      memoryPointer++;
   }
}
*/

func load_program() {
}

func main() {

    fmt.Println("=========================")
    fmt.Println("=== Welcome to My CPU ===")
    fmt.Println("=========================")

    exitFlag := false

    var c cpu.Cpu

    var m memory.Memory

    var h heap.Heap

    var s stack.Stack

    m.Write(0x00, 0x01) // SET0 5
    m.Write(0x01, 0x05)
    m.Write(0x02, 0x00) // HALT

    for !exitFlag  {
        switch displayMenu() {
            case 1: c.ReInitialize()
            case 2: m.ReInitialize()
            case 3: c.Dump()
            case 4: m.Dump()
            case 5: h.Dump()
            case 6: s.Dump()
            case 7:
                if c.IsHalt() {
                    fmt.Println("ERROR: CPU is Halted");
                } else {
                    ops.Run(&c, m)
                }
            case 8: load_program()
            case 9: exitFlag = true
            default:
                fmt.Println("\nERROR: Invalid input");

        }
    }
}
