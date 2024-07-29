package main

import "fmt"

func main() {
	var players map[string]int8
	//num := players["Durant"]
	//fmt.Println(num)
	fmt.Printf("%[1]v, %[1]T, %d, %t\n", players, len(players), players == nil) // map[], map[string]int8, 0, true

	//players["Curry"] = 4 // panic: assignment to entry in nil map
	players = map[string]int8{
		//"Curry":  4 // syntax error: unexpected newline in composite literal; possibly missing comma or }
		"Curry":  4, // syntax error: unexpected newline in composite literal; possibly missing comma or }
		"LeBron": 6,
	}

	num, ok := players["Durant"]
	fmt.Println(num, ok) // 0 false
	num, ok = players["LeBron"]
	fmt.Println(num, ok) // 6 true
	//fmt.Println(cap(players)) // invalid argument: players (variable of type map[string]int8) for cap
	players = nil
	fmt.Printf("%[1]v, %[1]T, %d, %t\n", players, len(players), players == nil)

	//equipments := make(map[string]float64, 3, 3) // invalid operation: make(map[string]float64, 3, 3) expects 1 or 2 arguments; found 3
	equipments := make(map[string]float64, 3)
	fmt.Printf("%[1]v, %[1]T, %d, %t\n", equipments, len(equipments), equipments == nil) // map[], map[string]float64, 0, false

}
