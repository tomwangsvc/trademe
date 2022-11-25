package main

import (
	"fmt"
	"os"
	"trademe/app"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'test1' 'test2' 'test3' 'test4' 'test5' subcommands")
		os.Exit(1)
	}

	properties := app.InitProperties("./files/properties.txt")
	app.LogInfo("properties initialised successfully with length %+v", len(properties))

	switch os.Args[1] {
	case "test1":
		app.LogInfo("result from test1: %+v", app.FilterAndUseLastItem(properties))

	case "test2":
		app.LogInfo("result from test2: %+v", app.FilterAndUseFirstItem(properties))

	case "test3":
		app.LogInfo("result from test3: %+v", app.FilterOutDuplicatedItem(properties))

	case "test4":
		app.LogInfo("result from test4: %+v", app.FilterByMutipleConditions(properties))

	case "test5":
		app.LogInfo("result from test4: %+v", app.FilterByMutipleConditions(properties))
	}
}
