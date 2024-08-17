package main

import (
	//"flag"
	"fmt"
	"os"

)

func main() {
	var out string
	for i := range os.Args{
		if(i == 0){
			continue
		}
		out = fmt.Sprintf("%v %v", out, os.Args[i])
	}
	
	fmt.Println(out)
}