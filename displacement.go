package main

import (
	"fmt"
)

/*s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.
You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo,
and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the
given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t,
representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2,
so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))*/

func main() {
	//variable declaration
	var acceleration, initV, initD, time float64
	fmt.Println("Enter values for acceleration, initial velocity, initial displacement")
	fmt.Scan(&acceleration, &initV, &initD, &time)
	//calling the function to calculate the displacement
	displacement := genDisplaceFn(acceleration, initV, initD)
	total := displacement(time)
	fmt.Println(total)
}

func genDisplaceFn(acc, iv, id float64) func(float64) float64 {
	return func(t float64) float64 {
		disp := (acc*t*t + iv*t + id) / 2
		return disp
	}
}
