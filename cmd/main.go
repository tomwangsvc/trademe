package main

import (
	"flag"
	"trademe/app"
)

func main() {
	chunkSize := flag.Int("size", 2, "chunk size to process data")
	result := flag.String("result", "test1", "result shows test result")
	flag.Parse()

	properties := app.InitProperties("./files/properties.txt")
	app.LogInfo("properties initialised successfully with length %d", len(properties))

	switch *result {
	case "test1":
		app.LogInfo("result from test1: %+v", app.FilterAndUseLastItem(properties))

	case "test2":
		app.LogInfo("result from test2: %+v", app.FilterAndUseFirstItem(properties))

	case "test3":
		app.LogInfo("result from test3: %+v", app.FilterOutDuplicatedItem(properties))

	case "test4":
		app.LogInfo("result from test4: %+v", app.FilterByMutipleConditions(properties))

	case "test5":
		app.LogInfo("result from test5 by chunk size %d: %+v", *chunkSize, app.ProcessByChunk(properties, *chunkSize))
	}
}
