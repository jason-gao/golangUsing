package fmt

import "fmt"

/**
- https://studygolang.com/articles/2644

 */
type Human struct{
	Name string
}

var people = Human{Name:"zhangsan"}

func TestPrintf()  {
	fmt.Printf("%v", people)
	fmt.Println()
	fmt.Printf("%+v", people)
}