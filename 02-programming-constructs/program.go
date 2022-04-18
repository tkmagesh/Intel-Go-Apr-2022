package main

import "fmt"

func main() {
	//if
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	//for v1.0
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	//for v2.0 (while)
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %d\n", numSum)

	//for v3.0 (infinite)
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Printf("x = %d\n", x)

	//switch case
	/* rank by score */
	/*
		score = 0 - 3  => "Terrible"
		score = 4 - 7 => "Not Bad"
		score = 8 - 9 => "Very Good"
		score = 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	score := 7
	/* switch score {
	case 0:
		fmt.Println("Terrible")
	case 1:
		fmt.Println("Terrible")
	case 2:
		fmt.Println("Terrible")
	case 3:
		fmt.Println("Terrible")
	case 4:
		fmt.Println("Not Bad")
	case 5:
		fmt.Println("Not Bad")
	case 6:
		fmt.Println("Not Bad")
	case 7:
		fmt.Println("Not Bad")
	case 8:
		fmt.Println("Very good")
	case 9:
		fmt.Println("Very good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid Score")
	}
	*/

	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid Score")
	}

	no := 21
	switch {
	case no%2 == 0:
		fmt.Printf("%d is an even number\n", no)
	case no%2 == 1:
		fmt.Printf("%d is an odd number\n", no)
	case no%3 == 0:
		fmt.Printf("%d is a multiple of 3\n", no)
	}

	//fallthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
		fallthrough
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
	}

	subscription := "pro"
	switch subscription {
	case "pro":
		fmt.Println("All Pro features")
		fallthrough
	case "premium":
		fmt.Println("All Premium features")
		fallthrough
	case "free":
		fmt.Println("All free features")
		fallthrough
	default:
		fmt.Println("All trial features")

	}
}
