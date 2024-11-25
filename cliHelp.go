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
    return nil
}
