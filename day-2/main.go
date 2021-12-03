package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Command string
	Value int
}

type FinalPosition struct {
	Horizontal int
	Depth int
}

func main(){
		file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := readData(file)
	finalPosition := processData(data)
	fmt.Println(finalPosition.Depth * finalPosition.Horizontal)
	finalPostionWithAim := processDataWithAim(data)
	fmt.Println(finalPostionWithAim.Depth * finalPosition.Horizontal)
}

func processData(data []Input) FinalPosition{
	finalPosition := FinalPosition{0, 0}
	for _, dataValue := range data {
		switch dataValue.Command {
		case "forward":
			finalPosition.Horizontal += dataValue.Value
		case "up":
			finalPosition.Depth += dataValue.Value
		case "down":
			finalPosition.Depth -= dataValue.Value
		}
	}
	return finalPosition
}

func processDataWithAim(data []Input) FinalPosition {
	finalPosition := FinalPosition{0, 0}
	aim := 0
		for _, dataValue := range data {
		switch dataValue.Command {
		case "forward":
			finalPosition.Horizontal += dataValue.Value
			finalPosition.Depth += (aim * dataValue.Value)
		case "up":
			// finalPosition.Depth -= dataValue.Value
			aim -= dataValue.Value
		case "down":
			// finalPosition.Depth += dataValue.Value
			aim += dataValue.Value
		}
	}
	return finalPosition
}

func readData(file *os.File) []Input{
	var data []Input

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		dataValue := scanner.Text()
		s := strings.Split(dataValue, " ")
		command := s[0]
		value, _ := strconv.Atoi(s[1])
		input := Input{command, value}
		data = append(data, input)
	}
	return data
}