package main

/**
 * User: lotus_zero
 * Date: 11/17/18
 * Time: 23:35 PM
 * Brief: A program to calculate Circle Area
 */

import "fmt"

const PI float32 = 3.1415926

func process(radius float32) float32 {
	return PI * radius * radius
}

func main() {
	var radius float32
	fmt.Println("Radius = ?")
	fmt.Scan(&radius)

	area := process(radius)

	fmt.Println("The area is ", area)

}