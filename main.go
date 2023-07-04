package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
}

func main() {
	// read standard.txt and convert to array of lines
	readFile, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile) // creates a scanner (fileScanner) to read the contents of the opened file.
	fileScanner.Split(bufio.ScanLines)        // configures the scanner to split the file content into lines based on newline characters ('\n').
	var fileLines []string                    // declares an empty slice (fileLines) to store the lines read from the file.
	for fileScanner.Scan() {                  // enters a loop that iterates until the scanner reaches the end of the file. In each iteration, it reads the next line
		fileLines = append(fileLines, fileScanner.Text()) // from the file using fileScanner.Scan() and adds the line to the fileLines slice using fileScanner.Text()
	}
	readFile.Close()

	if len(os.Args) < 3 || len(os.Args) > 4 { // if the number of command lines arguments are not within correct range
		fmt.Println("Error: Invalid number of arguments")
		printUsage()
		os.Exit(1)
	}

	colorFlag := os.Args[1]
	if !strings.HasPrefix(colorFlag, "--color=") { // if the prefix of 1st argument is not --color=
		fmt.Println("Error: Invalid format for --color flag. Please use --color=<color>")
		printUsage()
		os.Exit(2)
	}
	colorFlag = strings.TrimPrefix(colorFlag, "--color=") // get rid of --color=

	var lettersToColor, text string

	if len(os.Args) == 3 {
		text = os.Args[2]
	} else {
		lettersToColor = os.Args[2]
		text = os.Args[3]
	}

	textSlice := strings.Split(text, " ")

	// works with:     --color=red "hello world" "hello world"    or    --color=red "hello world"    or    --color=red "hello"
	if text == os.Args[2] {
		fmt.Println("process1Variable") // these print lines are just for clarity of variable values and which function was called
		fmt.Println("lettersToColor: ", lettersToColor)
		fmt.Println("text: ", text)
		process1Variable(text, text, colorFlag, fileLines)
	}

	// works when LettersToColor is equal as all slices of 'text'    --color=red "hello" "hello hello"
	LTCequalToSecVarSlices := false
	if len(textSlice) > 1 {
		count := 0
		for i := 0; i < len(textSlice); i++ {
			if lettersToColor != text && lettersToColor == textSlice[i] {
				count++
			}
		}
		if count == len(textSlice) {
			LTCequalToSecVarSlices = true
		}
	}
	if lettersToColor != text && LTCequalToSecVarSlices && lettersToColor == os.Args[2] {
		fmt.Println("process1Variable for LTCequalToSecVarSlices") // these print lines are just for clarity of variable values and which function was called
		fmt.Println("textslice length: ", len(textSlice))
		fmt.Println("lettersToColor: ", lettersToColor)
		fmt.Println("LTCequalToSecVarSlices: ", LTCequalToSecVarSlices)
		fmt.Println("text: ", text)
		process1Variable(text, lettersToColor, colorFlag, fileLines)
	}

	// works with:     --color=orange GuYs "HeY GuYs"      but not    --color=red hello "hello world"    and not with 3 words 'text'
	var matchingWord string
	for i := 0; i < len(textSlice); i++ {
		if lettersToColor == textSlice[i] {
			matchingWord = textSlice[i]
			var lenOfAdjacWord int
			if len(textSlice) == 1 {
				break
			}
			if lettersToColor == textSlice[0] {
				break
			} else {
				lenOfAdjacWord = len(textSlice[i-1])
			}
			fmt.Println("processMatchingWord") // these print lines are just for clarity of variable values and which function was called
			fmt.Println("lettersToColor: ", lettersToColor)
			fmt.Println("text: ", text)
			processMatchingWord(text, lettersToColor, colorFlag, fileLines, lenOfAdjacWord)
		}
	}

	/* works with command line variables that are not same i.e.:
	   Try specifying set of letters to be colored (the second until the last letter). --color=blue ram "Aram"
	   Try specifying letter to be colored (the second letter).             --color=blue r "Aram"
	   Try specifying a set of letters to be colored (just two letters).	--color=blue rm "Aram"  */
	if lettersToColor != text && lettersToColor != matchingWord {
		fmt.Println("processNotEqualVariables") // these print lines are just for clarity of variable values and which function was called
		fmt.Println("lettersToColor: ", lettersToColor)
		fmt.Println("text: ", text)
		processNotEqualVariables(text, lettersToColor, colorFlag, fileLines)
	}
}

// to handle/colorise single command line variable or two equal variables
func process1Variable(text string, lettersToColor string, colorFlag string, fileLines []string) {
	/* to work with "\n" as new line, uncomment below code, replace line 122 with comment and use closing brackets at the bottom of this function
	preLine := []rune(text)
	for m := 0; m < len(preLine); m++ {
		arrayMiddle := "n3wL!Ne"
		if preLine[m] == 92 && preLine[m+1] == 'n' {
			array1 := preLine[0:m]
			array2 := preLine[m+2:]
			s1 := string([]rune(array1))
			s2 := string([]rune(array2))
			text = s1 + arrayMiddle + s2
			preLine = []rune(text)
		}
	}

	line := strings.Split(string(preLine), "n3wL!Ne")
	for i := 0; i < len(line); i++ {
		if len(text) < 1 {
			break
		}
		if len(line[i]) < 1 && i == 0 {
			continue
		}
		if len(line[i]) < 1 {
			fmt.Println()
			continue
		}
	*/
	textSlice := []rune(text) // textSlice := []rune(line[i])
	for j := 1; j < 9; j++ {
		for k := 0; k < len(textSlice); k++ {
			asciiFetch := ((textSlice[k] - 32) * 9) + rune(j)
			fmt.Printf("%s", colorize(fileLines[asciiFetch], colorFlag))
		}
		fmt.Println()
	}
}

// to handle/colorise matching word in 'text' variable that is the same as 'lettersToColor'
func processMatchingWord(text string, lettersToColor string, colorFlag string, fileLines []string, lenOfAdjacWord int) {
	textSlice := []rune(text)
	for j := 1; j < 9; j++ {
		for k := 0; k < len(textSlice); k++ {
			asciiFetch := ((textSlice[k] - 32) * 9) + rune(j)
			letters := lenOfAdjacWord + 1
			if k == letters || (k >= lenOfAdjacWord && k <= letters+lenOfAdjacWord+1) {
				fmt.Printf("%s", colorize(fileLines[asciiFetch], colorFlag))
				letters++
			} else {
				fmt.Print(fileLines[asciiFetch])
			}
		}
		fmt.Println()
	}
}

// to match the characters in 'lettersToColor' with letters in 'text' when there is no same/matching words
func processNotEqualVariables(text string, lettersToColor string, colorFlag string, fileLines []string) {
	textSlice := []rune(text)
	for j := 1; j < 9; j++ {
		for k := 0; k < len(textSlice); k++ {
			asciiFetch := ((textSlice[k] - 32) * 9) + rune(j)
			if strings.ContainsRune(lettersToColor, textSlice[k]) {
				fmt.Printf("%s", colorize(fileLines[asciiFetch], colorFlag))
			} else {
				fmt.Print(fileLines[asciiFetch])
			}
		}
		fmt.Println()
	}
}

// It creates a colorMapping map that maps color flags to their corresponding ANSI escape code formats.
// Each color flag is associated with a specific escape code that sets the text color in the terminal.
func colorize(text string, colorFlag string) string {
	colorMapping := map[string]string{
		"black":   "\033[30m%s\033[0m", // The %s placeholder in the escape code format is replaced with the text value, resulting in the colorized version of the text.
		"red":     "\033[31m%s\033[0m",
		"green":   "\033[32m%s\033[0m",
		"yellow":  "\033[33m%s\033[0m",
		"blue":    "\033[34m%s\033[0m",
		"purple":  "\033[35m%s\033[0m",
		"magenta": "\033[35m%s\033[0m",
		"cyan":    "\033[36m%s\033[0m",
		"white":   "\033[37m%s\033[0m",
		"orange":  "\033[38;5;208m%s\033[0m",
		"gray":    "\033[90m%s\033[0m",
	}

	/* It attempts to retrieve the escape code format for the given colorFlag from the colorMapping map using colorMapping[colorFlag].
	The second variable found is a boolean flag indicating whether the colorFlag was found in the map. */
	format, found := colorMapping[colorFlag] // If the colorFlag was found in the map, the corresponding escape code format is retrieved into the format variable
	if !found {                              // If the colorFlag was not found in the map,
		return text // the original text is returned as is, without any color formatting.
	}
	return fmt.Sprintf(format, text) // The fmt.Sprintf function is used to format the text using the format escape code.
}
