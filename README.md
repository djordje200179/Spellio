# Spellio

Word processing engine built with trie structure than can predict words and correct typing mistakes.
For every word information about frequency of them will be stored in the structure.

### Word adding
You can add words in two ways:
1. Manually (one-by-one)
2. Read from file

### Usage
Currently, there are two main mods available:
#### 1. Word prediction
Engine will predict words based on prefix that was inputted. 
Possible words will be sorted by their frequency in the engine.

Example: 
```go
fmt.Printf("hospi...: %v\n", engine.CompleteWord("hospi", 5))
```
Output:
`hospi...: [hospital hospitality hospitably hospitable]`

#### 2. Word correction
Engine will correct the word based on inputted text and keyboard layout. 
Expected number of mistakes is `len(word) / 3`.
Possible words will be sorted by number of changed letters and
by their frequency in the engine.

Example:
```go
fmt.Printf("housr?: %v\n", engine.CorrectWord("housr", spellio.EnglishKeyboardLayout, 5))
```
Output:
`housr?: [house hoist]`