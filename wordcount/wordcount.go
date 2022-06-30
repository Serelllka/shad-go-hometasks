package wordcount

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func parseFiles(args []string) {
	files := args
	mapper := make(map[string]int)
	for _, str := range files {
		data, err := os.ReadFile(filepath.Join("tmp", str))
		if err != nil {
			panic(err)
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			words := strings.Split(line, " ")
			for _, word := range words {
				mapper[word]++
			}
		}
	}

	for word, count := range mapper {
		fmt.Printf("%d\t%s\n", count, word)
	}
}
