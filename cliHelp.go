package main

import("fmt")

func help() error {
    fmt.Println("Welcome!")
    fmt.Println("This tool will help get the latest pokemon information")
    fmt.Println("Usage: ")
    fmt.Println(" ")
    fmt.Println("Help: Shows this message with a summary of each command")
    fmt.Println("Exit: Exits the program")
    fmt.Println("Mapn: Show next 20 map locations")
    fmt.Println("Mapp: Show previous 20 map locations")
	fmt.Println("Explore <area>: Show the pokemon that spawn on the passed area")
	fmt.Println("Catch <pokemon>: Tries to catch the passed pokemon")
	fmt.Println("Inspect <pokemon>: Inspect a pokemon IF catched")
    return nil
}
