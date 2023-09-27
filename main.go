package main

import (
	"fmt"
	"log"
	"os"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/9/27 17:17:12
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/9/27 17:17:12
 */

func main() {
	output := "echo \"name=test1\" >> $GITHUB_OUTPUT"
	fmt.Println(output)
	_, err := os.Stdout.WriteString(output + "\n")
	if err != nil {
		log.Fatalf("write string error: %v", err)
	}
}
