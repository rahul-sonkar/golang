package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main() {
	_ = ex1
	_ = ex2
	_ = ex3
	_ = ex4
	_ = ex5
	_ = ex6
	_ = ex7
	_ = ex8
	_ = ex9
	_ = ex10
	_ = ex11
	_ = ex12
	_ = ex13
	_ = ex14
	_ = stdout
	_ = stdin
	_ = decode
}

// PicoCTF
func decode() {
	data := [...]int{112,105,99,111,67,84,70,123,103,48,48,100,95,107,49,116,116,121,33,95,110,49,99,51,95,107,49,116,116,121,33,95,55,99,48,56,50,49,102,53,125,10}
	
	var str string
	for _,v := range data {
		str += fmt.Sprintf("%c",v)
	}

	fmt.Println(str)
}

// Input from standar input.
func stdin() {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	io.WriteString(os.Stdout,scanner.Text())
}

// Output from standar output
func stdout() {
	msg := "this output from stdout."
	io.WriteString(os.Stdout,msg)
}

// Multiply two numbers by Russian peasant method.
func ex14() {
	var a,b,mul int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d %d",&a,&b)
	
	for {
		if a%2 != 0 {
			mul += b
		}
		if a == 1 {break}
		a = a/2
		b = b*2
	}

	fmt.Printf("Answer is: %d\n",mul)
}

// Check whether a date is valid or not.
func ex13() {
	var d, m, y int
	var flag bool = true
	fmt.Printf("Enter date: ")
	fmt.Scanf("%d-%d-%d", &d, &m, &y)
	if d < 1 && d > 31 || m < 1 && m > 12 || y < 1 {
		fmt.Println("Pleas check date.")
		return
	}

	isleap := func() bool {
		if y%4 == 0 || y%100 == 0 && y%400 == 0 {
			return true
		}
		return false
	}

	if y < 1850 || y > 2050 {
		flag = false
	} else if m == 2 && d>28 || d>29 && isleap() {
		flag = false
	} else if m == 4 || m == 6 || m  == 9 || m == 11 {
		if d == 31 {flag = false}
	}
	
	if flag {
		fmt.Println("Valid date.")
	} else {
		fmt.Println("Invalid date.")
	}
}

// Enter a number and test whether it is a fibonacci number or not.
func ex12() {
	var (
		a           = 1
		b           = 1
		temp        int
		num         int
		isFibonacci bool = false
	)

	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	for b <= num {
		temp = a
		a = b
		b = a + temp
		if b == num {
			isFibonacci = true
		}
	}

	if isFibonacci {
		fmt.Printf("%d is fibonacci number\n", num)
	} else {
		fmt.Printf("%d isn't fibonacci number\n", num)
	}
}

// Find out the value of x raised to the power y, where x and y are positive integers.
func ex11() {
	var x, y int
	fmt.Printf("Enter x and y: ")
	fmt.Scanf("%d %d", &x, &y)
	pow := func(a, b int) (p int) {
		p = 1
		for i := 0; i < b; i++ {
			p *= a
		}
		return
	}
	fmt.Printf("%d to the power %d is: %d\n", x, y, pow(x, y))
}

// Accept any number n and print the sum of cube of all numbers for 1 to n which are divisible by 3.
func ex10() {
	var num int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		if i%3 != 0 {
			continue
		}
		fmt.Printf("%d's cube is: %d\n", i, i*i*i)
	}
}

// Accept any number n and print the sum of squre of all nmbers from 1 to n.
func ex9() {
	var num int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		fmt.Println(i, "'s squre:", i*i)
	}
}

// Input a number and a digit, and count the number of times the digit occurs in the number.
func ex8() {
	var (
		num   int
		digit int
	)
	fmt.Printf("Enter number and digit: ")
	fmt.Scanf("%d %d", &num, &digit)

	count := func(n int, d int) (i int) {
		if n == d {
			return 1
		}

		for n > 0 {
			rem := n % 10
			if rem == d {
				i++
			}
			n = n / 10
		}
		return
	}

	fmt.Printf("%d times the digit (%d) occurs in the number (%d)\n", count(num, digit), digit, num)
}

// Input a number and count the digits in it.
func ex7() {
	var num int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)
	count := func(n int) (i int) {
		if n == 0 {
			return 1
		}

		for n > 0 {
			n = n / 10
			i++
		}
		return
	}

	fmt.Println(num, "is", count(num), "digit number.")
}

// Find the whether a number is palindrome or not.
func ex6() {
	var num int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	rev := func(n int) (r int) {
		for n > 0 {
			rem := n % 10
			n = n / 10
			r = (r * 10) + rem
		}
		return
	}

	if num == rev(num) {
		fmt.Printf("%d is palindrome.\n", num)
	} else {
		fmt.Println(num, "is not palindrome.")
	}
}

// Enter a number and find the reverse of that number. Also display the double of the reverse number.
func ex5() {
	var num int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)
	fmt.Printf("\nEntered number is  : %d\n", num)
	a := 0
	for num > 0 {
		rem := num % 10
		num = num / 10
		a = (a * 10) + rem
	}
	fmt.Printf("Reversed number is : %d\nDoubled number is  : %d\n", a, a*2)
}

// Print all prime numbers from 1 to 99.
func ex4() {
	num := 99
	for i := 1; i <= num; i++ {
		var j int
		for j = 2; j < num-1; j++ {
			if i%j == 0 {
				break
			}
		}
		if i == j {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

// Enter 10 numbers ranging from 1 to 50, and draw a histogram by displaying adjacent '=' signs for each number entered. For example if the number entered is 12 then a line of 12 equal-to '=' sign should be displayed.
func ex3() {
	numbers := make([]int, 10)
	fmt.Printf("Enter 10 numbers: ")
	for i := 0; i < len(numbers); i++ {
		fmt.Scanf("%d", &numbers[i])
		fmt.Printf("%4d | ", numbers[i])
		for j := 0; j < numbers[i]; j++ {
			fmt.Printf("=")
		}
		fmt.Println()
	}

}

// Input 10 numbers and find the largest number.
func ex2() {
	numbers := make([]int, 10)
	fmt.Printf("Enter your 10 numbers by space sperated: ")
	max := 0
	for i := 0; i < len(numbers); i++ {
		fmt.Scanf("%d", &numbers[i])
		if max < numbers[i] {
			max = numbers[i]
		}
	}
	fmt.Println("max number is: ", max)
}

// Print numbers from 1 to 80 separated by tab, 8 numbers per line.
func ex1() {
	num := 80

	for count := 1; count <= num; {
		for i := 0; i < 8; i++ {
			fmt.Printf("%d ", count)
			count++
		}
		fmt.Println()
	}
}
