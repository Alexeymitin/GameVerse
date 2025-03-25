package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	data, err := os.ReadFile("file.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

}

func WriteFile(content, name string) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.WriteString(content)

	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}

	fmt.Println("Write good")
	file.Close()
}
