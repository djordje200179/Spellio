# Spellio

A word processing engine built with trie structure than can finish words during typing
and correct typing mistakes. 

The project will be enhanced with new features in the future (like prediction
of the next word based on the previous one).

Engine tries to return the most probable words, and because of that, it stores
information about the frequency of each word.
Therefore, adding one word multiple times will increase its stored frequency.

### Word adding
You can add words in two ways:
1. Manually (one-by-one)
2. Read from file

### Usage
Currently, there are two main mods available:

#### 1. Word completion
Engine will try to complete word based on the prefix that was inputted.
Possible words will be sorted by their frequency in the engine.

Example: 
```go
fmt.Printf("hospi...: %v\n", engine.CompleteWord("hospi", 5))
// hospi...: [hospital hospitality hospitably hospitable]
```

#### 2. Word correction
Engine will correct inputted text based on given keyboard layout. 
Expected number of mistakes is `len(word) / 3`.
Possible words will be sorted by number of mistaken letters and
by their frequency in the engine.

Example:
```go
fmt.Printf("housr?: %v\n", engine.CorrectWord("housr", layouts.English, 5))
// `housr?: [house hoist]`
```