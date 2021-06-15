<h1 align="center">Useful links for GoLang</h1>

---

* https://golang.org/ref/spec#Numeric_types
* https://golang.org/pkg/fmt/
* https://golang.org/pkg/
```go 
    // string variables
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	// the following is allowed inside functions only
	nameFour := "yoshi"

	// int variables
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

    // bits & memory
	var numOne int8 = 25
	var numTwo int8 = 128 // too large a number for 8-bit
	var numTwo uint = -25 // unsigned ints cannot be negative

    // float variables
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 1965385877.5
	var scoreThree = 1.5 // inferred as float64
    
    // Arrays
    var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

    // Slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

    // slice ranges
	rangeOne := names[1:4] // doesn't include pos 4 element
	rangeTwo := names[2:]  //includes the last element
	rangeThree := names[:3]
    
    // While Loop
    x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++ // infinite loop without this
	}

    // For Loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

    for index, val := range names {
		fmt.Printf("the value at pos %v is %v \n", index, val)
		val = "new string"
	}

	// if else
	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 40")
	}

	// Functions
	func sayBye(n string) {
		fmt.Printf("Goodbye %v \n", n)
	}

	func cycleNames(n []string, f func(string)) {
		for _, v := range n {
			f(v)
		}
	}

	func circleArea(r float64) float64 {
		return math.Pi * r * r
	}