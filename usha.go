package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func generateHash(text string) string {
	h := sha256.New()
	io.WriteString(h, os.Args[1])
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <text>\n", os.Args[0])
		return
	}

	re := regexp.MustCompile(".......?")
	groups := re.FindAllString(generateHash(os.Args[1])[0:40], -1)
	fmt.Println(strings.Join(groups, "-"))
}
