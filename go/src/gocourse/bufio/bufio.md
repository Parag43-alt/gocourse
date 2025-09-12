# Go User Input with `bufio`

This guide explains how to use the `bufio` package in Go for efficient user input and output.

## Code Example

```go
package main // Go program starts here. Without 'package main', code won't run.

import (
    "bufio" // For fast input/output. Required for Reader/Writer/Scanner. // store data temporarily
    "fmt"   // For printing/output. Required for Println/Print.
    "os"    // For input/output sources (keyboard, file, screen).
)

func main() { // main() is compulsory; it's the program's entry point.

    // ---------------- READER ----------------
    // bufio.NewReader creates a "basket" to store input together.
    // Without bufio, you'd use fmt.Scan or fmt.Scanln (less flexible).
    reader := bufio.NewReader(os.Stdin) // os.Stdin = keyboard input source
    fmt.Print("Enter your name: ")      // Prompt user for input
    name, _ := reader.ReadString('\n')  // Reads input until '\n' : jab tak new linw nhi aa jati hai 
    fmt.Println("Hello,", name)         // Displays output

    // ---------------- WRITER ----------------
    // Writer is a "bag" that stores output and sends it to the screen after flushing.
    // Without bufio.Writer, fmt.Println sends output directly (no buffering).
    writer := bufio.NewWriter(os.Stdout) // os.Stdout = screen output
    writer.WriteString("This line is written using bufio.Writer\n") // Stored in buffer
    writer.Flush() // Output appears on screen only after Flush()

    // ---------------- SCANNER ----------------
    // Scanner is a "machine" that breaks input into lines/words.
    // Without bufio.Scanner, you'd need to parse strings manually.
    scanner := bufio.NewScanner(os.Stdin) // Keyboard input again
    fmt.Print("Enter your city: ")
    scanner.Scan()         // Reads input
    city := scanner.Text() // Gets input from scanner
    fmt.Println("You live in:", city) // Displays output
}
```

## Explanation

- **Reader**: Use `bufio.NewReader` for flexible input handling.
- **Writer**: Use `bufio.NewWriter` for buffered output; remember to call `Flush()`.
- **Scanner**: Use `bufio.NewScanner` for parsing input into lines or words.

Each section demonstrates how `bufio` improves input/output operations in Go compared to basic `fmt` functions.
