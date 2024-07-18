package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

var after int
var before int
var contextFlag int
var countFlag bool
var ignoreCaseFlag bool
var invertFlag bool
var fixedFlag bool
var lineNumFlag bool

func init() {
	flag.IntVar(&after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&contextFlag, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&countFlag, "c", false, "подсчет количества строк с совпадениями")
	flag.BoolVar(&ignoreCaseFlag, "i", false, "игнорировать регистр")
	flag.BoolVar(&invertFlag, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&fixedFlag, "F", false, "точное совпадение со строкой")
	flag.BoolVar(&lineNumFlag, "n", false, "напечатать номер строки")
}

func main() {
	flag.Parse()

	if after == 0 && contextFlag != 0 {
		after = contextFlag
	}
	if before == 0 && contextFlag != 0 {
		before = contextFlag
	}

	pattern, fileName := flag.Arg(0), flag.Arg(1)
	if len(pattern) == 0 || len(fileName) == 0 {
		log.Fatal("Введите паттерн и название файла в формате: PATTERN FILE")
	}

	if fixedFlag {
		pattern = regexp.QuoteMeta(pattern)
	}
	if ignoreCaseFlag {
		pattern = `(?i)` + pattern
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Ошибка в паттерне: %s", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %s", err)
	}
	defer file.Close()

	inputStrings := readLines(file)

	if countFlag {
		count := countEntries(inputStrings, re)
		fmt.Printf("Количество совпадений: %d\n", count)
		return
	}
	matchLines := findMatch(inputStrings, re)
	printResult(matchLines, inputStrings)
}

func readLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка чтения файла: %s", err)
	}
	return lines
}

func countEntries(strs []string, r *regexp.Regexp) int {
	count := 0
	for _, line := range strs {
		if r.MatchString(line) != invertFlag {
			count++
		}
	}
	return count
}

func findMatch(strs []string, r *regexp.Regexp) []int {
	matchLines := make([]int, 0)
	for idx, line := range strs {
		if r.MatchString(line) != invertFlag {
			matchLines = append(matchLines, idx)
		}
	}
	return matchLines
}

func printResult(data []int, strs []string) {
	printedLines := make(map[int]bool)
	for _, numberLine := range data {
		if before > 0 || after > 0 {
			printContext(numberLine, before, after, strs, printedLines)
		} else if !printedLines[numberLine] {
			printLine(numberLine, strs[numberLine], numberLine+1)
			printedLines[numberLine] = true
		}
	}
}

func printContext(numberLine, before, after int, strs []string, printedLines map[int]bool) {
	start := max(0, numberLine-before)
	finish := min(len(strs)-1, numberLine+after)
	for i := start; i <= finish; i++ {
		if !printedLines[i] {
			printLine(i, strs[i], i+1)
			printedLines[i] = true
		}
	}
}

func printLine(_ int, line string, lineNum int) {
	if lineNumFlag {
		fmt.Printf("%d:%s\n", lineNum, line)
	} else {
		fmt.Println(line)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
