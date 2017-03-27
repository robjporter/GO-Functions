package main

import (
	"fmt"

	"../as"
)

func main() {
	fmt.Println("TO STRING")
	fmt.Println(as.ToString(32))
	fmt.Println(as.ToString(bool(true)))
	fmt.Println(as.ToString("mayonegg"))         // "mayonegg"
	fmt.Println(as.ToString(8))                  // "8"
	fmt.Println(as.ToString(8.31))               // "8.31"
	fmt.Println(as.ToString([]byte("one time"))) // "one time"
	fmt.Println(as.ToString(nil))                // ""
	var foo interface{} = "one more time"
	fmt.Println(as.ToString(foo)) // "one more time"

	fmt.Println("\nTO FLOAT")
	fmt.Println(as.ToFloat(32.4400))

	fmt.Println("\nTO RUNE LENGTH")
	fmt.Println(as.ToRuneLength("test"))
	fmt.Println(as.ToRuneLength("TEST"))
	fmt.Println(as.ToRuneLength("iiii"))
	fmt.Println(as.ToRuneLength("QQKK"))
	fmt.Println(as.ToRuneLength("Lllm"))

	fmt.Println("\nTO BOOL")
	fmt.Println(as.ToBool(1))
	fmt.Println(as.ToBool(0))
	fmt.Println(as.ToBool("1"))
	fmt.Println(as.ToBool("true"))
	fmt.Println(as.ToBool("down"))

	fmt.Println("\nTO BYTES")
	fmt.Println(as.ToBytes("Testing"))

	fmt.Println("\nTO SLICE")
	var foo2 []interface{}
	foo2 = append(foo2, "one") //more time"
	fmt.Println(as.ToSlice(foo2))

	fmt.Println("\nCONVERT")
	fmt.Println(as.ToBytes("Testing"))

	fmt.Println("\nTO INT")
	fmt.Println(as.ToInt("1"))

	fmt.Println("\nTO IP")
	fmt.Println(as.ToIP("192.168.0.1")) // "one more time"

	fmt.Println("\nTO BASE64")
	fmt.Println(as.ToBase64("This is a test"))

	fmt.Println("\nFROM BASE64")
	fmt.Println(as.FromBase64("VGhpcyBpcyBhIHRlc3Q="))

	fmt.Println("\nIS EMPTY")
	fmt.Println(as.IsEmpty(0))
	fmt.Println(as.IsEmpty(""))
	fmt.Println(as.IsEmpty([]string{}))

	fmt.Println("\nIS KIND")
	fmt.Println(as.IsKind("string", 0))
	fmt.Println(as.IsKind("string", ""))

	fmt.Println("\nOF KIND")
	fmt.Println(as.OfKind("string"))
	fmt.Println(as.OfKind([]string{}))
	fmt.Println(as.OfKind(nil))
	fmt.Println(as.OfKind([]byte("one time")))
	fmt.Println(as.OfKind(bool(true)))
	fmt.Println(as.OfKind(32))

	fmt.Println("\nOF TYPE")
	fmt.Println(as.OfType(32))
	fmt.Println(as.OfType(""))
}
