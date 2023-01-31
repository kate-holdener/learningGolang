package input

import (
   "fmt"
   "bufio"
   "os"
)

func GetUserName(){
   reader:= bufio.NewReader(os.Stdin)
   fmt.Print("Enter text: ")
   text:= reader.ReadString('\n')
   return text
}
