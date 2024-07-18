package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var fields int
var delimiter string
var separated bool

func init() {
	flag.IntVar(&fields, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(&delimiter, "d", "\t", "выбрать разделитель")
	flag.BoolVar(&separated, "s", false, "выводить только строки с разделителем")
}

func main() {
	flag.Parse()
	fields = fields - 1
	if fields < 0 {
		log.Fatal("выберете корректный номер поля, номер должен быть положительным целым числом")
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), delimiter)
		if separated {
			if len(words) == 1 {
				continue
			} else {
				if len(words) <= fields {
					fmt.Println(words[0])
				} else {
					fmt.Println(words[fields])
				}
			}
		} else {
			if len(words) <= fields {
				fmt.Println(words[0])
			} else {
				fmt.Println(words[fields])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
