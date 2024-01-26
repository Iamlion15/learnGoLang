// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	cycleNames([]string{"cloud", "tifa", "roul"}, sayGreeting)
// 	cycleNames([]string{"lorenzo", "gustavo", "gusto"}, sayGoodAfternoon)
// 	a1 := circleArea(10.5)
// 	a2 := circleArea(15)
// 	fmt.Println("Areas", a1)
// 	fmt.Printf("Area a2 is %0.3f", a2)
// }
// func cycleNames(n []string, f func(string)) {
// 	for _, value := range n {
// 		f(value)
// 	}
// }
// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v\n", n)
// }
// func sayGoodAfternoon(n string) {
// 	fmt.Printf("Good afternoon %v\n", n)
// }
// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }
