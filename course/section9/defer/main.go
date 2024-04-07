package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./dummy_settings.txt")
	defer file.Close()
	defer ShowMessage("File manipulation is over...")
	if err != nil {
		panic(fmt.Sprintf("Failed to create file : %s", err))
	}

	_, err = file.Write([]byte("DB_HOST=POSTGRES_DEV"))

	if err != nil {
		panic(fmt.Sprintf("Failed to modify file : %s", err))
	}
}

func ShowMessage(msg string) {
	fmt.Println(msg)
}
