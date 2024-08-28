package main

import (
	//"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var out string
	for i := range os.Args{
		if(i == 0){
			continue
		}
		out = fmt.Sprintf("%v %v", out, os.Args[i])
	}
    fmt.Print(" _" + strings.Repeat("_", len([]rune(out))) + "_\n") 
    fmt.Print("/ " + strings.Repeat(" ", len([]rune(out))) + " \\\n")
	fmt.Print("| " + out + "  |\n")
    fmt.Print("\\_" + strings.Repeat("_", len([]rune(out))) + "_/\n")

    fmt.Println("   |")
    fmt.Println("   |")
    fmt.Println("   /")

    fmt.Println(" @")
    fmt.Println("/|\\")
    fmt.Println("/ \\")
}
