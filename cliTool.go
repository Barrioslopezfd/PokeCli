package main
import ("fmt"
    "bufio"
    "os"
    "strings"
)

func start(){
    for {
        fmt.Print("pokedex > ")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        input := scanner.Text()
        fmt.Print("\n")
        c := separateInput(input, "cmd")
        param := separateInput(input, "param")
        cmd := getCommands(c, param)
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
    caught      []pokemonResponse
}

var conf = config{
    next: "https://pokeapi.co/api/v2/location-area",
    previous: "",
}

func getCommands(cmd string, param string) cliCommand {

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
        "explore": {
            name: "Explore",
            description: "Shows pokemons in the area",
            function: func() error {
                return explore(param)
            },
        },
        "catch": {
            name: "Catch",
            description: "Tries to catch a pokemon",
            function: func() error {
                return catch(param, &conf.caught)
            },
        },
    }
    return command[cmd]

}

func separateInput(input string, whatIWant string) string {
    parameter := strings.Fields(input)
    if len(parameter) == 1 {
        return input
    }
    if whatIWant== "cmd"{
        return parameter[0]
    }
   return parameter[1]
}
