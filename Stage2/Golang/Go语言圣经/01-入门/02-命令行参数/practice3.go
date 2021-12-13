/*
	做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
*/
package main

import (
	"fmt"
	"os"
	"time"
)

func joinEcho() {
	start := time.Now()
	fmt.Printf("joinEcho: %v\n", time.Since(start))
}

func plusEcho() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("plusEcho: %v\n", time.Since(start))
}

func main() {
	joinEcho()
	plusEcho()
}
