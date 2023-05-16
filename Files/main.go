package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readFileAndPrintToterminal(string(os.Args[1]))
}
func readFileAndPrintToterminal(filename string) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(string(bs))

}
