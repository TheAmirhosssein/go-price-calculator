package filemanger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type stringFloatMap map[string]float64

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

func WriteJson(data stringFloatMap, key string) error {
	jsonData, err := os.ReadFile("result.json")

	if os.IsNotExist(err) {
		startData := make([]map[string]stringFloatMap, 1)
		startData[0] = make(map[string]stringFloatMap)
		startData[0][key] = data
		jsonData, _ = json.Marshal(startData)
		os.WriteFile("result.json", []byte(jsonData), 0644)

	}

	var fileData []map[string]stringFloatMap

	err = json.Unmarshal(jsonData, &fileData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}
	for _, jsonItem := range fileData {
		jsonItem[key] = data
	}
	jsonResultData, err := json.Marshal(fileData)
	os.WriteFile("result.json", []byte(jsonResultData), 0644)
	return err
}
