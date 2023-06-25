/**
 * @Time    : 2023/6/25 16:31
 * @File    : main.go
 * @Project : chapter2
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    :
 */

package main

import (
	"chapter2/src/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Fatal("123")
	search.Run("president")
}
