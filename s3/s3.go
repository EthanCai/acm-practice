package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")

		switch text {
		case "":
			os.Exit(0)
		default:
			n, _ := strconv.Atoi(text)
			x, err := calculatex(n)
			if err != nil {
				fmt.Println("X")
			} else {
				fmt.Println(x)
			}
		}
	}
}

func calculatex(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("no x")
	}

	i := -1
	for {
		p := math.Pow(2.0, float64(i))
		if p >= float64(n) {
			break
		}

		i = i + 1
	}

	return i - 1, nil
}
