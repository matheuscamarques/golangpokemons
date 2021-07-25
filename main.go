package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
)

/**
* @see https://transform.tools/json-to-go
 */

type IIE8_NAME struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Chinese  string `json:"chinese"`
	French   string `json:"french"`
}

type BaseAtribute struct {
	HP        int `json:"HP"`
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
	SpAttack  int `json:"Sp_Attack"`
	SpDefense int `json:"Sp_Defense"`
	Speed     int `json:"Speed"`
}

type Pokemon struct {
	ID   int          `json:"id"`
	Name IIE8_NAME    `json:"name"`
	Type []string     `json:"type"`
	Base BaseAtribute `json:"base"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("pokemons.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("OPENED: pokemons.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// create a new variable to hold our json data
	var pokemons []Pokemon
	// unmarshal our json data
	json.Unmarshal(byteValue, &pokemons)
	// print the json data with the chosen language
	language := getLanguage()
	fmt.Println("\n" + language + ":")
	getMainMenu(pokemons, language)
}

// make a function for get stdin terminal input
func getStdin(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	input, _ := reader.ReadString('\n')
	return input
}

// make a menu for the user to choose options
func getMainMenu(pokemons []Pokemon, language string) {
	// print main menu
	fmt.Println("Choose your option:")
	fmt.Println("1. Print all pokemons")
	fmt.Println("2. Print pokemons by type")
	fmt.Println("3. Print pokemons by name")
	fmt.Println("4. Change language")
	fmt.Println("5. Exit")
	// get user input
	mainMenu := getStdin("Choose your option: ")
	// verify runtime os is windows
	if runtime.GOOS == "windows" {
		mainMenu = strings.TrimRight(mainMenu, "\r\n")
	} else {
		mainMenu = strings.TrimRight(mainMenu, "\n")
	}
	// check if the input is valid stdin
	if strings.Compare(mainMenu, "1") == 0 {
		// print all pokemons
		for _, pokemon := range pokemons {
			// use name of pokemon as key
			r := reflect.ValueOf(pokemon.Name)
			f := reflect.Indirect(r).FieldByName(language)
			for _, Ctype := range pokemon.Type {
				printPokemon(pokemon, pokemon.ID, f.String(), Ctype)
			}

		}
		getMainMenu(pokemons, language)
	} else if strings.Compare(mainMenu, "2") == 0 {
		getPokemonsByType(pokemons, language)
	} else if strings.Compare(mainMenu, "3") == 0 {
		getPokemonsByName(pokemons, language)
	} else if strings.Compare(mainMenu, "4") == 0 {
		// change language
		language = getLanguage()
		getMainMenu(pokemons, language)
	} else if strings.Compare(mainMenu, "5") == 0 {
		// exit
		os.Exit(0)
	} else {
		fmt.Println("Invalid input")
		getMainMenu(pokemons, language)
	}

}

func getPokemonsByName(pokemons []Pokemon, language string) {
	// get user input
	name := getStdin("Enter the name of the pokemon: ")
	// verify runtime os is windows
	if runtime.GOOS == "windows" {
		name = strings.TrimRight(name, "\r\n")
	} else {
		name = strings.TrimRight(name, "\n")
	}
	// check if the input is valid stdin
	if strings.Compare(name, "") == 0 {
		fmt.Println("Invalid input")
		getPokemonsByName(pokemons, language)
	} else {
		// print all pokemons
		achou := false
		for _, pokemon := range pokemons {
			// use name of pokemon as key
			r := reflect.ValueOf(pokemon.Name)
			f := reflect.Indirect(r).FieldByName(language)

			if strings.Contains(strings.ToLower(f.String()), strings.ToLower(name)) {
				for _, Ctype := range pokemon.Type {
					achou = true
					printPokemon(pokemon, pokemon.ID, f.String(), Ctype)
				}
			}

		}
		if !achou {
			fmt.Println("Pokemon not founded!")
			getPokemonsByName(pokemons, language)
		} else {
			fmt.Println()
		}

		getMainMenu(pokemons, language)
	}
}

// make a list pokemons by type
func getPokemonsByType(pokemons []Pokemon, language string) {
	// print all possibles types of pokemons make menu
	fmt.Println("Choose your type:")
	fmt.Println("1. Normal")
	fmt.Println("2. Fire")
	fmt.Println("3. Water")
	fmt.Println("4. Electric")
	fmt.Println("5. Grass")
	fmt.Println("6. Ice")
	fmt.Println("7. Poison")
	fmt.Println("8. Ground")
	fmt.Println("9. Flying")
	fmt.Println("10. Psychic")
	fmt.Println("11. Bug")
	fmt.Println("12. Rock")
	fmt.Println("13. Ghost")
	fmt.Println("14. Dragon")
	fmt.Println("15. Dark")
	fmt.Println("16. Steel")
	fmt.Println("17. Fairy")
	fmt.Println("18. None")
	fmt.Println("any to main menu")
	// get user input
	typeMenu := getStdin("Choose your type: ")
	// verify runtime os is windows
	if runtime.GOOS == "windows" {
		typeMenu = strings.TrimRight(typeMenu, "\r\n")
	} else {
		typeMenu = strings.TrimRight(typeMenu, "\n")
	}
	// check if the input is valid stdin
	if strings.Compare(typeMenu, "1") == 0 {
		// search for pokemon by type
		listPokemonsByType(pokemons, "Normal", language)
	}
	if strings.Compare(typeMenu, "2") == 0 {
		listPokemonsByType(pokemons, "Fire", language)
	}
	if strings.Compare(typeMenu, "3") == 0 {
		listPokemonsByType(pokemons, "Water", language)
	}
	if strings.Compare(typeMenu, "4") == 0 {
		listPokemonsByType(pokemons, "Electric", language)
	}
	if strings.Compare(typeMenu, "5") == 0 {
		listPokemonsByType(pokemons, "Grass", language)
	}
	if strings.Compare(typeMenu, "6") == 0 {
		listPokemonsByType(pokemons, "Ice", language)
	}
	if strings.Compare(typeMenu, "7") == 0 {
		listPokemonsByType(pokemons, "Poison", language)
	}
	if strings.Compare(typeMenu, "8") == 0 {
		listPokemonsByType(pokemons, "Ground", language)
	}
	if strings.Compare(typeMenu, "9") == 0 {
		listPokemonsByType(pokemons, "Flying", language)
	}
	if strings.Compare(typeMenu, "10") == 0 {
		listPokemonsByType(pokemons, "Psychic", language)
	}
	if strings.Compare(typeMenu, "11") == 0 {
		listPokemonsByType(pokemons, "Bug", language)
	}
	if strings.Compare(typeMenu, "12") == 0 {
		listPokemonsByType(pokemons, "Rock", language)
	}
	if strings.Compare(typeMenu, "13") == 0 {
		listPokemonsByType(pokemons, "Ghost", language)
	}
	if strings.Compare(typeMenu, "14") == 0 {
		listPokemonsByType(pokemons, "Dragon", language)
	}
	if strings.Compare(typeMenu, "15") == 0 {
		listPokemonsByType(pokemons, "Dark", language)
	}
	if strings.Compare(typeMenu, "16") == 0 {
		listPokemonsByType(pokemons, "Steel", language)
	}
	if strings.Compare(typeMenu, "17") == 0 {
		listPokemonsByType(pokemons, "Fairy", language)
	}
	if strings.Compare(typeMenu, "18") == 0 {
		listPokemonsByType(pokemons, "None", language)
	}
	getMainMenu(pokemons, language)
}

// list pokemons of a specific type
func listPokemonsByType(pokemons []Pokemon, types string, language string) {
	fmt.Printf("ID\tNAME\t\tTYPE\n")
	for _, pokemon := range pokemons {
		for _, Ctype := range pokemon.Type {
			if strings.Compare(Ctype, types) == 0 {
				r := reflect.ValueOf(pokemon.Name)
				f := reflect.Indirect(r).FieldByName(language)
				printPokemon(pokemon, pokemon.ID, f.String(), Ctype)
			}
		}
	}
}

func printPokemon(pokemon Pokemon, id int, name string, ctype string) {
	if len(name) >= 8 {
		fmt.Printf("%d\t%s\t%s\n", id, name, ctype)
	} else {
		fmt.Printf("%d\t%s\t\t%s\n", id, name, ctype)
	}

}

// make a menu for the user to choose an language
func getLanguage() string {
	var language string
	// print language menu
	fmt.Println("Choose your language:")
	fmt.Println("1. English")
	fmt.Println("2. Japanese")
	fmt.Println("3. Chinese")
	fmt.Println("4. French")
	// get user input
	language = getStdin("Choose your language: ")

	// verify runtime os is windows
	if runtime.GOOS == "windows" {
		language = strings.TrimRight(language, "\r\n")
	} else {
		language = strings.TrimRight(language, "\n")
	}

	// check if the input is valid stdin
	if strings.Compare(language, "1") == 0 {
		return "English"
	} else if strings.Compare(language, "2") == 0 {
		return "Japanese"
	} else if strings.Compare(language, "3") == 0 {
		return "Chinese"
	} else if strings.Compare(language, "4") == 0 {
		return "French"
	} else {
		fmt.Println("Invalid input")
		return getLanguage()
	}
}

// make a function to open http requests to the api and get the json data
func getData(url string) ([]byte, error) {
	// open http request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// read the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
