package keyboardlayouts

import "github.com/djordje200179/spellio"

// EnglishKeyboardLayout is a keyboard layout for the
// English language QUERTY keyboard.
var EnglishKeyboardLayout = spellio.KeyboardLayoutNearbyKeys{
	'a': {'q', 'w', 's', 'z'},
	'b': {'v', 'g', 'h', 'n'},
	'c': {'x', 'd', 'f', 'v'},
	'd': {'s', 'e', 'r', 'f', 'c', 'x'},
	'e': {'w', 's', 'd', 'r'},
	'f': {'d', 'r', 't', 'g', 'v', 'c'},
	'g': {'f', 't', 'y', 'h', 'b', 'v'},
	'h': {'g', 'y', 'u', 'j', 'n', 'b'},
	'i': {'u', 'j', 'k', 'o'},
	'j': {'h', 'u', 'i', 'k', 'm', 'n'},
	'k': {'j', 'i', 'o', 'l', 'm'},
	'l': {'k', 'o', 'p', 'č'},
	'm': {'n', 'j', 'k'},
	'n': {'b', 'h', 'j', 'm'},
	'o': {'i', 'k', 'l', 'p'},
	'p': {'o', 'l', 'č', 'š'},
	'q': {'a', 'w'},
	'r': {'e', 'd', 'f', 't'},
	's': {'w', 'e', 'd', 'x', 'z', 'a'},
	't': {'r', 'f', 'g', 'y'},
	'u': {'y', 'h', 'j', 'i'},
	'v': {'c', 'f', 'g', 'b'},
	'w': {'q', 'a', 's', 'e'},
	'x': {'y', 's', 'd', 'c'},
	'y': {'t', 'g', 'h', 'u'},
	'z': {'x', 's', 'a'},
}
