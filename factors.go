package main

import (
    "fmt"
    "bufio"
    "os"
    "math/big"

)
func main() {
   
    if len(os.Args) < 2 {
        fmt.Println("Usage: factors <file>")
    }else{    

        arg := os.Args[1]

        file, err := os.Open(arg)

        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        defer file.Close()

        scanner := bufio.NewScanner(file)
        n := new(big.Int)

        for scanner.Scan() {
            line := scanner.Text()
            n.SetString(line, 10)

            var flag_ bool
            half := new(big.Int).Div(n, big.NewInt(2))
            for i := big.NewInt(2); i.Cmp(half) <= 0; i = new(big.Int).Add(i, big.NewInt(1)) {
                if new(big.Int).Mod(n, i).Cmp(big.NewInt(0)) == 0 {
                    fmt.Printf("%v=%v*%v\n", n, new(big.Int).Div(n, i), i)
                    flag_ = true
                    break
                }
            }

            if !flag_ {
                fmt.Println("Cannot factorize the given number.")
            }
        }

        if err := scanner.Err(); err != nil {
            fmt.Println("Error:", err)
        }
           
    }
}
