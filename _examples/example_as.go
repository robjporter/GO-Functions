package main

import (
	"fmt"

	"../as"
)

func main() {
	fmt.Println("TO STRING")
	fmt.Println("=========================================================")
	fmt.Println("STRING: (32)                        >", `"`+as.ToString(32)+`"`)
	fmt.Println("STRING: (true)                      >", `"`+as.ToString(bool(true))+`"`)
	fmt.Println("STRING: ('mayonegg')                >", `"`+as.ToString("mayonegg")+`"`)         // "mayonegg"
	fmt.Println("STRING: (8)                         >", `"`+as.ToString(8)+`"`)                  // "8"
	fmt.Println("STRING: (8.31)                      >", `"`+as.ToString(8.31)+`"`)               // "8.31"
	fmt.Println("STRING: ([]byte('one time'))        >", `"`+as.ToString([]byte("one time"))+`"`) // "one time"
	fmt.Println("STRING: (nil)                       >", `"`+as.ToString(nil)+`"`)                // ""
	var foo interface{} = "one more time"
	fmt.Println("STRING: (interface{'one more time}) >", `"`+as.ToString(foo)) // "one more time"

	fmt.Println("TRIMMED")
	fmt.Println("=========================================================")
	fmt.Println("TRIMMED: ('    TEST      ')         >", `"`+as.Trimmed("    TEST      ")+`"`)

	fmt.Println("\nTO FLOAT")
	fmt.Println("=========================================================")
	fmt.Println("FLOAT: (32.4400)                    >", as.ToFloat(32.4400))
	fmt.Println("FLOAT32: (32.4400)                  >", as.ToFloat32(32.4400))

	fmt.Println("\nTO RUNE LENGTH")
	fmt.Println("=========================================================")
	fmt.Println("RUNELENGTH: ('test')                >", as.ToRuneLength("test"))
	fmt.Println("RUNELENGTH: ('TEST')                >", as.ToRuneLength("TEST"))
	fmt.Println("RUNELENGTH: ('iiii')                >", as.ToRuneLength("iiii"))
	fmt.Println("RUNELENGTH: ('QQKK')                >", as.ToRuneLength("QQKK"))
	fmt.Println("RUNELENGTH: ('Lllm')                >", as.ToRuneLength("Lllm"))

	fmt.Println("\nTO BOOL")
	fmt.Println("=========================================================")
	fmt.Println("BOOL: (1)                           >", as.ToBool(1))
	fmt.Println("BOOL: (0)                           >", as.ToBool(0))
	fmt.Println("BOOL: ('1')                         >", as.ToBool("1"))
	fmt.Println("BOOL: ('true')                      >", as.ToBool("true"))
	fmt.Println("BOOL: ('down')                      >", as.ToBool("down"))

	fmt.Println("\nTO BYTES")
	fmt.Println("=========================================================")
	fmt.Println("BYTES: ('Testing')                  >", as.ToBytes("Testing"))

	fmt.Println("\nTO SLICE")
	fmt.Println("=========================================================")
	var foo2 []interface{}
	foo2 = append(foo2, "one") //more time"
	fmt.Println("SLICE: ('one')                      >", as.ToSlice(foo2))

	fmt.Println("\nTO INT")
	fmt.Println("=========================================================")
	fmt.Println("INT: ('1')                          >", as.ToInt("1"))
	fmt.Println("INT64: ('1')                        >", as.ToInt64("1"))
	fmt.Println("INT32: ('1')                        >", as.ToInt32("1"))
	fmt.Println("INT16: ('1')                        >", as.ToInt16("1"))
	fmt.Println("INT8: ('1')                         >", as.ToInt8("1"))

	fmt.Println("\nTO IP")
	fmt.Println("=========================================================")
	fmt.Println("IP ADDRESS: ('192.168.0.1')          >", as.ToIP("192.168.0.1"))   // "one more time"
	fmt.Println("IP ADDRESS: ('one more time')        >", as.ToIP("one more time")) //
	fmt.Println("IP ADDRESS: ('1')                    >", as.ToIP("1"))             // "one more time"
	fmt.Println("IP ADDRESS: ('1.0')                  >", as.ToIP("1.0"))           // "one more time"
	fmt.Println("IP ADDRESS: ('1.0.0')                >", as.ToIP("1.0.0"))         // "one more time"
	fmt.Println("IP ADDRESS: ('1.0.0.0/8')            >", as.ToIP("1.0.0.0/8"))     // "one more time"

	fmt.Println("\nTO BASE64")
	fmt.Println("=========================================================")
	fmt.Println("TOBASE64: ('This is a test')         >", as.ToBase64("This is a test"))

	fmt.Println("\nFROM BASE64")
	fmt.Println("=========================================================")
	fmt.Println("FROMBASE64: ('VGhpcyBpcyBhIHRlc3Q=') >", as.FromBase64("VGhpcyBpcyBhIHRlc3Q="))

	fmt.Println("\nIS EMPTY")
	fmt.Println("=========================================================")
	fmt.Println("IP EMPTY: ('0')                      >", as.IsEmpty(0))
	fmt.Println("IP EMPTY: ('1')                      >", as.IsEmpty(1))
	fmt.Println("IP EMPTY: ('')                       >", as.IsEmpty(""))
	fmt.Println("IP EMPTY: ('sdasdass')               >", as.IsEmpty("sdasdass"))
	fmt.Println("IP EMPTY: ('[]string{}')             >", as.IsEmpty([]string{}))

	fmt.Println("\nIS KIND")
	fmt.Println("=========================================================")
	fmt.Println("IS KIND: (string,0)                  >", as.IsKind("string", 0))
	fmt.Println("IS KIND: (string,'')                 >", as.IsKind("string", ""))
	fmt.Println("IS KIND: (int,0)                     >", as.IsKind("int", 0))
	fmt.Println("IS KIND: (int,'test')                >", as.IsKind("int", "test"))

	fmt.Println("\nOF KIND")
	fmt.Println("=========================================================")
	fmt.Println("KIND OF: ('string')                  >", as.OfKind("string"))
	fmt.Println("KIND OF: ([]string{})                >", as.OfKind([]string{}))
	fmt.Println("KIND OF: (nil)                       >", as.OfKind(nil))
	fmt.Println("KIND OF: ([]byte('one time))         >", as.OfKind([]byte("one time")))
	fmt.Println("KIND OF: (bool(true))                >", as.OfKind(bool(true)))
	fmt.Println("KIND OF: (32)                        >", as.OfKind(32))

	fmt.Println("\nOF TYPE")
	fmt.Println("=========================================================")
	fmt.Println("TYPE: (32)                           >", as.OfType(32))
	fmt.Println("TYPE: ('')                           >", as.OfType(""))
	fmt.Println("TYPE: ([]string{}])                  >", as.OfType([]string{}))
	fmt.Println("TYPE: (true)                         >", as.OfType(true))
	fmt.Println("TYPE: (1.0f)                         >", as.OfType(1.00))
	fmt.Println("TYPE: (int64(22))                    >", as.OfType(int64(22)))

	fmt.Println("\nTO TIME")
	fmt.Println("=========================================================")
	fmt.Println("TIME: ('2016-04-04')                 >", as.ToTime("2016-04-04"))
	fmt.Println("TIME: ('04-04-2016')                 >", as.ToTime("04-04-2016"))
	fmt.Println("TIME: ('2016-04-04 16:20:40')        >", as.ToTime("2016-04-04 16:20:40"))
	fmt.Println("TIME: ('2016-04-04 16:20:40 +1 BST') >", as.ToTime("2016-04-04 16:20:40 +1 BST"))

	fmt.Println("\nTO DURATION")
	fmt.Println("=========================================================")
	fmt.Println("DURATION: (1h44m)                    >", as.ToDuration("1h44m"))
	fmt.Println("DURATION: (44)                       >", as.ToDuration("44"))
	fmt.Println("DURATION: (44s)                      >", as.ToDuration("44s"))
	fmt.Println("DURATION: (444h)                     >", as.ToDuration("444h"))
	fmt.Println("DURATION: (88m)                      >", as.ToDuration("88m"))

	fmt.Println("\nTO FIXED LENGTH AFTER")
	fmt.Println("=========================================================")
	fmt.Println("FIXED LENGTH AFTER (*,20):           >", as.ToFixedLengthAfter("Test String", "*", 20))
	fmt.Println("FIXED LENGTH AFTER (-,50):           >", as.ToFixedLengthAfter("Test String", "-", 50))
	fmt.Println("FIXED LENGTH AFTER (*,10):           >", as.ToFixedLengthAfter("Test String", "*", 10))
	fmt.Println("FIXED LENGTH AFTER (*,8):            >", as.ToFixedLengthAfter("Test String", "*", 8))

	fmt.Println("\nTO FIXED LENGTH BEFORE")
	fmt.Println("=========================================================")
	fmt.Println("FIXED LENGTH BEFORE (*,20):          >", as.ToFixedLengthBefore("Test String", "*", 20))
	fmt.Println("FIXED LENGTH BEFORE (-,50):          >", as.ToFixedLengthBefore("Test String", "-", 50))
	fmt.Println("FIXED LENGTH BEFORE (*,10):          >", as.ToFixedLengthBefore("Test String", "*", 10))
	fmt.Println("FIXED LENGTH BEFORE (*,8):           >", as.ToFixedLengthBefore("Test String", "*", 8))

	fmt.Println("\nTO FIXED LENGTH CENTER")
	fmt.Println("=========================================================")
	fmt.Println("FIXED LENGTH CENTER (*,20):          >", as.ToFixedLengthCenter("Test String", "*", 20))
	fmt.Println("FIXED LENGTH CENTER (-,50):          >", as.ToFixedLengthCenter("Test String", "-", 50))
	fmt.Println("FIXED LENGTH CENTER (*,10):          >", as.ToFixedLengthCenter("Test String", "*", 10))
	fmt.Println("FIXED LENGTH CENTER (*,8):           >", as.ToFixedLengthCenter("Test String", "*", 8))

	fmt.Println("\nIS INT")
	fmt.Println("=========================================================")
	fmt.Println("INT: (44)                           >", as.IsInt(44))
	fmt.Println("INT: (true)                         >", as.IsInt(true))
	fmt.Println("INT: (44.44)                        >", as.IsInt(44.44))
	fmt.Println("INT: ('test')                       >", as.IsInt("test"))
	fmt.Println("INT: ('14:14:14')                   >", as.IsInt(as.ToTime("14:14:14")))

	fmt.Println("\nIS BOOL")
	fmt.Println("=========================================================")
	fmt.Println("BOOL: (44)                           >", as.IsBool(44))
	fmt.Println("BOOL: (true)                         >", as.IsBool(true))
	fmt.Println("BOOL: (44.44)                        >", as.IsBool(44.44))
	fmt.Println("BOOL: ('test')                       >", as.IsBool("test"))
	fmt.Println("BOOL: ('14:14:14')                   >", as.IsBool(as.ToTime("14:14:14")))

	fmt.Println("\nIS FLOAT")
	fmt.Println("=========================================================")
	fmt.Println("FLOAT: (44)                           >", as.IsFloat(44))
	fmt.Println("FLOAT: (true)                         >", as.IsFloat(true))
	fmt.Println("FLOAT: (44.44)                        >", as.IsFloat(44.44))
	fmt.Println("FLOAT: ('test')                       >", as.IsFloat("test"))
	fmt.Println("FLOAT: ('14:14:14')                   >", as.IsFloat(as.ToTime("14:14:14")))

	fmt.Println("\nIS STRING")
	fmt.Println("=========================================================")
	fmt.Println("STRING: (44)                         >", as.IsString(44))
	fmt.Println("STRING: (true)                       >", as.IsString(true))
	fmt.Println("STRING: (44.44)                      >", as.IsString(44.44))
	fmt.Println("STRING: ('test')                     >", as.IsString("test"))
	fmt.Println("STRING: ('14:14:14')                 >", as.IsString(as.ToTime("14:14:14")))

	fmt.Println("\nIS TIME")
	fmt.Println("=========================================================")
	fmt.Println("TIME: (44)                           >", as.IsTime(44))
	fmt.Println("TIME: (true)                         >", as.IsTime(true))
	fmt.Println("TIME: (44.44)                        >", as.IsTime(44.44))
	fmt.Println("TIME: ('test')                       >", as.IsTime("test"))
	fmt.Println("TIME: ('14:14:14')                   >", as.IsTime(as.ToTime("14:14:14")))

	fmt.Println("\nIS Nillable")
	fmt.Println("=========================================================")
	fmt.Println("NILLABLE: ('')                       >", as.IsNillable(""))
	fmt.Println("NILLABLE: ([]string{})               >", as.IsNillable([]string{}))

	fmt.Println("\nTO FORMATTED BYTES")
	fmt.Println("=========================================================")
	fmt.Println("FORMAT: (44)                       >", as.FormatIntToByte(44))
	fmt.Println("FORMAT: (444)                      >", as.FormatIntToByte(444))
	fmt.Println("FORMAT: (4444)                     >", as.FormatIntToByte(4444))
	fmt.Println("FORMAT: (44444)                    >", as.FormatIntToByte(44444))
	fmt.Println("FORMAT: (444444)                   >", as.FormatIntToByte(444444))
	fmt.Println("FORMAT: (4444444)                  >", as.FormatIntToByte(4444444))
	fmt.Println("FORMAT: (44444444)                 >", as.FormatIntToByte(44444444))
	fmt.Println("FORMAT: (444444444)                >", as.FormatIntToByte(444444444))
	fmt.Println("FORMAT: (4444444444)               >", as.FormatIntToByte(4444444444))
	fmt.Println("FORMAT: (44444444444)              >", as.FormatIntToByte(44444444444))
	fmt.Println("FORMAT: (444444444444)             >", as.FormatIntToByte(444444444444))
	fmt.Println("FORMAT: (4444444444444)            >", as.FormatIntToByte(4444444444444))
	fmt.Println("FORMAT: (44444444444444)           >", as.FormatIntToByte(44444444444444))
	fmt.Println("FORMAT: (444444444444444)          >", as.FormatIntToByte(444444444444444))
	fmt.Println("FORMAT: (4444444444444444)         >", as.FormatIntToByte(4444444444444444))
	fmt.Println("FORMAT: (44444444444444444)        >", as.FormatIntToByte(44444444444444444))
	fmt.Println("FORMAT: (444444444444444444)       >", as.FormatIntToByte(444444444444444444))
	fmt.Println("FORMAT: (999999999999999999)       >", as.FormatIntToByte(999999999999999999))
	fmt.Println("FORMAT: (1000000000000000000)      >", as.FormatIntToByte(1152921504606846976))

}
