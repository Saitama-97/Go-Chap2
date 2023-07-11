package search

import (
	"fmt"
	"log"
)

/**
 * @Time    : 2023/6/25 16:33
 * @File    : match.go
 * @Project : chapter2
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    :
 */

// Result 结构体，保存搜索结果
type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 函数，为每个数据源单独启动 goroutine 来执行这个函数
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	// 将搜索结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}

// Display 从每个单独的 goroutine 接收到结果后打印在终端窗口
func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
