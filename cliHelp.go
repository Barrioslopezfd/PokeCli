package main

import("fmt")

func help() error {
    fmt.Print("Welcome!\nThis tool will help you get the latest pokemon information\nUsage:\n\nHelp: Shows this message with a summary of each command\nExit: Exits the program\n\n")
    return nil 
}

