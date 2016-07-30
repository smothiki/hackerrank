package main

import "fmt"
import "bufio"
import "os"

func main(){
    reader := bufio.NewReader(os.Stdin)
text, _ := reader.ReadString('\n')
    fmt.Println("Hello, World.")
    fmt.Println(text)
}
