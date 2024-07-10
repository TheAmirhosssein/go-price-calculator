package filemanger

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return nil, err
	}

	scan := bufio.NewScanner(file)
	var fileLines []string
	for scan.Scan() {
		fileLines = append(fileLines, scan.Text())
	}
	file.Close()

	if scan.Err() != nil {
		fmt.Println(scan.Err())
		return nil, err
	}

	return fileLines, nil
}
