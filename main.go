package main

import (
	"fmt"
	"os"
)

func main() {
	res, _ := os.UserCacheDir()
	fmt.Println(res)
}
