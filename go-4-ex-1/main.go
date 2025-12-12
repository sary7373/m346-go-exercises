package main

import (
	"errors"
	"fmt"
)

// TODO: implement the function computeGrade
func computeGrade(gotPoints float32, maxPoints float32) (float32, error) {
	grade := gotPoints/maxPoints*5 + 1
	var err error
	if gotPoints < 0 || maxPoints < 0 || gotPoints > maxPoints {
		err = errors.New("invalid grade")
	} else {
		err = nil
	}
	return grade, err
}

func main() {
	// TODO: call the function computeGrade
	fmt.Println(computeGrade(17.5, 28.0))  //4.125
	fmt.Println(computeGrade(67, 69))      //5.8550725
	fmt.Println(computeGrade(8.25, 15.5))  //3.6612904
	fmt.Println(computeGrade(-8.25, 15.5)) //invalid grade
	fmt.Println(computeGrade(20, 15.5))    //invalid grade
}
