package main

import("fmt")

func help(commands map[string]cliCommand) error {
    
    fmt.Println("Welcome!")
    fmt.Println("This tool will help get the latest pokemon information")
    fmt.Println("Usage: ")
    fmt.Println(" ")
    for _, cmd := range commands {
	fmt.Println(cmd.name+": "+cmd.description)
    }
    return nil
}
