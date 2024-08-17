package main

import "fmt"

func main() {
	players1 := map[string]int8{
		"Curry":  4,
		"LeBron": 6,
	}

	players2 := map[string]int8{
		"Curry":  4,
		"LeBron": 6,
	}

	teams := map[string]string{
		"Warriors": "Golden State",
		"Lakers":   "Los Angeles",
	}

	fmt.Printf("%t\n", players1 == nil)                    // false
	fmt.Printf("%t\n", players1 == players1)               // invalid operation: players1 == players1 (map can only be compared to nil)
	fmt.Printf("%t\n", players1 == players2)               // invalid operation: players1 == players2 (map can only be compared to nil)
	fmt.Printf("%t\n", players1 == teams)                  // invalid operation: players1 == teams (mismatched types map[string]int8 and map[string]string)
	fmt.Printf("%t\n", players1 >= teams)                  // invalid operation: players1 == teams (mismatched types map[string]int8 and map[string]string)
	fmt.Printf("%t,%t\n", players1 > nil, players1 >= nil) // invalid operation: players1 > nil (operator > not defined on map) 以及 invalid operation: players1 >= nil (operator >= not defined on map)
	fmt.Printf("%t,%t\n", players1 < nil, players1 <= nil) // invalid operation: players1 > nil (operator < not defined on map) 以及 invalid operation: players1 <= nil (operator <= not defined on map)

	players1 = nil
	fmt.Printf("%t\n", players1 == nil)                    // true
	fmt.Printf("%t,%t\n", players1 > nil, players1 >= nil) // invalid operation: players1 > nil (operator > not defined on map) 以及 invalid operation: players1 >= nil (operator >= not defined on map)
	fmt.Printf("%t,%t\n", players1 < nil, players1 <= nil) // invalid operation: players1 > nil (operator < not defined on map) 以及 invalid operation: players1 <= nil (operator <= not defined on map)
}
