package main

import (
	"fmt"

	"../jsonw"
)

func main() {
	w := jsonw.NewDictionary()
	if w.IsEmpty() {
		fmt.Println("EMPTY")
	}

	w.SetKey("dos", jsonw.NewString("deux"))
	w.SetKey("tres", jsonw.NewString("trois"))
	w.SetKey("quatro", jsonw.NewInt(4))
	w.SetKey("others", jsonw.NewArray(3))

	w.AtKey("others").SetIndex(0, jsonw.NewInt(100))
	w.AtKey("others").SetIndex(1, jsonw.NewInt(101))
	w.AtKey("others").SetIndex(2, jsonw.NewInt(102))

	number := 20
	w.SetKey("Numbers", jsonw.NewArray(number))
	for i := 0; i < number; i++ {
		w.AtKey("Numbers").SetIndex(i, jsonw.NewInt(i))
	}

	w.SetKey("Tests", jsonw.NewArray(2))
	w.AtKey("Tests").SetIndex(0, jsonw.NewDictionary())
	w.AtKey("Tests").AtIndex(0).SetKey("String", jsonw.NewString("string"))
	w.AtKey("Tests").AtIndex(0).SetKey("Bool", jsonw.NewBool(true))
	w.AtKey("Tests").AtIndex(0).SetKey("Int", jsonw.NewInt(8))
	w.AtKey("Tests").AtIndex(0).SetKey("Float", jsonw.NewFloat(44.4422))

	w.AtKey("Tests").SetIndex(1, jsonw.NewDictionary())
	w.AtKey("Tests").AtIndex(1).SetKey("String", jsonw.NewString("string"))
	w.AtKey("Tests").AtIndex(1).SetKey("Bool", jsonw.NewBool(true))
	w.AtKey("Tests").AtIndex(1).SetKey("Int", jsonw.NewInt(8))
	w.AtKey("Tests").AtIndex(1).SetKey("Float", jsonw.NewFloat(44.4422))

	w.SetKey("Depth", jsonw.NewDictionary())
	w.AtKey("Depth").SetKey("Depth2A", jsonw.NewString("DEPTHA"))
	w.AtKey("Depth").SetKey("Depth2B", jsonw.NewString("DEPTHB"))
	w.AtKey("Depth").SetKey("Depth2C", jsonw.NewDictionary())
	w.AtKey("Depth").AtKey("Depth2C").SetKey("Depth3", jsonw.NewString("DEPTHA"))

	fmt.Println("**OUTPUT********************")
	fmt.Println(w.GetJsonString())
	fmt.Println(w.GetJsonPrettyString())

	fmt.Println("**PATH********************")
	fmt.Println("PATH 1: >", w.AtPath("Tests.0.String").GetString())
	fmt.Println("PATH 2: >", w.AtPath("Tests.4.String").GetString())
	fmt.Println("PATH 3: >", w.AtPath("Numbers.2").GetInt())
	fmt.Println("PATH 4: >", w.AtPath("Numbers.16").GetString())

	fmt.Println("**ISEXIST********************")
	fmt.Println("IS PATH 1: >", w.IsExist("Numbers.2"))
	fmt.Println("IS PATH 2: >", w.IsExist("Numbers.26"))

	fmt.Println("**IS********************")
	if w.AtKey("Tests").AtIndex(0).AtKey("Int").IsInt() {
		fmt.Println("INT:    > TRUE")
	} else {
		fmt.Println("INT:    > FALSE")
	}
	if w.AtKey("Tests").AtIndex(0).AtKey("String").IsString() {
		fmt.Println("STRING: > TRUE")
	} else {
		fmt.Println("STRING: > FALSE")
	}
	if w.AtKey("Tests").AtIndex(0).AtKey("Bool").IsBool() {
		fmt.Println("BOOL:   > TRUE")
	} else {
		fmt.Println("BOOL:   > FALSE")
	}
	if w.AtKey("Tests").AtIndex(0).AtKey("Float").IsFloat() {
		fmt.Println("FLOAT:  > TRUE")
	} else {
		fmt.Println("FLOAT:  > FALSE")
	}

	fmt.Println("**GET********************")
	fmt.Println("INT:    >", w.AtKey("Tests").AtIndex(0).AtKey("Int").GetInt())
	fmt.Println("STRING: >", w.AtKey("Tests").AtIndex(0).AtKey("Int").GetString())
	fmt.Println("BOOL:   >", w.AtKey("Tests").AtIndex(0).AtKey("Int").GetBool())
	fmt.Println("FLOAT:  >", w.AtKey("Tests").AtIndex(0).AtKey("Int").GetFloat())

	fmt.Println("**LENGTH********************")
	length, _ := w.AtKey("Tests").Len()
	fmt.Println("LENGTH 1: >", length)
	for i := 0; i < length; i++ {
		fmt.Println(w.AtKey("Tests").AtIndex(i))
	}
	length, _ = w.Size()
	fmt.Println("SIZE 1: >", length)
	length, _ = w.AtKey("Tests").AtIndex(0).Size()
	fmt.Println("SIZE 2: >", length)

	fmt.Println("**GET********************")
	if w.AtKey("Depth").AtKey("Depth2A").GetString() == "" {
		fmt.Println("Dictionary fail for birds.sparrow:")
	}
	if w.AtKey("Depth").AtKey("Depth2B").GetInt() == 0 {
		fmt.Println("Dictionary fail for birds.sparrow:")
	}

	fmt.Println("**COPY********************")
	v, _ := w.Copy()

	fmt.Println("**KEYS********************")
	keys, _ := v.Keys()
	fmt.Println("KEYS: >", keys)
	keys, _ = v.AtKey("Tests").AtIndex(0).Keys()
	fmt.Println("KEYS: >", keys)
	for i := 0; i < len(keys); i++ {
		typef := w.AtKey("Tests").AtIndex(0).AtKey(keys[i]).TypeOf()
		if typef == "string" {
			fmt.Println(w.AtKey("Tests").AtIndex(0).AtKey(keys[i]).GetString())
		} else if typef == "int" || typef == "int16" || typef == "int32" || typef == "int64" {
			fmt.Println(w.AtKey("Tests").AtIndex(0).AtKey(keys[i]).GetInt())
		} else if typef == "float32" || typef == "float64" {
			fmt.Println(w.AtKey("Tests").AtIndex(0).AtKey(keys[i]).GetFloat())
		} else if typef == "bool" {
			fmt.Println(w.AtKey("Tests").AtIndex(0).AtKey(keys[i]).GetBool())
		}
	}

	fmt.Println("**DICT********************")
	dict, _ := v.ToDictionary()
	fmt.Println("DICT: >", dict)

	fmt.Println("**ARRAY*******************")
	arr, _ := v.AtKey("Tests").ToArray()
	fmt.Println("ARRAY: >", arr)
	length, _ = arr.Len()
	fmt.Println("ARRAY LENGTH: >", length)
	arr, _ = v.AtKey("Numbers").ToArray()
	fmt.Println("ARRAY: >", arr)
	length, _ = arr.Len()
	fmt.Println("ARRAY LENGTH: >", length)
	fmt.Println("ARRAY ELEMENT: >", arr.AtIndex(0))
	arr2, _ := v.AtKey("Numbers").ToIntArray()
	fmt.Println("ARRAY LENGTH: >", length)
	fmt.Println("ARRAY ELEMENT: >", arr2[4])

	json := `{"Depth":{"Depth2A":"DEPTHA","Depth2B":"DEPTHB","Depth2C":{"Depth3":"DEPTHA"}},"Numbers":[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19],"Tests":[{"Bool":true,"Float":44.4422,"Int":8,"String":"string"},{"Bool":true,"Float":44.4422,"Int":8,"String":"string"}],"dos":"deux","others":[100,101,102],"quatro":4,"tres":"trois"}`
	x := jsonw.NewDictionaryFromString(json)
	fmt.Println("STRING: >", x.AtKey("Tests").AtIndex(0).AtKey("Int").GetString())
	fmt.Println("**CLEAR*******************")
	x.Clear()
	fmt.Println("STRING: >", x.AtKey("Tests").AtIndex(0).AtKey("Int").GetString())
	if x.IsEmpty() {
		fmt.Println("EMPTY")
	}
}
