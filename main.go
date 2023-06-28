package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// do poprawek bo teraz przy  go run . --color=red aram "aram bravo" koloruje na czerwono aram *ra**
func main() {
	// reading standard.txt and convert to array of lines
	readFile, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	// get argument as a string
	flag.Usage = func() {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
	}
	var colorFlag string
	flag.StringVar(&colorFlag, "color", "", "Specify the color for highlighting")
	flag.Parse()

	// will print above instructions in func if to many inputs
	if len(os.Args) > 4 { 
		flag.Usage()
		os.Exit(1)
	}
	
	lettersToColor := flag.Arg(0) //letters/word that we want to be in color if 2 strings are provided/inputed  // fmt.Println("flag.Arg(0):", lettersToColor)

	text := flag.Arg(1) // whole string          // fmt.Println("flag.Arg(1):", text)

	/* check wether lettersToColor is substring of text
	if strings.Contains(text, lettersToColor) {
		fmt.Println("true")
	}
	*/
	
	// if there is no letters/word provided/inputed, first string becomes the text to be colored
	if len(os.Args) == 3 {
		text = flag.Arg(0)
		lettersToColor = text
	} 
	
/*
	textSlice := strings.Split(text, " ")
	for _, word := range textSlice {
		if word == lettersToColor {
			//fmt.Println("yes")
		}
	}
*/
	
	// looking for "\n" and turn it into "n3wL1ne" so string.Split can find it
	preLine := []rune(text)
	for m := 0; m < len(preLine); m++ {
		arrayMiddle := "n3wL!Ne"
		if preLine[m] == 92 && preLine[m+1] == 'n' {
			array1 := preLine[0:m]
			array2 := preLine[m+2:]
			s1 := string([]rune(array1))
			s2 := string([]rune(array2))
			text = s1 + arrayMiddle + s2
			preLine = ([]rune(text))
		}
	}

	// split the text into lines if required
	nextStep := string(preLine)
	line := strings.Split(nextStep, "n3wL!Ne")

	// loop to work on lines
	for j := 0; j < len(line); j++ {

		// to make or not make new lines in situations with no other text
		if len(text) < 1 {
			break
		}
		if len(line[j]) < 1 && j == 0 {
			continue
		}
		if len(line[j]) < 1 {
			fmt.Println()
			continue
		}

		word := []rune(line[j])

		// row by row loop
		for k := 1; k < 9; k++ {

			// character by character loop
			for i := 0; i < len(word); i++ {
				m := rune(k)

				// reduce each character value by 32 in ascii table,
				// multiply by the 9 rows each character uses in standard.txt,
				// add the row number
				asciiFetch := ((word[i] - 32) * 9) + m

				if strings.ContainsRune(lettersToColor, word[i]) {
					fmt.Printf("%s", colorize(fileLines[asciiFetch], colorFlag))
				} else {
					fmt.Print(fileLines[asciiFetch])
				}
			}
			fmt.Println()
		}
	}

}
// colorize applies the specified color to the text
func colorize(text string, colorFlag string) string {
	colorMapping := map[string]string{
		"red":    "\033[31m%s\033[0m",
		"green":  "\033[32m%s\033[0m",
		"yellow": "\033[33m%s\033[0m",
		"blue":   "\033[34m%s\033[0m",
		"purple": "\033[35m%s\033[0m",
		"cyan":   "\033[36m%s\033[0m",
		"white":  "\033[37m%s\033[0m",
	}

	format, found := colorMapping[colorFlag]
	if !found {
		return text
	}
	return fmt.Sprintf(format, text)
}