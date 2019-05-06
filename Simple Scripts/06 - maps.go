package main

import "fmt"

func main() {

	presAge := make(map[string]int)

	presAge["Name 1"] = 42
	fmt.Println(presAge)
	fmt.Println(presAge["Name 1"])
	fmt.Println(len(presAge))

	presAge["Name 2"] = 43
	fmt.Println(presAge)
	fmt.Println(presAge["Name 2"])
	fmt.Println(len(presAge))

	delete(presAge, "Name 1")
	fmt.Println(presAge)

	// map inside a map
	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"realname": "Clark Kent",
			"city":     "Metropolis",
		},

		"Batman": map[string]string{
			"realname": "Bruce Wayne",
			"city":     "Gotham City",
		},
	}
	fmt.Println(superhero)

	// We can output data where the key matches Superman

	temp, hero := superhero["Supermammeta"]
	fmt.Println("temp - hero", temp, hero)

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["realname"], temp["city"])
	}

}
