package main

import (
	"fmt"

	"github.com/abdukhashimov/sample/storage"
)

func main() {
	storage := storage.New()
	storage.User().Create()
	fmt.Println(5 << 4)
}
