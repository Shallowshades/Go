package main

import (
	"fmt"
	"sync"
)

/*
练习：Web 爬虫
在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。

修改 Crawl 函数来并行地抓取 URL，并且保证不重复。

提示：你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！
*/

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// 存储找到的url和body
var (
	m         = make(map[string]int)
	mutex     sync.Mutex
	waitGroup sync.WaitGroup
	//群组等待
	//当添加的任务没有完成时（done（））， wait（） 会一直等待
	//三个方法   Add()  Done()  Wait()
)

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：

	//告诉等待组其中一个完成，对应Add
	defer waitGroup.Done()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	//锁
	mutex.Lock()
	defer mutex.Unlock()

	//判断该url是否已经被遍历过
	if m[url] == 0 {
		m[url]++
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			//等待组添加一个成员
			waitGroup.Add(1)
			go Crawl(u, depth-1, fetcher)
		}
	}
}

func main() {
	waitGroup.Add(1)
	go Crawl("https://golang.org/", 4, fetcher)
	waitGroup.Wait()

	for k := range m {
		fmt.Println(k)
	}
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
