/*
Create: 2022/7/14
Project: go-leak-example
Github: https://github.com/landers1037
Copyright Renj
*/

// Package go_leak_example
package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"
)

var Examples = map[string]func(){
	"slice1":     ExampleSlice,
	"slice2":     ExampleString,
	"goroutine1": ExampleGoroutine,
	"defer":      ExampleDefer,
	"http":       ExampleHTTP,
	"ticker":     ExampleTicker,
	"cgo":        ExampleCgo,
}

func main() {
	arg := flag.String("r", "", "example to run")
	examples := flag.Bool("e", false, "show examples map")
	flag.Parse()

	if *examples {
		for k, _ := range Examples {
			fmt.Println(k)
		}

		return
	}
	go runExample(*arg)
	err := http.ListenAndServe(":8899", nil)
	fmt.Println(err)
}

func runExample(arg string) {
	if arg == "" {
		return
	}
	f := Examples[arg]
	f()

	runtime.GC()
	debug.FreeOSMemory()
}
