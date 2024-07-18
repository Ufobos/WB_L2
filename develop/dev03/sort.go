package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type sortOptions struct {
	column  int
	numeric bool
	reverse bool
	unique  bool
}

func parseArgs() sortOptions {
	var opts sortOptions

	flag.IntVar(&opts.column, "k", 0, "column number to sort by (default 0, meaning the whole line)")
	flag.BoolVar(&opts.numeric, "n", false, "sort by numeric value")
	flag.BoolVar(&opts.reverse, "r", false, "sort in reverse order")
	flag.BoolVar(&opts.unique, "u", false, "output only unique lines")

	flag.Parse()

	return opts
}

func readLines(filePath string) ([]string, error) {
	fmt.Println("Reading lines from file:", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeLines(lines []string, filePath string) error {
	fmt.Println("Writing lines to file:", filePath)
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	fmt.Println("Absolute path for output file:", absPath)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Попробуем сразу записать в файл, чтобы проверить его создание и запись
	_, err = file.WriteString("Test write to file\n")
	if err != nil {
		return err
	}
	fmt.Println("Successfully wrote test line to file:", filePath)

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func sortLines(lines []string, opts sortOptions) []string {
	fmt.Println("Sorting lines with options:", opts)
	if opts.column > 0 {
		sort.SliceStable(lines, func(i, j int) bool {
			wordsI := strings.Fields(lines[i])
			wordsJ := strings.Fields(lines[j])
			if len(wordsI) <= opts.column-1 || len(wordsJ) <= opts.column-1 {
				return lines[i] < lines[j]
			}
			if opts.numeric {
				numI, errI := strconv.ParseFloat(wordsI[opts.column-1], 64)
				numJ, errJ := strconv.ParseFloat(wordsJ[opts.column-1], 64)
				if errI == nil && errJ == nil {
					if opts.reverse {
						return numI > numJ
					}
					return numI < numJ
				}
			}
			if opts.reverse {
				return wordsI[opts.column-1] > wordsJ[opts.column-1]
			}
			return wordsI[opts.column-1] < wordsJ[opts.column-1]
		})
	} else {
		sort.SliceStable(lines, func(i, j int) bool {
			if opts.numeric {
				numI, errI := strconv.ParseFloat(lines[i], 64)
				numJ, errJ := strconv.ParseFloat(lines[j], 64)
				if errI == nil && errJ == nil {
					if opts.reverse {
						return numI > numJ
					}
					return numI < numJ
				}
			}
			if opts.reverse {
				return lines[i] > lines[j]
			}
			return lines[i] < lines[j]
		})
	}

	if opts.unique {
		uniqueLines := make([]string, 0, len(lines))
		seen := make(map[string]struct{})
		for _, line := range lines {
			if _, exists := seen[line]; !exists {
				seen[line] = struct{}{}
				uniqueLines = append(uniqueLines, line)
			}
		}
		return uniqueLines
	}

	return lines
}

func main() {
	opts := parseArgs()
	if len(flag.Args()) != 2 {
		fmt.Println("Usage: sort [options] inputfile outputfile")
		os.Exit(1)
	}

	inputFile := flag.Args()[0]
	outputFile := flag.Args()[1]

	// Вывод текущего рабочего каталога
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting working directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Current working directory:", workingDir)

	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Lines read from file:", lines)

	sortedLines := sortLines(lines, opts)

	fmt.Println("Sorted lines:", sortedLines)

	if err := writeLines(sortedLines, outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Output written to file:", outputFile)
}
