/*
	修改dup2，出现重复的行时打印文件名称。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		_, ok := counts[key]
		if ok {
			fmt.Println("当前文件名：", f.Name())
			counts[input.Text()]++
		} else {
			counts[input.Text()]++
		}
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("line: %d\tn:%s\n", n, line)
		}
	}
}
