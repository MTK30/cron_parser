package main

import (
	"fmt"
	"os"

	"github.com/MTK30/cron_parser/parser"
)

func main() {
	var cronString string
	if len(os.Args) < 2 {
		cronString = "1-4 * */15 * 1-5 /usr/bin/find"

	} else {
		cronString = os.Args[1]
	}
	parse, err := parser.GetParserInstance(cronString)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	err = parse.Handler()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	fmt.Println(parse.String())
}
