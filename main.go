package main
import(
    "fmt"
    "bufio"
    "os"
)

func main() {
    type cliCommand struct {
        name            string
        description     string
        callback            func() error
    }

    

    for {
        fmt.Print("pokedex > ")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        input := scanner.Text()
        fmt.Printf("Just wrote: %q\n", input)
    }
}
