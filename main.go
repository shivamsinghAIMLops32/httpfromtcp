package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	myFile,err := os.Open("./message.txt")

	if err != nil{
		fmt.Println("error in opening the file")
		os.Exit(1)
	}

	defer myFile.Close()

	bufferArr := make([]byte, 8)

	for {
		numberofByteSRead, err := myFile.Read(bufferArr)

         if numberofByteSRead > 0{
            fmt.Printf("%s\n",string(bufferArr[:numberofByteSRead]))
		 }
		if err != nil{
			if err == io.EOF{
				fmt.Println("End of file reached")
			      break
		     }
			fmt.Println("error in reading the file")
			os.Exit(1)
		}
	}
}