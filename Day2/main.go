package main

import (
	"adventdaytwo/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	idcapture   = regexp.MustCompile(`(\d{1,3}:)`)
	cubepattern = regexp.MustCompile(`(\d+ (?:blue|green|red))`)
)

func StringToNumberKeyPair(input string) (int, string) {
	//a function that takes an input of a number in a written string format or a numeric string format and converts it to a number
	digitpattern := regexp.MustCompile(`\d{1,3}`)

	Number_To_Return, _ := strconv.Atoi(digitpattern.FindString(input))
	keyString := strings.Trim(input[slices.Max(digitpattern.FindStringIndex(input)):], " ")

	return Number_To_Return, keyString
}

func StringToSlice(input string) [][]string {
	//A helper function not part of the main solution
	var MasterSlice [][]string

	lines := strings.Split(input, "\n") // Split the text into lines

	for _, line := range lines {
		words := strings.Fields(line) // Split each line into words
		if len(words) > 0 {
			MasterSlice = append(MasterSlice, words)
		}
	}
	return MasterSlice
}

func PartOne(outcomes, contained string) int {
	var RunningSum int
	MyContainedCubeMap := make(map[string]int)
	//lets make a map to compare our values to later
	for _, objs := range StringToSlice(contained) {
		//this step will be slightly different as we know the input string will be a single line
		//we would like to populate our map as such key = red : value =2
		//so we would want to find the pairs by their index
		mycontainedcubefindings := cubepattern.FindAllString(strings.Join(objs, " "), -1)
		for _, inner := range mycontainedcubefindings {
			val, key := StringToNumberKeyPair(inner)
			MyContainedCubeMap[key] = val
		}

	}

	//now we can do the same for our outcomes.
	//but we must consider two things,
	//1: we have to track each id we are working with, so we would want to get that before working
	//2: there can be multi colors for any one id, so we would want to check as we loop until some condition is broken
	for _, objs := range StringToSlice(outcomes) {
		var valueispossible bool
		myobjid := idcapture.FindString(strings.Join(objs, " "))
		myoutcomescubefindings := cubepattern.FindAllString(strings.Join(objs, " "), -1)
		
		for _, inner := range myoutcomescubefindings {
			val, key := StringToNumberKeyPair(inner)
			if MyContainedCubeMap[key] >= val {
				valueispossible = true
			} else {
				valueispossible = false
				break
			}
		}
		if valueispossible == true {
			numbertoadd, _ := StringToNumberKeyPair(myobjid)
			RunningSum += numbertoadd
		}
	}

	return RunningSum
}

func PartTwo(outcomes string)int{
	var RunningSum int
	
	//lets make a map to compare our values to later
	for _, objs := range StringToSlice(outcomes) {
		MyContainedCubeMap := make(map[string][]int)
		//this step will be slightly different as we know the input string will be a single line
		//we would like to populate our map as such key = red : value =2
		//so we would want to find the pairs by their index
		mycontainedcubefindings := cubepattern.FindAllString(strings.Join(objs, " "), -1)
		for _, inner := range mycontainedcubefindings {
			val, key := StringToNumberKeyPair(inner)
			MyContainedCubeMap[key] = append(MyContainedCubeMap[key], val)
		}
		//now lets get the max for each color, before multiplying and adding the results
		RunningSum += slices.Max(MyContainedCubeMap["blue"])*slices.Max(MyContainedCubeMap["red"])*slices.Max(MyContainedCubeMap["green"])

	}
	
	return RunningSum
}
func main() {
	input := `12 red, 13 green, 14 blue`
	fmt.Println("Part One Solution: ",PartOne(utils.PuzzleInput,input))
	fmt.Println("Part One Solution: ",PartTwo(utils.PuzzleInput))
}
