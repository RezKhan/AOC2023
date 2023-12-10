package readfile

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	var lines []string
	for fScanner.Scan() {
		line := fScanner.Text()
		lines = append(lines, line)
	}
	return lines
}
