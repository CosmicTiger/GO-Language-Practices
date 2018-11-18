package main

import "fmt"

/**
 * User: lotus_zero
 * Date: 11/17/18
 * Time: 18:20 PM
 * Brief: A program to calculate Multiple Indeterminated # of Circle Area with Check
 */

 const PI float32 = 3.14159

 func process(radius float32) float32 {
 	return PI * radius * radius
 }

 func main() {
 	var radius, area float32
	var i int32

 	fmt.Println("TO STOP THE PROGRAM, ENTER 0 IN THE VALUE RADIUS")
 	fmt.Println("Radius = ?")
 	fmt.Scan(&radius)

	 for i = 1; i != 0 ; i++  {
		 if radius < 0 {
		 	area = 0
		 }else {
		 	p := process(radius)
		 	area = p
		 }

		 fmt.Println("Area = ", area)
		 fmt.Println("Radius = ?")
	 }

	 return
 }