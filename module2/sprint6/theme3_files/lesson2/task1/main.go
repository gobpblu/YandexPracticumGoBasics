package main

import (
	"log"
	"os"
	"time"
)

func main() {

	// открываем файл
	fRun, err := os.OpenFile("run.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// не забываем о закрытии файла
	defer fRun.Close()

	// ограничим количество действий
	fRun.WriteString(time.Now().Format("02.01.06 15.04.05") + "\n")

}
