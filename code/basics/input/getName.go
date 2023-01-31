package input

import (
   "fmt"
   "bufio"
   "os"
)

func GetUserName() string{
   reader:= bufio.NewReader(os.Stdin)
   fmt.Print("Enter text: ")
   text,_:= reader.ReadString('\n')
   return text
}
