package main

import (
	"fmt"
	"os"

	ps "github.com/mitchellh/go-ps"
)

func psa(name string) {
	if name == "" {
		processList, err := ps.Processes()
		if err != nil {
			panic(err)
		}

		for x := range processList {
			var process ps.Process = processList[x]
			process = processList[x]
			fmt.Printf("%d\t%s\n", process.Pid(), process.Executable())
		}
		os.Exit(0)
	} else {
		processList, err := ps.Processes()
		if err != nil {
			panic(err)
		}

		for x := range processList {
			var process ps.Process = processList[x]
			if process.Executable() == name {
				fmt.Printf("%d\t%s\n", process.Pid(), process.Executable())
			}
		}
		os.Exit(0)
	}
}

func pspid(name string) {
	if name == "" {
		processList, err := ps.Processes()
		if err != nil {
			panic(err)
		}

		for x := range processList {
			var process ps.Process = processList[x]
			fmt.Printf("%d\n", process.Pid())
		}
		os.Exit(0)
	} else {
		processList, err := ps.Processes()
		if err != nil {
			panic(err)
		}

		for x := range processList {
			var process ps.Process = processList[x]
			if process.Executable() == name {
				fmt.Printf("%d\n", process.Pid())
			}
		}
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) == 1 {
		psa("")
	}
	if len(os.Args) == 3 {
		if os.Args[1] == "--name" || os.Args[1] == "-n" {
			psa(os.Args[2])
		} else if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println("Usage: psaux [--name|-n <name>]")
			os.Exit(0)
		} else if os.Args[1] == "--pid" || os.Args[1] == "-p" {
			pspid(os.Args[2])
		} else {
			fmt.Println("Usage: psaux [--name|-n <name>]")
			os.Exit(1)
		}
	}
}
