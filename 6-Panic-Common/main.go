package main

func main() {
	// 1. Nil Pointer Dereference
	// var p *int
	// fmt.Println(*p) // This will cause a panic because p is nil

	// 2. Array or Slice Index Out of Bounds
	// arr := []int{1, 2, 3}
	// fmt.Println(arr[10]) // This will cause a panic because the index is out of bounds

	// 3. Map Access with Nil Map
	// var m map[string]int
	// fmt.Println(m["key"]) // This will cause a panic because m is nil

	// 4. Division by Zero
	// result := 10 / 0 // This will cause a panic because dividing by zero is not allowed

	// 5. Type Assertion on Wrong Type
	// var i interface{} = 42
	// str := i.(string) // This will cause a panic because i holds an int, not a string

	// 6. Using Panic Function
	// panic("Something unexpected happened")

	// 7. Recursive Panics
	// recursivePanic() // Uncomment this line to see a recursive panic

}

// func recursivePanic() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from panic:", r)
// 			panic("Another panic") // Recursive panic
// 		}
// 	}()
// 	panic("Initial panic")
// }
