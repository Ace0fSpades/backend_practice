package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var dictionaryForNotations = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func findIndexByValue(value string) int {
	for i, v := range dictionaryForNotations {
		if v == value {
			return i
		}
	}
	return -1 // Если значение не найдено
}

func changeNotationTask1_1(number string, baseNotation int, toChangeNotation int) (final string) {
	buff := 0
	for i := 0; i < len(number); i++ {
		buff += findIndexByValue(string(number[i])) * int(math.Pow(float64(baseNotation), float64(len(number)-i-1)))
	}
	for buff >= toChangeNotation {
		final += dictionaryForNotations[buff%toChangeNotation]
		buff /= toChangeNotation
	}
	final += dictionaryForNotations[buff]

	return reverse(final)

}

func quadRootsTask1_2(a, b, c float64) (complex128, complex128) {
	if a == 0 {
		panic("Уравение не квадратичное")
	}

	D := b*b - 4*a*c

	if D > 0 {
		root1 := (-b + math.Sqrt(D)) / (2 * a)
		root2 := (-b - math.Sqrt(D)) / (2 * a)
		return complex(root1, 0), complex(root2, 0)
	}

	if D == 0 {
		root := -b / (2 * a)
		return complex(root, 0), complex(root, 0)
	}

	realPart := -b / (2 * a)
	imaginaryPart := math.Sqrt(-D) / (2 * a)
	return complex(realPart, imaginaryPart), complex(realPart, -imaginaryPart)
}

func sortAbsArrayTask1_3(arr []int) []int {
	var tmp int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if math.Abs(float64(arr[i])) > math.Abs(float64(arr[j])) {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}

func joinArraysTask1_4(arr []int, brr []int) []int {
	result := make([]int, 0, len(arr)+len(brr))

	i, j := 0, 0

	for i < len(arr) && j < len(brr) {
		if arr[i] <= brr[j] {
			result = append(result, arr[i])
			i++
		} else {
			result = append(result, brr[j])
			j++
		}
	}

	result = append(result, arr[i:]...)
	result = append(result, brr[j:]...)

	return result
}

func findSubstringTask1_5(mainString, subString string) int {
	if len(subString) == 0 {
		panic("Длина подстроки 0")
	}

	subLength := len(subString)
	mainLength := len(mainString)

	for i := 0; i < mainLength; i++ {
		match := true
		for j := 0; j < subLength; j++ {
			if i+j >= mainLength || mainString[i+j] != subString[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}

func calculateTask2_1(x, y int, operation string) (result float64) {
	validOperations := []string{"+", "-", "*", "/", "%", "^"}
	if !slices.Contains(validOperations, operation) {
		panic("Недопустимая операция")
	}

	switch operation {
	case "+":
		{
			result = float64(x + y)
		}
	case "-":
		{
			result = float64(x - y)
		}
	case "*":
		{
			result = float64(x * y)
		}
	case "/":
		{
			if y == 0 {
				panic("Деление на ноль")
			}
			result = float64(x) / float64(y)
		}
	case "%":
		{
			if y == 0 {
				panic("Деление на ноль")
			}
			result = float64(x % y)
		}
	case "^":
		{
			if y < 0 {
				y *= -1
				for i := 0; i < y; i++ {
					result *= float64(x)
				}
				result = 1 / result
			} else {
				for i := 0; i < y; i++ {
					result *= float64(x)
				}
			}
		}
	}
	return result
}

func isPalindromeTask2_2(s string) bool {

	s = strings.ToLower(strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r >= 'A' && r <= 'Z' {
			return r
		}
		return -1
	}, s))

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func hasIntersectionTask2_3(a1, a2, b1, b2, c1, c2 int) bool {

	if a1 > a2 {
		a1, a2 = a2, a1
	}
	if b1 > b2 {
		b1, b2 = b2, b1
	}
	if c1 > c2 {
		c1, c2 = c2, c1
	}

	maxStart := max(a1, b1, c1)
	minEnd := min(a2, b2, c2)

	return maxStart <= minEnd
}

func findLongestWordTask2_4(input string) string {

	words := strings.FieldsFunc(input, func(r rune) bool {
		if unicode.IsPunct(r) || r == ' ' {
			return true
		}
		return false
	})

	maxLength := 0
	for _, word := range words {
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}

	for _, word := range words {
		if len(word) == maxLength {
			return word
		}
	}

	return ""
}

func leapYearTask2_5(year int) string {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return "Високосный"
			}
		} else {
			return "Високосный"
		}
	}
	return "Не високосный"
}

func fiboTask3_1(number int) []int {
	arr := make([]int, number)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < number; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

func eratosthenesTask3_2(leftNumber, rightNumber int) []int { //Дописать диапазон для двух чисел
	arr := make([]bool, rightNumber+1)
	for i := leftNumber; i <= rightNumber; i++ {
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				arr[i] = true
				break
			}

		}
	}
	var primes []int

	for i, isComposite := range arr {
		if i >= leftNumber && !isComposite {
			primes = append(primes, i)
		}
	}

	return primes
}

func isArmstrong(num int) bool {
	digitCount := len(strconv.Itoa(num))
	tmpNum := num
	sum := 0
	for tmpNum > 0 {
		digit := tmpNum % 10
		sum += int(math.Pow(float64(digit), float64(digitCount)))
		tmpNum /= 10
	}
	return sum == num
}

func findArmstrongNumbersTask3_3(start, end int) []int {
	var armstrongs []int
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			armstrongs = append(armstrongs, i)
		}
	}
	return armstrongs
}

func reverseStringTask3_4(s string) string {
	b := []byte(s)
	result := make([]byte, len(b))

	left := 0
	right := len(b) - 1

	for left <= right {
		result[left] = b[right]
		result[right] = b[left]

		left++
		right--
	}

	return string(result)
}

func GCDTask3_5(a, b int) int {
	if b > a {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var taskNumber string

	fmt.Print("Выберите номер задания 1.1-1.5, 2.1-2.5, 3.1-3.5\nНапример для задания 1.3: 1.3\nВведите номер задания: ")
	fmt.Scanln(&taskNumber)

	switch taskNumber {
	case "1.1":
		{
			var number string
			var baseNotation, toChangeNotation int
			fmt.Println("Введите число, его СС и в какую СС надо перевести: ")
			fmt.Scan(&number, &baseNotation, &toChangeNotation)
			fmt.Printf("Входные данные: %s, %d, %d\nРезультат перевода: %s", number, baseNotation, toChangeNotation, changeNotationTask1_1(number, baseNotation, toChangeNotation))
		}
	case "1.2":
		{
			var a, b, c float64
			fmt.Print("Введите коэф. квадратного уравнения (a, b, c): ")
			fmt.Scanln(&a, &b, &c)
			r1, r2 := quadRootsTask1_2(a, b, c)
			fmt.Printf("Входные данные: %f, %f, %f\nРешение уравнения: %v, %v", a, b, c, r1, r2)
		}
	case "1.3":
		{
			var a []int
			var l int
			fmt.Print("Введите кол-во вводимых чисел: ")
			fmt.Scan(&l)
			fmt.Print("Введите целочисленные значения для массива: ")
			for i := 0; i < l; i++ {
				var tmp int
				fmt.Scan(&tmp)
				a = append(a, tmp)
			}
			fmt.Printf("Входные данные: %v\nСортировка по абсолютному значению: %v", a, sortAbsArrayTask1_3(a))
		}
	case "1.4":
		{
			var a, b []int
			var l int
			fmt.Print("Введите кол-во вводимых чисел: ")
			fmt.Scan(&l)
			fmt.Print("Введите целые значения для первого массива по порядку возрастания: ")
			for i := 0; i < l; i++ {
				var tmp int
				fmt.Scan(&tmp)
				a = append(a, tmp)
			}
			fmt.Print("Введите целые значения для второго массива по порядку возрастания: ")
			for i := 0; i < l; i++ {
				var tmp int
				fmt.Scan(&tmp)
				b = append(b, tmp)
			}
			fmt.Printf("Входные данные: %v, %v\nСлияние двух отсортированных массивов: %v", a, b, joinArraysTask1_4(a, b))
		}
	case "1.5":
		{
			var stra, strb string
			fmt.Print("Введите основную строку и подстроку: ")
			fmt.Scanln(&stra, &strb)
			fmt.Printf("Входные данные: %s, %s\nИндекс первого вхождения: %d", stra, strb, findSubstringTask1_5(stra, strb))
		}
	case "2.1":
		{
			var a, b int
			var op string
			fmt.Print("Введите число операцию число через пробел: ")
			fmt.Scanln(&a, &op, &b)
			fmt.Printf("Входные данные: %d, %s, %d\nОтвет: %.2f", a, op, b, calculateTask2_1(a, b, op))
		}
	case "2.2":
		{
			var s string
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Введите строку: ")
			s, _ = reader.ReadString('\n')
			fmt.Printf("Входные данные: %s\nПроверка на палиндром: %t", s, isPalindromeTask2_2(s))
		}
	case "2.3":
		{
			var a1, a2, b1, b2, c1, c2 int
			fmt.Print("Введите координаты на числовой оси для отрезков A1A2, B1B2, C1C2: ")
			fmt.Scanln(&a1, &a2, &b1, &b2, &c1, &c2)
			fmt.Printf("Входные данные: A1A2(%d, %d), B1B2(%d, %d), C1C2(%d, %d)\nПересечение отрезков: %t", a1, a2, b1, b2, c1, c2, hasIntersectionTask2_3(a1, a2, b1, b2, c1, c2))
		}
	case "2.4":
		{
			var s string
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Введите строку: ")
			s, _ = reader.ReadString('\n')
			fmt.Printf("Входные данные: %s\nСамое длинное слово: %s", s, findLongestWordTask2_4(s))
		}
	case "2.5":
		{
			var year int
			fmt.Print("Введите год: ")
			fmt.Scanln(&year)
			fmt.Printf("Входные данные: %d\nГод: %s", year, leapYearTask2_5(year))
		}
	case "3.1":
		{
			var a int
			fmt.Print("Введите целочисленное число: ")
			fmt.Scanln(&a)
			fmt.Printf("Входные данные: %d\nРяд Фибоначчи: %v", a, fiboTask3_1(a))
		}
	case "3.2":
		{
			var lNumber, rNumber int
			fmt.Print("Введите границы для простых чисел: ")
			fmt.Scanln(&lNumber, &rNumber)
			fmt.Printf("Входные данные:%d, %d\nПростые числа: %v", lNumber, rNumber, eratosthenesTask3_2(lNumber, rNumber))
		}
	case "3.3":
		{
			var lNumber, rNumber int
			fmt.Print("Введите границы для чисел Армстронга: ")
			fmt.Scanln(&lNumber, &rNumber)
			fmt.Printf("Входные данные:%d, %d\nЧисла Армстронга: %v", lNumber, rNumber, findArmstrongNumbersTask3_3(lNumber, rNumber))
		}
	case "3.4":
		{
			var s string
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Введите строку: ")
			s, _ = reader.ReadString('\n')
			fmt.Printf("Входные данные: %s\nРеверс строки: %s", s, reverseStringTask3_4(s))
		}
	case "3.5":
		{
			var a, b int
			fmt.Print("Введите два числа: ")
			fmt.Scanln(&a, &b)
			fmt.Printf("Входные данные:%d, %d\nНОД: %d", a, b, GCDTask3_5(a, b))
		}
	}

}
