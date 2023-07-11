package search

/**
 * @Time    : 2023/6/25 16:33
 * @File    : default.go
 * @Project : chapter2
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    :
 */

// defaultMatcher 默认匹配器
type defaultMatcher struct {
}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
