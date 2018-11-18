package main

import "fmt"

/**
 * User: lotus_zero
 * Date: 11/17/18
 * Time: 16:54 PM
 * Brief: A program to calculate Multiple Circle Area with Check
 */

 const PI float32 = 3.14159

 func process(radius float32) float32 {
 	return PI * radius * radius
 }

 func main() {
 	var radius, area float32
 	var n, i int32
 	fmt.Println("# of Circles?")
 	fmt.Scan(&n)

	 for i = 0; i < n ; i++  {

	 	fmt.Println("Circle # ",i," Radius = ?")
	 	fmt.Scan(&radius)

	 	if radius < 0 {
	 		area = 0
		}else {
			p := process(radius)
			area = p
		}

	 	fmt.Println("Area = ", area)
	 }

	 return
 }