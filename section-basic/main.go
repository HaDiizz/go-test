package main

import (
	"fmt"
	"strconv"
)

func main() {
	//! ข้อ0
	no0()

	//! ข้อ1
	// no1()

	//! ข้อ1.2
	// no1_2()

	//! ข้อ2
	// no2()

	//! ข้อ3
	// no3()

	//! ข้อ4
	// no4()

	//! ข้อ4.1
	// no4_1()

	//! ข้อ5
	// no5()

	//! ข้อ6
	// no6()

	//! ข้อพิเศษ
	// special()
}

func no0() {
	i := 2
	fmt.Println("Example-: if case condition")
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Your i not in case.")
	}
}

func no1() {
	count := 0
	numbers := []int{}
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			count++
			numbers = append(numbers, i)
		}
	}
	fmt.Println("Total Numbers: ", count)
	fmt.Println("Number List: ", numbers)

}

func no1_2() {
	result := expoCalculation(20, 2)
	fmt.Println("No.2 | Expo: ", result)
}

func expoCalculation(num float64, expo int) float64 {
	if expo < 0 {
		num = 1 / num
		expo = -expo
	}
	result := 1.0
	for i := 0; i < expo; i++ {
		result *= num
	}
	return result
}

func no2() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	max := maxNum(x)
	println("Max number: ", max)
	min := minNum(x)
	println("Min number: ", min)
}

func maxNum(num []int) int {
	max := num[0]
	for i := 0; i < len(num); i++ {
		if num[i] > max {
			max = num[i]
		}
	}
	return max
}

func minNum(num []int) int {
	min := num[0]
	for i := 0; i < len(num); i++ {
		if num[i] < min {
			min = num[i]
		}
	}
	return min
}

func no3() {
	count := 0
	for i := 1; i <= 1000; i++ {
		for _, c := range strconv.Itoa(i) {
			if c == '9' {
				count++
			}
		}
	}
	fmt.Println("No.3 | 1-1000 มีเลข9รวมทั้งหมด: ", count)
}

func no4() {
	result := ""
	var myWords = "AW SOME GO!"
	for i := 0; i < len(myWords); i++ {
		c := string(myWords[i])
		if c == " " {
			result += ""
		} else {
			result += c
		}
	}
	println(result)
}

func no4_1() {
	const myWords = "ine t"

	result := cutText(myWords)

	println("cutText:", result)
}

func cutText(str string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		c := string(str[i])
		if c == " " {
			result += ""
		} else {
			result += c
		}
	}
	return result
}

func no5() {
	data := []map[string]string{
		{
			"nameTitle": "Mr.",
			"firstName": "Time",
			"lastName":  "Curry",
			"age":       "25",
			"address":   "142rd roads, Virginir, 22202",
		},
		{
			"nameTitle": "Ms.",
			"firstName": "Laura",
			"lastName":  "McCoy",
			"age":       "22",
			"address":   "123/western Street, New York, 12304",
		},
		{
			"nameTitle": "Mr",
			"firstName": "John",
			"lastName":  "Doe",
			"age":       "24",
			"address":   "123/Hatyai Street, Thailand, 90110",
		},
		{
			"nameTitle": "Mr",
			"firstName": "Jim",
			"lastName":  "Ron",
			"age":       "18",
			"address":   "123/Pattaya Street, Thailand, 1234",
		},
	}

	fmt.Println("Result -:")
	for _, item := range data {
		fmt.Printf("Name -: %s %s %s (Age: %s)\n", item["nameTitle"], item["firstName"], item["lastName"], item["age"])
		fmt.Printf("Address -: %s\n", item["address"])
	}
}

type Company struct {
	Name           string
	Address        string
	RegistrationAt string
	Value          float64
	Type           string
}

func no6() {
	var company Company = Company{
		Name:           "INTERNET THAILAND PUBLIC COMPANY LIMITED",
		Address:        "1768 อาคารไทยซัมมิท ทาวเวอร์ ชั้น 10-12 และชั้น IT ถ.เพชรบุรีตัดใหม่ แขวงบางกะปิ เขตห้วยขวาง กรุงเทพมหานคร 10310",
		RegistrationAt: "14 Sept. 2001",
		Value:          2200838143,
		Type:           "บริษัทมหาชนจำกัด"}

	fmt.Println("Company Name:", company.Name)
	fmt.Println("Address:", company.Address)
	fmt.Println("RegistrationAt:", company.RegistrationAt)
	fmt.Println("Value:", company.Value)
	fmt.Println("Type:", company.Type)
}

func special() {
	for row := 1; row <= 6; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
