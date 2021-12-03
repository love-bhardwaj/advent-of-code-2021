package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := readData(file)

	fmt.Println(checkIncrease(data))
	fmt.Println(getIncreaseWindow(data))
}

func checkIncrease(data []int) int{
	count := 0
	for index := range data {
		if index > 0 {
			if data[index] > data[index-1]{
				count++;
			}
		}
	}
	return count
}

func getIncreaseWindow(data []int) int{
	count := 0
	for index := range data {
		if index+4 > len(data) {
			break
		}
		currWin := data[index:index+3]
		nextWindow := data[index+1:index+4]
		if sumSlice(nextWindow) > sumSlice(currWin) {
			count++;
		}
	}
	return count
}

func sumSlice(data []int) int {
	sum := 0
	if len(data) < 3{
		return sum
	} 

	for _, value := range data {
		sum += value
	}
	return sum
}

func readData(file *os.File) []int{
	var data []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataValue, _ := strconv.Atoi(scanner.Text())
		data = append(data, dataValue)
	}
	return data;
}