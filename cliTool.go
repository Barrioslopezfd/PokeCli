package main
import ("fmt"
    "bufio"
    "os"
)

func start(){
    for {
        fmt.Print("pokedex > ")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        input := scanner.Text()
        fmt.Print("\n")
        cmd := getCommands(input)

        cmd.function()
    }
}

type cliCommand struct {
    name            string
    description     string
    function        func() error
}

type pokeResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
    next	string
    previous	string
}

var conf = config{
    next: "https://pokeapi.co/api/v2/location-area",
    previous: "",
}

func getCommands(cmd string) cliCommand {

    command :=  map[string]cliCommand {
        "help": {
                name: "Help",
                description: "Brief summay on how to use the tool",
                function: help,
            },
        "exit": {
                name: "Exit",
                description: "Exits the program",
                function: exit,
            },
        "mapn": {
                name: "Map",
                description: "Shows the next 20 locations",
                function: func() error {
                    return maps(&conf, "next")
                },
            },
        "mapp": {
                name: "Map",
                description: "Shows the previous 20 locations",
                function: func() error {
                    return maps(&conf, "previous")
                },
            },
    }
    return command[cmd]

}

