package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type IpItem struct {
	Time string `json:"time"`
	Ip   string `json:"ip"`
	Url  string `json:"url"`
}

func ReadLogToMap(fileName string) (error, map[string]int) {
	// открываем файл
	file, err := os.Open(fileName)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	// создаем сканер для чтения файла
	scanner := bufio.NewScanner(file)
	ipAdresses := make(map[string]int, 16)

	// читаем файл построчно
	var ipItem IpItem
	for scanner.Scan() {
		rawJson := scanner.Text()

		err = json.Unmarshal([]byte(rawJson), &ipItem)
		if err != nil {
			return err, nil
		}

		_, exists := ipAdresses[ipItem.Ip]

		if exists {
			ipAdresses[ipItem.Ip]++
		} else {
			ipAdresses[ipItem.Ip] = 1
		}
	}

	fmt.Println(ipAdresses)
	
	return scanner.Err(), ipAdresses
}

func WriteMapToResult(fileName string, logs map[string]int) error {
	fResult, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	defer fResult.Close()

	for key, value := range logs {
		fmt.Println(key, value)
		fmt.Fprintf(fResult, "%s %d\n", key, value)
	}
	return nil
}

func main() {
	err, logs := ReadLogToMap("access.log")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = WriteMapToResult("result.txt", logs)
	if err != nil {
		log.Fatal(err.Error())
	}
}
