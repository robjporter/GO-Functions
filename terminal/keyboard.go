/**
 *
 *	Keyboard IO library.
 *	For creating interesting CLI apps.
 *
 * 	Author: Timothy Lorens (tlorens@cyberdyne.org)
 * 	https://github.com/tlorens
 *	http://www.cyberdyne.org/~icebrkr
 *
 */
package terminal

/*
#include "cgetkey.h"
*/
import "C"

const HOME = 149   // DOS Keycode #0 + #71
const END = 152    // DOS Keycode #0 + #79
const PG_UP = 153  // DOS Keycode #0 + #73
const PG_DN = 154  // DOS Keycode #0 + #81
const KEY_UP = 165 // DOS Keycode #0 + #72
const KEY_DN = 166 // DOS Keycode #0 + #80
const KEY_RT = 167 // DOS Keycode #0 + #77
const KEY_LF = 168 // DOS Keycode #0 + #75

/**
 *
 *	ReadKey
 *
 *	Reads a single keystroke and returns it's ASCII value.
 *	If a special key is pressed such as an arrow-key, it will process
 *	the ^] + ] return codes and return the ASCII value + 100.
 *
 * 	@return Integer
 *
 */
func ReadKey() int {
	var ch int

	ch = int(C.readkey())
	if ch == 27 { // ESC
		if int(C.readkey()) == 91 { // [ bracket
			ch = int(C.readkey()) + 100 // Add 100 to ASCII value.
		}
	}

	return ch
}

/**
 *
 *	RawKey
 *
 *	Reads a single keystroke and returns it's ASCII value.
 *	Does nothing with special keys being pressed.
 *	You will have to read each character code and process them
 *	yourself.
 *
 * 	@return Integer
 *
 */
func RawKey() int {
	return int(C.readkey())
}
