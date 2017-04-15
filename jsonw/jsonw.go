package jsonw

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/robjporter/go-functions/as"
)

type Wrapper struct {
	dat    interface{}
	err    *Error
	access []string
}

type Error struct {
	msg string
}

/*==========================================================================
TYPE
==========================================================================*/
func (w *Wrapper) TypeOf() string {
	return fmt.Sprintf("%T", w.dat)
}

/*==========================================================================
CLEAR
==========================================================================*/
func (w *Wrapper) Clear() {
	w.dat = nil
}

/*==========================================================================
COPY
==========================================================================*/
func (w *Wrapper) Copy() (*Wrapper, error) {
	ret := *w
	return &ret, nil
}

/*==========================================================================
ERROR
==========================================================================*/
func (e Error) Error() string { return e.msg }

func (w *Wrapper) Error() (e error) {
	if w.err != nil {
		e = *w.err
	}
	return
}

/*==========================================================================
LENGTH
==========================================================================*/
func (w *Wrapper) Len() (ret int, err error) {
	tmp, v := w.asArray()
	if v == nil {
		err = tmp.err
	} else {
		ret = len(v)
	}
	return
}

func (w *Wrapper) Size() (ret int, err error) {
	v, err := w.Keys()
	if v == nil {
		ret = 0
	} else {
		ret = len(v)
	}
	return
}

/*==========================================================================
KEYS
==========================================================================*/
func (w *Wrapper) Keys() (v []string, err error) {
	tmp, d := w.asDictionary()
	if d == nil {
		err = tmp.err
	} else {
		v = make([]string, len(d))
		var i int = 0
		for k, _ := range d {
			v[i] = k
			i++
		}
	}
	return
}

/*==========================================================================
IS
==========================================================================*/
func (w *Wrapper) IsOk() bool { return w.Error() == nil }
func (w *Wrapper) IsInt() bool {
	if as.IsInt(w.dat) {
		return true
	}
	return false
}

func (w *Wrapper) IsString() bool {
	if as.IsString(w.dat) {
		return true
	}
	return false
}

func (w *Wrapper) IsFloat() bool {
	if as.IsFloat(w.dat) {
		return true
	}
	return false
}

func (w *Wrapper) IsBool() bool {
	if as.IsBool(w.dat) {
		return true
	}
	return false
}

func (w *Wrapper) IsExist(path string) bool {
	bits := strings.Split(path, ".")
	ret := w
	for _, bit := range bits {
		if len(bit) > 0 && (bit[0] >= '0' && bit[0] <= '9') {
			// this is probably an int, use AtIndex instead
			if val, e := strconv.Atoi(bit); e == nil {
				ret = ret.AtIndex(val)
			} else {
				ret = ret.AtKey(bit)
			}
		} else if len(bit) > 0 {
			ret = ret.AtKey(bit)
		} else {
			break
		}

		if ret.dat == nil || ret.err != nil {
			break
		}
	}
	if ret.dat != nil {
		return true
	}
	return false
}

func (w *Wrapper) IsEmpty() bool {
	length, _ := w.Len()
	size, _ := w.Size()
	if length == 0 && size == 0 {
		return true
	}
	return false
}

/*==========================================================================
GET
==========================================================================*/
func (w *Wrapper) GetJsonString() (string, error) {
	data, err := json.Marshal(w.getData())
	return as.ToString(data), err
}

func (w *Wrapper) GetJsonPrettyString() (string, error) {
	data, err := json.MarshalIndent(w.getData(), "", "    ")
	return as.ToString(data), err
}

func (w *Wrapper) getData() interface{} { return w.dat }

func (w *Wrapper) GetDataOrNil() interface{} { return w.getData() }

func (w *Wrapper) GetInt() int {
	return as.ToInt(w.dat)
}

func (w *Wrapper) GetInt64() int64 {
	return as.ToInt64(w.dat)
}

func (w *Wrapper) GetString() string {
	return as.ToString(w.dat)
}

func (w *Wrapper) GetBool() bool {
	return as.ToBool(w.dat)
}

func (w *Wrapper) GetFloat() float64 {
	return as.ToFloat(w.dat)
}

/*==========================================================================
NEW
==========================================================================*/
func NewWrapper(i interface{}) (rd *Wrapper) {
	rd = new(Wrapper)
	rd.dat = i
	rd.access = make([]string, 1, 1)
	rd.access[0] = "<root>"
	return rd
}

func NewArray(l int) *Wrapper {
	m := make([]interface{}, l)
	return NewWrapper(m)
}

func NewDictionary() *Wrapper {
	m := make(map[string]interface{})
	return NewWrapper(m)
}

func NewDictionaryFromString(s string) *Wrapper {
	var data interface{}
	json.Unmarshal([]byte(s), &data)
	return NewWrapper(data)
}

func NewString(s string) *Wrapper {
	return NewWrapper(s)
}

func NewInt(i int) *Wrapper {
	return NewWrapper(i)
}

func NewBool(b bool) *Wrapper {
	return NewWrapper(b)
}

func NewFloat(f float64) *Wrapper {
	return NewWrapper(f)
}

func NewNil() *Wrapper {
	return NewWrapper(nil)
}

func NewInt64(i int64) *Wrapper {
	return NewWrapper(i)
}

/*==========================================================================
SET
==========================================================================*/
func (w *Wrapper) SetKey(s string, val *Wrapper) error {
	d, ok := (w.dat).(map[string]interface{})
	if ok {
		if d != nil {
			d[s] = val.getData()
		}
	}
	return w.Error()
}

func (w *Wrapper) SetIndex(i int, val *Wrapper) error {
	b, d := w.asArray()
	if d != nil {
		d[i] = val.getData()
	}
	return b.Error()

}

/*==========================================================================
AT
==========================================================================*/
func (w *Wrapper) AtKey(s string) *Wrapper {
	ret, d := w.asDictionary()

	if d != nil {
		val, found := d[s]
		if found {
			ret.dat = val
		} else {
			ret.dat = nil
		}
	}
	ret.access = append(ret.access, fmt.Sprintf(".%s", s))
	return ret
}

func (w *Wrapper) AtIndex(i int) *Wrapper {
	ret, v := w.asArray()
	if v == nil {

	} else if len(v) <= i {
		data := fmt.Sprintf("index out of bounds %d >= %d", i, len(v))
		ret.err = &Error{msg: data}
	} else {
		ret.dat = v[i]
	}
	ret.access = append(ret.access, fmt.Sprintf("[%d]", i))
	return ret
}

func (w *Wrapper) AtPath(path string) (ret *Wrapper) {
	bits := strings.Split(path, ".")
	ret = w
	for _, bit := range bits {
		if len(bit) > 0 && (bit[0] >= '0' && bit[0] <= '9') {
			// this is probably an int, use AtIndex instead
			if val, e := strconv.Atoi(bit); e == nil {
				ret = ret.AtIndex(val)
			} else {
				ret = ret.AtKey(bit)
			}
		} else if len(bit) > 0 {
			ret = ret.AtKey(bit)
		} else {
			break
		}

		if ret.dat == nil || ret.err != nil {
			break
		}
	}
	return ret
}

/*==========================================================================
AS
==========================================================================*/
func (w *Wrapper) asArray() (ret *Wrapper, v []interface{}) {
	if w.err != nil {
		ret = w
	} else {
		var ok bool
		v, ok = (w.dat).([]interface{})
		ret = new(Wrapper)
		ret.access = w.access
		if !ok {
			ret.err = &Error{msg: "NOT AN ARRAY"}
		}
	}
	return
}

func (w *Wrapper) asDictionary() (ret *Wrapper, d map[string]interface{}) {
	if w.err != nil {
		ret = w
	} else {
		var ok bool
		d, ok = (w.dat).(map[string]interface{})
		ret = new(Wrapper)
		ret.access = w.access
		if !ok {
			ret.err = &Error{msg: "NOT A DICTIONARY"}
		}
	}
	return
}

/*==========================================================================
TO
==========================================================================*/
func (w *Wrapper) ToDictionary() (out *Wrapper, e error) {
	tmp, _ := w.asDictionary()
	if tmp.err != nil {
		e = tmp.err
	} else {
		out = w
	}
	return
}

func (w *Wrapper) ToArray() (out *Wrapper, e error) {
	tmp, _ := w.asArray()
	if tmp.err != nil {
		e = tmp.err
	} else {
		out = w
	}
	return
}

func (w *Wrapper) ToIntArray() ([]int, error) {
	var out []int
	var e error
	tmp, _ := w.asArray()
	if tmp.err != nil {
		e = tmp.err
	} else {
		tmp2 := as.ToSlice(w.dat)
		for i := 0; i < len(tmp2); i++ {
			out = append(out, as.ToInt(tmp2[i]))
		}
	}
	return out, e
}

func (w *Wrapper) ToStringArray() ([]string, error) {
	var out []string
	var e error
	tmp, _ := w.asArray()
	if tmp.err != nil {
		e = tmp.err
	} else {
		tmp2 := as.ToSlice(w.dat)
		for i := 0; i < len(tmp2); i++ {
			out = append(out, as.ToString(tmp2[i]))
		}
	}
	return out, e
}
