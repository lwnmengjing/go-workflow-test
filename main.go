package main

import (
	"log"
	"os/exec"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/9/27 17:17:12
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/9/27 17:17:12
 */

func main() {
	err := exec.Command("echo", "\"name=test1\"", ">>", "$GITHUB_OUTPUT").Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
