package main

func main() {
	var mass = []string{"Hello", "2", "world", ":-)"}
	var result = make([]string, 0)
	for i := 0; i < len(mass); i++ {
		if len(mass[i]) <= 3 {
			result = append(result, mass[i])
		}
	}
}
