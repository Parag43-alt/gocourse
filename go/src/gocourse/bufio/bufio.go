package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// ------------------- SCANNER -------------------
	// ✅ Scanner hamesha banate ho: bufio.NewScanner(os.Stdin)
	// ✅ Must call: scanner.Scan() -> phir hi scanner.Text() se input milega
	// ⚡ Change kar sakte ho: scanner.Split(bufio.ScanWords) ya ScanLines (default)

	fmt.Println("=== Scanner Example ===")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter something (Scanner): ")
	scanner.Scan()          // user input wait karega
	input := scanner.Text() // ek line ya word return karega
	fmt.Println("Scanner Input ->", input)

	// Agar words alag alag chahiye:
	// scanner.Split(bufio.ScanWords)

	// ------------------- READER -------------------
	// ✅ Reader hamesha banate ho: bufio.NewReader(os.Stdin)
	// ✅ Must call: reader.ReadString('\n') -> ek line lega Enter tak
	// ⚡ Change kar sakte ho: '\n' ki jagah koi aur delimiter (jaise ',') de sakte ho

	fmt.Println("\n=== Reader Example ===")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter something (Reader): ")
	line, _ := reader.ReadString('\n') // Enter tak lega (newline bhi include karega)
	line = strings.TrimSpace(line)     // ✅ newline hata diya
	fmt.Println("Reader Input ->", line)

	// ------------------- WRITER -------------------
	// ✅ Writer hamesha banate ho: bufio.NewWriter(os.Stdout)
	// ✅ Must call: writer.Flush() warna output buffer me hi rahega
	// ⚡ Change kar sakte ho: os.Stdout ki jagah file handle pass kar do -> file me likhega

	fmt.Println("\n=== Writer Example ===")
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hello from Writer!\n") // output buffer me gaya
	writer.Flush()                             // ✅ Flush karna zaroori hai warna dikhai nahi dega

	// Agar file me likhna hai:
	// file, _ := os.Create("output.txt")
	// writer := bufio.NewWriter(file)
	// writer.WriteString("Data inside file\n")
	// writer.Flush()
}
