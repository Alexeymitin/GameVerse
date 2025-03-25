package main

import (
	"gameverse/files"
)

func main() {
	files.ReadFile()
	files.WriteFile("privet! fail", "file.txt")
}
