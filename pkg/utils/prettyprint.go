package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(s interface{}) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
