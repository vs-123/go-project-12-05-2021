/*
* PROJECT: go-project-12-05-2021
* DESCRIPTION: Given an IPv4 address in the form of string, indentify if it's valid.
* AUTHOR: Vahin Sharma
*/

package main

import (
  "fmt"
  "regexp"
)

func isValidIP(ip string) bool {
  re, _ := regexp.MatchString(`^((25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)(\.|$)){4}\b`, ip)
  return re
}

func main() {
  tests := map[string]bool {
    "123.456.789.0": false,
    "83.247.294.22": false,
    "207.298.128.206": false,
    
    "12.255.56.1": true,
    "160.179.194.224": true,
    "17.84.144.122": true,
    
  }
  score := 0
  
  for q, a := range tests {
    fmt.Println("Input:", q)
    if returnedA := isValidIP(q); returnedA != a {
      fmt.Printf("Expected %d, instead got %d\n\n", a, returnedA)
    } else {
      score++
      fmt.Println("Passed\n")
    }
  }
  
  fmt.Printf("Score: %d out of %d\n", score, len(tests))
}
