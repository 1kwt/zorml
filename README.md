# ZORML Quick Guide - Version 2.3 (Dec 2024)

## Basic Syntax
ZORML is a stack-based programming language where each command is a single character. Commands manipulate the **stack**, execute functions, or interact with the user. Here's a breakdown of available commands:

---

## Core Commands

| Symbol  | Description                                                                                   | Example                      |
|---------|-----------------------------------------------------------------------------------------------|------------------------------|
| `>`     | Pushes the next character onto the stack.                                                     | `>A^` → Stack: `A` → Prints `A`.      |
| `<`     | Prompts the user for input and pushes the entered value onto the stack.                       | `<^` → Input: `5` → Prints `5`.      |
| `^`     | Prints the entire stack as a single string.                                                   | `>A>B^` → Prints: `AB`.      |
| `!`     | Prints the value at a specific stack index.                                                   | `>1>2!0` → Prints: `1`.      |
| `#`     | Clears the entire stack.                                                                      | `>A#^` → Prints nothing.     |
| `[ ]`   | Adds a string or long sequence of characters to the stack.                                    | `[hello]^` → Prints: `hello`.|
| `{ }`   | Defines a function to be executed later.                                                      | `{>A^}` → No output yet.     |
| `~`     | Executes the most recently defined function.                                                  | `{>A^}~` → Prints: `A`.      |
| `;`     | Stops the interpreter from processing further commands.                                       | `>1>2;` → Stops at `;`.      |

---

## Mathematical Operations

| Symbol  | Description                                                                                   | Example                      |
|---------|-----------------------------------------------------------------------------------------------|------------------------------|
| `+`     | Adds the top two items of the stack (if they are numbers) or concatenates them as strings.    | `>1>2+^` → Prints: `3`.      |
| `-`     | Subtracts the second item from the top item (if they are numbers).                            | `>5>3-^` → Prints: `2`.      |
| `/`     | Divides the second item by the top item (if they are numbers).                                | `>6>2/^` → Prints: `3`.      |
| `*`     | Multiplies the top two items of the stack (if they are numbers).                              | `>3>4*^` → Prints: `12`.     |

---

## Conditional Logic

| Symbol  | Description                                                                                   | Example                      |
|---------|-----------------------------------------------------------------------------------------------|------------------------------|
| `?`     | Starts an **if condition** that checks if two stack indices are equal.                        | `>1>1?0=1[>A^]` → Prints: `A`.      |
| `{ }`   | Encloses commands to be executed if the condition is true.                                    | `?0=1{>B^}` executes if true.|

**Condition Syntax**:
- `?index1=index2[commands]`: Executes commands if values at `index1` and `index2` are equal.

---

## Examples

### Print "hello world"
```
[hello world]^;
```
### Add 2 numbers together
```
>3>5+^; // Output: 8
```
### If-else
```
>5>5?0=1{>A^}; // Output: A because 5 = 5
```
### Define and Execute a Function
```
{[Hello]^}~; // Output: Hello
```
## Further Experimentation

### Factorial calculator in zorml
```
// code indented and commented for clarity, what you would actually put into the interpreter is at the bottom
<             // Get user input
>1            // Push initial result value to stack (1)
{             // Define factorial function
  >0!0?0=1;   // Check if input equals 1, stop if true
  ?0!=1{      // If input is not 1:
    >0>1*#    // Multiply input by result and update stack
    >0>1-#    // Decrement input and update stack
    ~         // Call function recursively until its finished
  }
}
~             // Execute factorial function
^             // Print result
;             // Finish
---------------------------------------------------------------
// This is the grammatically correct way of putting this code into zorml:
<>1{>0!0?0=1;??0!=1{>0>1*#>0>1-#~}}~;
```
### Fibonacci Generator
```
// code indented and commented for clarity, what you would actually put into the interpreter is at the bottom
<             // Get user input for the number of terms (e.g., 5)
>0>1          // Push initial Fibonacci numbers: 0 and 1
>1            // Push term index (1, representing the first term)
{             // Define Fibonacci generator function
  >!1!0+      // Add the last two Fibonacci numbers (!1 and !0)
  ^           // Print the result
  >!1         // Push the new Fibonacci number to the stack
  !2>1+>2#    // Increment the term index and update the stack
  !2!0?0=1[   // Check if the term index equals the user input
    ~         // If not, recursively call the function
  ];          // Stop if all terms are generated
}
~;            // Execute Fibonacci generator function
---------------------------------------------------------------
// This is the grammatically correct way of putting this code into zorml:
<>0>1>1{>!1!0+^>!1!2>1+>2#!2!0?0=1{~};}~;
```
