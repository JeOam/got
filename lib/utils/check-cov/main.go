package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func main() {
	checkCoverage(100)
}

func checkCoverage(min float64) {
	out, _ := exec.Command("go", "tool", "cover", "-func=coverage.txt").CombinedOutput()
	totalStr := regexp.MustCompile(`(\d+.\d+)%\n\z`).FindSubmatch(out)[1]
	total, _ := strconv.ParseFloat(string(totalStr), 64)
	if total < min {
		fmt.Printf("[lint] Test coverage %.1f%% must >= %.1f%%\n", total, min)
		os.Exit(1)
	}
	fmt.Printf("Test coverage: %.1f%%\n", total)
}
