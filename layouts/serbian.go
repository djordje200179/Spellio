package layouts

import "github.com/djordje200179/spellio"

// Serbian is a keyboard layout for
// the Serbian language QWERTZ keyboard.
var Serbian = spellio.KeyboardLayout{
	'a': {'q', 'w', 's', 'y'},
	'b': {'v', 'g', 'h', 'n'},
	'c': {'x', 'd', 'f', 'v'},
	'č': {'l', 'p', 'š', 'ć'},
	'ć': {'č', 'š', 'đ', 'z'},
	'd': {'s', 'e', 'r', 'f', 'c', 'x'},
	'đ': {'š', 'ć', 'ž'},
	'e': {'w', 's', 'd', 'r'},
	'f': {'d', 'r', 't', 'g', 'v', 'c'},
	'g': {'f', 't', 'z', 'h', 'b', 'v'},
	'h': {'g', 'z', 'u', 'j', 'n', 'b'},
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
	's': {'w', 'e', 'd', 'x', 'y', 'a'},
	'š': {'p', 'č', 'ć', 'đ'},
	't': {'r', 'f', 'g', 'z'},
	'u': {'z', 'h', 'j', 'i'},
	'v': {'c', 'f', 'g', 'b'},
	'w': {'q', 'a', 's', 'e'},
	'x': {'y', 's', 'd', 'c'},
	'z': {'t', 'g', 'h', 'u'},
	'ž': {'ć', 'đ'},
	'y': {'x', 's', 'a'},
}
