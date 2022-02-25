package main

import "fmt"

func main()  {
	var gamingConsole []string

	gamingConsole = append(gamingConsole, "PlayStation4, ")
	gamingConsole = append(gamingConsole, "Nintendo Switch, ")
	gamingConsole = append(gamingConsole, "Xbox One")

	fmt.Println(gamingConsole)

	for _, console := range gamingConsole {
		fmt.Println(console)
	}

	gamingConsoles := []string{"PlayStation4, ", "Nintendo Switch, ", "Xbox One"}
	fmt.Println(gamingConsoles)
}