# ac automaton

This package implements ac automaton algorithm for fast matching many patterns in string.

## example

A simple use case:

```go
package main

import (
"fmt"
"github.com/we-miks/ac_automaton"
)

func main() {
    patterns := []string{"i", "is", "his", "her", "hers", "she"}
    ac := ac_automaton.NewACAutomaton()
    for _, pattern := range patterns {
        ac.Insert(pattern)
    }
    ac.Build()
    s := "uthersheis"
    res := ac.FindMatches(s)
    fmt.Println(res)
}
```