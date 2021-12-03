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

func readData(file *os.File) []int{
	var data []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataValue, _ := strconv.Atoi(scanner.Text())
		data = append(data, dataValue)
	}
	return data;
}