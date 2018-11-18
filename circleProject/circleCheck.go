package main

/**
 * User: lotus_zero
 * Date: 11/17/18
 * Time: 16:25 PM
 * Brief: A program to calculate Circle Area with Check
 */

import "fmt"

const PI float32 = 3.14159

func process(radius float32) float32 {
	return PI * radius * radius
}

func main() {
	var radius, area float32
	fmt.Printf("Radius = ? ")
	fmt.Scan(&radius)

	if radius < 0 {
		process := process(radius)
		area = process
	} else{
		area = 0
	}

	fmt.Println("Area = ", area)

}