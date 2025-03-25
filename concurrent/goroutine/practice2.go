package goroutine

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 所指向页面的 body 内容，
	// 并将该页面上找到的所有 URL 放到一个切片中。
	Fetch(url string) (body string, urls []string, err error)
}

type visited struct {
	mu   sync.Mutex
	urls map[string]bool
}

func (v *visited) visited(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.urls[url]
}

func (v *visited) markVisited(url string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.urls[url] = true
}

var v = visited{urls: make(map[string]bool)}

// Crawl 用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {

	// 下面并没有实现上面两种情况：
	//if depth <= 0 {
	//	return
	//}
	//body, urls, err := fetcher.Fetch(url)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("found: %s %q\n", url, body)
	//for _, u := range urls {
	//	Crawl(u, depth-1, fetcher)
	//}

	// TODO: 并行地爬取 URL。
	// TODO: 不重复爬取页面。
	if depth <= 0 {
		return
	}

	if v.visited(url) {
		return
	}
	v.markVisited(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found:%s %q\n", url, body)

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher)
		}(u)
	}
	wg.Wait()
	return
}

// fakeFetcher 是待填充结果的 Fetcher。
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

func WebTest() {
	Crawl("https://golang.org/", 4, fetcher)
}
