package data

import (
	_ "encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Version string `json:"version"`
	Author  string `json:"author"`
	Config Config `json:"config"`
}

type Config struct {
	Language string `json:"language"`
}

func Init() {
	_, err := os.Stat("data.json")
	
	if os.IsNotExist(err) {
		// 检查文件是否不存在
		fmt.Println("data.json does not exist, creating a new one...")
	} else if err != nil {
		// 检查是否出现其他错误
		fmt.Println("Error checking data.json:", err)
		return
	} else { 
		// 文件存在
		fmt.Println("data.json exists, loading...")
	}
}