package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
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
        var n, i, flag int
        for scanner.Scan() {
            line := scanner.Text()
            n, err = strconv.Atoi(line)


            for i = 2; i <= n/2; i++ {
                if n%i == 0 {
                    fmt.Printf("%d=%d*%d\n", n, n/i, i)
                    flag = 1
                    break
                }
            }

            if flag == 0 {
                fmt.Println("Cannot factorize the given number.")
            }
        }

        if err := scanner.Err(); err != nil {
            fmt.Println("Error:", err)
        }
           
    }
}
