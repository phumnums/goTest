package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("============ข้อ 0============")
	num0()

	fmt.Println("============ข้อ 1============")
	num1()

	fmt.Println("============ข้อ 1.2============")
	num1_2()

	fmt.Println("============ข้อ 2============")
	num2()

	fmt.Println("============ข้อ 3============")
	num3()

	fmt.Println("============ข้อ 3.1============")
	num3_1()

	fmt.Println("============ข้อ 4============")
	num4()

	fmt.Println("============ข้อ 4.1============")
	num4_1()

	fmt.Println("============ข้อ 5============")
	num5()

	fmt.Println("============ข้อ 6============")
	num6()

	fmt.Println("============ข้อพิเศษ============")
	num7()
}

/*-------------------------------ข้อ 0-------------------------------*/
func num0() {
	i := 2

	if i == 0 {
		fmt.Println(i, "is zero")
	}
	if i == 1 {
		fmt.Println(i, "is one")
	}
	if i == 2 {
		fmt.Println(i, "is two")
	}
	if i == 3 {
		fmt.Println(i, "is three")
	}
	if i != 0 && i != 1 && i != 2 && i != 3 {
		fmt.Println(i, "is not in cases")
	}
}

/*-------------------------------ข้อ 1-------------------------------*/
func num1() {
	countDiv3 := 0
	countNotDiv3 := 0

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, "is divisible by 3")
			countDiv3++

		} else {
			fmt.Println(i, "is not divisible by 3")
			countNotDiv3++
		}
	}

	fmt.Println("Divisible by 3 =", countDiv3)
	fmt.Println("Not Divisible by 3 =", countNotDiv3)
}

/*-------------------------------ข้อ 1.2-------------------------------*/

func num1_2() {
	resultPow := pow(4, 0)
	fmt.Println(resultPow)
}

func pow(x, n int) int {
	result := math.Pow(float64(x), float64(n))
	return int(result)
}

/*-------------------------------ข้อ 2-------------------------------*/

func num2() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	maxNum := Findmax(x)
	minNum := Findmin(x)
	fmt.Println("maxNum = ", maxNum)
	fmt.Println("minNum = ", minNum)
}

func Findmax(x []int) int {
	max := x[0]

	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}

func Findmin(x []int) int {
	min := x[0]

	for i := 0; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}

/*-------------------------------ข้อ 3-------------------------------*/
func num3() {
	count := 0
	for i := 1; i <= 1000; i++ {
		str := strconv.Itoa(i)
		count += strings.Count(str, "9")
	}
	fmt.Println("9 in 1-1000 =", count)
}

/*-------------------------------ข้อ 3.1-------------------------------*/
func num3_1() {
	someFunc(20)
}

func someFunc(x int) {
	count := 0
	for i := 1; i <= x; i++ {
		str := strconv.Itoa(i)
		count += strings.Count(str, "9")
	}
	fmt.Println("9 in 1-", x, " =", count)
}

/*-------------------------------ข้อ 4-------------------------------*/
func num4() {
	var myWords = "AW SOME GO!"
	var result string

	for i := 0; i < len(myWords); i++ {
		if myWords[i] != ' ' {
			result += string(myWords[i])
		}
		//
	}
	fmt.Println(result)
}

/*-------------------------------ข้อ 4.1-------------------------------*/
func num4_1() {
	cutText("ine t")
}

func cutText(text string) {
	var result string
	for i := 0; i < len(text); i++ {
		if text[i] != ' ' {
			result += string(text[i])
		}
	}
	fmt.Println(result)
}

/*-------------------------------ข้อ 5-------------------------------*/
type Person struct {
	Title     string
	Firstname string
	Lastname  string
	Age       string
	Address   string
}

func num5() {
	people := map[string]Person{
		"employee1": {
			Title:     "Mr.",
			Firstname: "Tim",
			Lastname:  "Carry",
			Age:       "25",
			Address:   "142rd roads, Virginir, 22202",
		},
		"employee2": {
			Title:     "Ms.",
			Firstname: "Laura",
			Lastname:  "McCory",
			Age:       "22",
			Address:   "123/western Street, New York, 12304",
		},
		"employee3": {
			Title:     "Mrs.",
			Firstname: "Alice",
			Lastname:  "Der",
			Age:       "27",
			Address:   "123/central Street, London, 12345",
		},
	}

	for _, person := range people {
		fmt.Println("Name:", person.Title, person.Firstname, person.Lastname)
		fmt.Println("Age:", person.Age)
		fmt.Println("Address:", person.Address)
		fmt.Println()
	}
}

/*-------------------------------ข้อ 6-------------------------------*/
type Company struct {
	Name     string
	Age      int
	Address  string
	Position string
	Salary   int
}

func num6() {
	var c Company
	c.Name = "Forest Gump"
	c.Age = 26
	c.Address = "123/central Street, London, 12345"
	c.Position = "Manager"
	c.Salary = 50000
	fmt.Println(c)
	fmt.Println("Name:", c.Name)
	fmt.Println("Age:", c.Age)
	fmt.Println("Address:", c.Address)
	fmt.Println("Position:", c.Position)
	fmt.Println("Salary:", c.Salary)
}

/*-------------------------------ข้อพิเศษ-------------------------------*/

func num7() {
	for i := 0; i < 6; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
