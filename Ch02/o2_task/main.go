package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1 = ")
	str1, _ := reader.ReadString('\n')
	value1, err := getFloat(str1)
	if err != nil {
		fmt.Println("Value 1 should be a float:", err)
		return
	}

	fmt.Print("Value 2 = ")
	str2, _ := reader.ReadString('\n')
	value2, err := getFloat(str2)
	if err != nil {
		fmt.Println("Value 2 should be a float:", err)
		return
	}

	sum := value1 + value2
	sum = math.Round(sum*100) / 100
	fmt.Printf("Sum = %.2f\n", sum)
}

func getFloat(str string) (float64, error) {
	value, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}
