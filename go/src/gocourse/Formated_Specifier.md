# 🔢 Numbers

* `%d`: integer ko decimal (**base 10**) mein print karta hai.  
   👉 Base 10 matlab 0–9 ke 10 digits use hote hain. Example: `25` (2×10 + 5)  

* `%b`: integer ko binary (**base 2**) mein print karta hai.  
   👉 Base 2 matlab sirf `0` aur `1`. Example: Decimal `5` → Binary `101`  

* `%o`: integer ko octal (**base 8**) mein print karta hai.  
   👉 Base 8 matlab 0–7 ke digits. Example: Decimal `10` → Octal `12`  

* `%x`: integer ko hexadecimal (**base 16**, chote letters) mein print karta hai.  
   👉 Base 16 matlab 0–9 aur A–F symbols. Example: Decimal `255` → Hex `ff`  

* `%X`: integer ko hexadecimal (bade letters) mein print karta hai.  

* `%c`: integer ko ASCII character ke form mein print karta hai.  

---

# ✍️ Strings

* `%s`: string ko normal form mein print karta hai.  
* `%q`: string ko double quotes ke andar print karta hai.  

---

# 💧 Floating Point

* `%f`: floating point ko decimal format mein print karta hai (default 6 decimal places).  
* `%e`: floating point ko scientific notation (chote 'e') mein print karta hai.  
* `%E`: floating point ko scientific notation (bade 'E') mein print karta hai.  
* `%g`: floating point ko short form mein print karta hai (decimal ya scientific, jo short ho).  

---

# ✅ Boolean

* `%t`: boolean value (`true` / `false`) print karta hai.  

---

# 📋 General

* `%v`: value ko default format mein print karta hai.  
* `%+v`: struct print karte waqt field names ke saath print karta hai.  
* `%#v`: value ko Go syntax ke format mein print karta hai.  
* `%T`: variable ka type print karta hai.  

---

# 📍 Pointers

* `%p`: pointer ka memory address print karta hai.  

---

# 📝 Example Code (Bases ka fark dekhne ke liye)

```go
package main

import "fmt"

func main() {
    num := 25

    fmt.Printf("Decimal (base 10): %d\n", num)  // 25
    fmt.Printf("Binary (base 2): %b\n", num)   // 11001
    fmt.Printf("Octal (base 8): %o\n", num)    // 31
    fmt.Printf("Hexadecimal (base 16, small): %x\n", num) // 19
    fmt.Printf("Hexadecimal (base 16, caps): %X\n", num)  // 19
    fmt.Printf("Character (ASCII): %c\n", num) // ASCII char for 25 (non-printable)
}
