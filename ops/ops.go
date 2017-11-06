package ops

import (
    "fmt"
    "github.com/jmacIndy/cpugo/cpu"
    "github.com/jmacIndy/cpugo/memory"
)

// FUNCTION: opHalt (HALT op code 0x00)
func opHalt(c *cpu.Cpu) {
    c.SetHalt()
    c.IncrementProgramCounter()
}

// FUNCTION: opSet0 (SET0 op code 0x01)
func opSet0(c *cpu.Cpu, m memory.Memory) {
    c.IncrementProgramCounter()
    c.SetRegister0(m.Read(c.ProgramCounter()))
    c.IncrementProgramCounter()
}

// FUNCTION: opSet1 (SET1 op code 0x02)
func opSet1(c *cpu.Cpu, m memory.Memory) {
    c.IncrementProgramCounter()
    c.SetRegister1(m.Read(c.ProgramCounter()))
    c.IncrementProgramCounter()
}

/*
// FUNCTION: opAdd (ADD op code 0x03)
void opAdd(Cpu &cpu)
{
   cpu.setRegister0(cpu.getRegister0() + cpu.getRegister1());
   cpu.incrementProgramCounter();
}

// FUNCTION: opStore (STOR op code 0x04)
void opStore(Cpu &cpu, Memory &memory, Heap &heap)
{
   cpu.incrementProgramCounter();
   cpu.setHeapPointer(memory.read(cpu.getProgramCounter()));
   heap.write(cpu.getHeapPointer(), cpu.getRegister0());
   cpu.incrementProgramCounter();
}

// FUNCTION: opPrint (PRT op code 0x05)
void opPrint(Cpu &cpu, Memory &memory, Heap &heap)
{
   cpu.incrementProgramCounter();
   cpu.setHeapPointer(memory.read(cpu.getProgramCounter()));
   std::cout << "HEAP ADDRESS: "
             << cpu.getHeapPointer()
             << " VALUE: "
             << heap.read(cpu.getHeapPointer())
             << std::endl;
   cpu.incrementProgramCounter();
}

// FUNCTION: opMultiply (MULT op code 0x07)
void opMultiply(Cpu &cpu)
{
   cpu.setRegister0(cpu.getRegister0() * cpu.getRegister1());
   cpu.incrementProgramCounter();
}

// FUNCTION; opDivide (DIV op code 0x08)
void opDivide(Cpu &cpu)
{
   cpu.setRegister0(cpu.getRegister0() / cpu.getRegister1());
   cpu.incrementProgramCounter();
}

// FUNCTION: opSubtract (SUB op code 0x09)
void opSubtract(Cpu &cpu)
{
   cpu.setRegister0(cpu.getRegister0() - cpu.getRegister1());
   cpu.incrementProgramCounter();
}

// FUNCTION: opJumpEqual (JEQ op code 0x0A)
void opJumpEqual(Cpu &cpu, Memory &memory)
{
   cpu.incrementProgramCounter();
   if (cpu.isGreaterThan() || cpu.isLessThan())
   {
      cpu.incrementProgramCounter();
   }
   else
   {
      cpu.setProgramCounter(memory.read(cpu.getProgramCounter()));
   }
}

// FUNCTION: opJumpNotEqual (JNE op code 0x0B)
void opJumpNotEqual(Cpu &cpu, Memory &memory)
{
   cpu.incrementProgramCounter();
   if (cpu.isGreaterThan() || cpu.isLessThan())
   {
      cpu.setProgramCounter(memory.read(cpu.getProgramCounter()));
   }
   else
   {
      cpu.incrementProgramCounter();
   }
}

// FUNCTION: opJumpLessThan (JLT op code 0x0C)
void opJumpLessThan(Cpu &cpu, Memory &memory)
{
   cpu.incrementProgramCounter();
   if (cpu.isLessThan())
   {
      cpu.setProgramCounter(memory.read(cpu.getProgramCounter()));
   }
   else
   {
      cpu.incrementProgramCounter();
   }
}

// FUNCTION: opJumpGreaterThan (JGT op code 0x0D)
void opJumpGreaterThan(Cpu &cpu, Memory &memory)
{
   cpu.incrementProgramCounter();
   if (cpu.isGreaterThan())
   {
      cpu.setProgramCounter(memory.read(cpu.getProgramCounter()));
   }
   else
   {
      cpu.incrementProgramCounter();
   }
}

// FUNCTION: opCall (CALL op code 0x0E)
void opCall(Cpu &cpu, Stack &stack, Memory &memory)
{
   cpu.incrementProgramCounter();
   stack.push(cpu.getProgramCounter() + 1);
   cpu.setProgramCounter(memory.read(cpu.getProgramCounter()));
}

// FUNCTION: opReturn (RET op code 0x0F)
void opReturn(Cpu &cpu, Stack &stack)
{
   cpu.setProgramCounter(stack.pop());
}

// FUNCTION: opInterrupt (INT op code 0x10)
void opInterrupt(Cpu &cpu, Memory &memory, Stack &stack) 
{

   // interrupt number is operand
   //    type goes into r0
   // interrupt 0x01 - keyboard services
   //    type 01 - read from keyboard
   //            - read value goes into r1
   // interrupt 0x02 - display services
   //    type 01 - write number to display
   //            - r1 contains number to write
   //    type 02 - write character to display
   //            - r1 contains character to write
   //    type 03 - write string to display
   //            - r1 contains address of start of string

   cpu.incrementProgramCounter();
   byte interruptCode = memory.read(cpu.getProgramCounter());

   cpu.pushState(stack);
   handleInterrupt(interruptCode, cpu.getRegister0(), cpu.getRegister1());
   cpu.popState(stack);
   cpu.incrementProgramCounter();
}

// FUNCTION: opCompare (CMP op code 0x11)
void opCompare(Cpu &cpu)
{
   cpu.resetGreaterThan();
   cpu.resetLessThan();

   if (cpu.getRegister0() > cpu.getRegister1())
   {
      cpu.setGreaterThan();
   }
   else if (cpu.getRegister0() < cpu.getRegister1())
   {
      cpu.setLessThan();
   }
}

// FUNCTION: opJumpNotZero (JNZ op code 0x12)
void opJumpNotZero(Cpu &cpu, Memory &memory)
{
   cpu.incrementProgramCounter();
   if (cpu.isNonZero())
   {
      cpu.setProgramCounter(memory.read(cpu.getProgramCounter()));
   }
   else
   {
      cpu.incrementProgramCounter();
   }
}

// FUNCTION: opJump (JMP op code 0x13)
void opJump(Cpu &cpu, Memory &memory)
{
   cpu.incrementProgramCounter();
   cpu.setProgramCounter(memory.read(cpu.getProgramCounter()));
}

// FUNCTION: opJumpZero (JZ opcode 0x14)
void opJumpZero(Cpu &cpu, Memory &memory)
{
   cpu.incrementProgramCounter();
   if (cpu.isZero())
   {
      cpu.setProgramCounter(memory.read(cpu.getProgramCounter()));
   }
   else
   {
      cpu.incrementProgramCounter();
   }
}

// FUNCTION: opIncrement0 (INC0 opcode 0x15)
void opIncrement0(Cpu &cpu)
{
   cpu.setRegister0(cpu.getRegister0() + 1);
   cpu.incrementProgramCounter();
}

// FUNCTION: opIncrement1 (INC1 opcode 0x16)
void opIncrement1(Cpu &cpu)
{
   cpu.setRegister1(cpu.getRegister1() + 1);
   cpu.incrementProgramCounter();
}

// FUNCTION; opDecrement0 (DEC0 opcode 0x17)
void opDecrement0(Cpu &cpu)
{
   cpu.setRegister0(cpu.getRegister0() - 1);
   cpu.incrementProgramCounter();
}

// FUNCTION: opDecrement1 (DEC1 opcode 0x18)
void opDecrement1(Cpu &cpu)
{
   cpu.setRegister1(cpu.getRegister1() - 1);
   cpu.incrementProgramCounter();
}

// FUNCTION: opLoad0 (LD0 opcode 0x19)
void opLoad0(Cpu &cpu, Heap &heap)
{
   cpu.incrementProgramCounter();
   cpu.setRegister0(heap.read(cpu.getProgramCounter()));
   cpu.incrementProgramCounter();
}

// FUNCTION: opLoad1 (LD1 opcode 0x1A)
void opLoad1(Cpu &cpu, Heap &heap)
{
   cpu.incrementProgramCounter();
   cpu.setRegister1(heap.read(cpu.getProgramCounter()));
   cpu.incrementProgramCounter();
}

// FUNCTION: opTest (TST opcode 0x1B)
void opTest(Cpu &cpu)
{
   cpu.resetZero();
   cpu.resetNonZero();

   if (cpu.getRegister0() == 0x00)
   {
      cpu.setZero();
   }
   else
   {
      cpu.setNonZero();
   }

   cpu.incrementProgramCounter();
}
*/

// FUNCTION: run
func Run(c *cpu.Cpu, m memory.Memory) {
    fmt.Println("... Running ...")

    for {
      switch m.Read(c.ProgramCounter()) {
      case 0x00: opHalt(c)
      case 0x01: opSet0(c, m)
      case 0x02: opSet1(c, m)
/*
      case 0x03:
         opAdd(cpu);
         break;
      case 0x04:
         opStore(cpu, memory, heap);
         break;
      case 0x05:
         opPrint(cpu, memory, heap);
         break;
      case 0x07:
         opMultiply(cpu);
         break;
      case 0x08:
         opDivide(cpu);
         break;
      case 0x09:
         opSubtract(cpu);
         break;
      case 0x0A:
         opJumpEqual(cpu, memory);
         break;
      case 0x0B:
         opJumpNotEqual(cpu, memory);
         break;
      case 0x0C:
         opJumpLessThan(cpu, memory);
         break;
      case 0x0D:
         opJumpGreaterThan(cpu, memory);
         break;
      case 0x0E:
         opCall(cpu, stack, memory);
         break;
      case 0x0F:
         opReturn(cpu, stack);
         break;
      case 0x10:
         opInterrupt(cpu, memory, stack);
         break;
      case 0x11:
         opCompare(cpu);
         break;
      case 0x12:
         opJumpNotZero(cpu, memory);
         break;
      case 0x13:
         opJump(cpu, memory);
         break;
      case 0x14:
         opJumpZero(cpu, memory);
         break;
      case 0x15:
         opIncrement0(cpu);
         break;
      case 0x16:
         opIncrement1(cpu);
         break;
      case 0x17:
         opDecrement0(cpu);
         break;
      case 0x18:
         opDecrement1(cpu);
         break;
      case 0x19:
         opLoad0(cpu, heap);
         break;
      case 0x1A:
         opLoad1(cpu, heap);
         break;
      case 0x1B:
         opTest(cpu);
         break;
*/
      default:
         fmt.Println("ERROR: Bad OpCode")
      }

      if (c.IsHalt()) {
         break;
      }
   }
}
