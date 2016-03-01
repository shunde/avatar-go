package main

import (
	_ "bufio"
	"fmt"
	"github.com/shunde/avatar-go/avatar"
	"image/png"
	"os"
)

func main() {
	//str1 := "hello"
	//str2 := "shunde"
	//str3 := "中国人"

	file, err := os.Create("str1.png")
	if err != nil {
		fmt.Printf("Error: %s", nil)
		return
	}
	defer file.Close()
	//w := bufio.NewWriter(file)
	m := avatar.NewAvatar("张三")
	png.Encode(file, m)
	//png.Encode(w, m)
	//w.Flush()
}
