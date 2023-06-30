package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Flag package not used so missing '=' triggers an error. But because code modification, /n doesn't work"

func main() {
	// read standard.txt and convert to array of lines
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

	if len(os.Args) < 3 || len(os.Args) > 4 {
		printUsage()
		os.Exit(1)
	}

	colorFlag := os.Args[1]
	if !strings.HasPrefix(colorFlag, "--color=") {
		fmt.Println("Error: Invalid format for --color flag. Please use --color=<color>")
		printUsage()
		os.Exit(1)
	}
	colorFlag = strings.TrimPrefix(colorFlag, "--color=")

	lettersToColor := os.Args[2]
	text := ""

	if len(os.Args) == 4 {
		text = os.Args[3]
	} else {
		text = lettersToColor
	}

	lenOfPrevWord := -1
	if strings.Contains(text, lettersToColor) {
		textSlice := strings.Split(text, " ")
		for i, matchingWord := range textSlice {
			if lettersToColor == matchingWord {
				lenOfPrevWord = len(textSlice[i-1])
			}
		}
		processTextWord(text, lettersToColor, colorFlag, fileLines, lenOfPrevWord)
	} else {
		processText(text, lettersToColor, colorFlag, fileLines)
	}
}

func processTextWord(text string, lettersToColor string, colorFlag string, fileLines []string, lenOfPrevWord int) {
	textSlice := []rune(text)
	for j := 1; j < 9; j++ {
		for k := 0; k < len(textSlice); k++ {
			m := rune(j)
			asciiFetch := ((textSlice[k] - 32) * 9) + m
			letters := lenOfPrevWord + 1
			if k == letters {
				fmt.Printf("%s", colorize(fileLines[asciiFetch], colorFlag))
				letters++
			} else {
				fmt.Print(fileLines[asciiFetch])
			}
		}
		fmt.Println()
	}
}

func processText(text string, lettersToColor string, colorFlag string, fileLines []string) {
	nextStep := text
	line := strings.Split(nextStep, "\n")

	for j := 0; j < len(line); j++ {
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

		for k := 1; k < 9; k++ {
			for i := 0; i < len(word); i++ {
				m := rune(k)
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

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
}
