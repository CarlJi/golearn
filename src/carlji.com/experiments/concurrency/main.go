package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

type Scenario struct {
	Name        string
	Description []string
	Examples    []string
	RunExample  func()
}

var s1 = &Scenario{
	Name: "s1",
	Description: []string{
		"无脑并发执行任务",
	},
	Examples: []string{
		"比如并发的请求后端某个接口",
	},
	RunExample: RunScenario1,
}

var s2 = &Scenario{
	Name: "s2",
	Description: []string{
		"按数量并行执行任务, goroutine worker pool",
	},
	Examples: []string{
		"比如技术支持要给某个客户删除几个TB/GB的文件",
	},
	RunExample: RunScenario2,
}

var s3 = &Scenario{
	Name: "s3",
	Description: []string{
		"按时间来持续并发",
	},
	Examples: []string{
		"在规定时间内，持续的高并发请求后端服务， 防止服务死循环",
	},
	RunExample: RunScenario3,
}

var Scenarios []*Scenario

func init() {
	Scenarios = append(Scenarios, s1)
	Scenarios = append(Scenarios, s2)
	Scenarios = append(Scenarios, s3)
}

// 常用的并发与同步场景
func main() {
	if len(os.Args) == 1 {
		fmt.Println("=============>")
		return
	}
	for _, arg := range os.Args[1:] {
		sc := matchScenario(arg)
		if sc != nil {
			printDescription(sc.Description)
			printExamples(sc.Examples)
			sc.RunExample()
		}
	}
}

func printDescription(str []string) {
	fmt.Printf("场景描述: %s \n", str)
}

func printExamples(str []string) {
	fmt.Printf("场景举例: %s \n", str)
}

func matchScenario(name string) *Scenario {
	for _, sc := range Scenarios {
		if sc.Name == name {
			return sc
		}
	}
	return nil
}

var doSomething = func(i int) {
	fmt.Printf("Goroutine %d do things .... \n", i)
	time.Sleep(time.Millisecond * time.Duration(10))
}

// 场景1: 按数量，并发做事/发送请求

func RunScenario1() {
	count := 10
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			doSomething(index)
		}(i)
	}

	wg.Wait()
}

// 场景2：以 worker pool 方式 并发做事/发送请求

func RunScenario2() {
	numOfConcurrency := runtime.NumCPU()
	taskTool := 10
	jobs := make(chan int, taskTool)
	results := make(chan int, taskTool)
	var wg sync.WaitGroup

	// workExample
	workExampleFunc := func(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for job := range jobs {
			res := job * 2
			fmt.Printf("Worker %d do things, produce result %d \n", id, res)
			time.Sleep(time.Millisecond * time.Duration(100))
			results <- res
		}
	}

	for i := 0; i < numOfConcurrency; i++ {
		wg.Add(1)
		go workExampleFunc(i, jobs, results, &wg)
	}

	totalTasks := 100

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < totalTasks; i++ {
			n := <-results
			fmt.Printf("Got results %d \n", n)
		}
		close(results)
	}()

	for i := 0; i < totalTasks; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

// 场景3: 按时间来持续并发

func RunScenario3() {
	timeout := time.Now().Add(time.Second * time.Duration(10))
	n := runtime.NumCPU()

	waitForAll := make(chan struct{})
	done := make(chan struct{})
	concurrentCount := make(chan struct{}, n)

	for i := 0; i < n; i++ {
		concurrentCount <- struct{}{}
	}

	go func() {
		for time.Now().Before(timeout) {
			<-done
			concurrentCount <- struct{}{}
		}

		waitForAll <- struct{}{}
	}()

	go func() {
		for {
			<-concurrentCount
			go func() {
				// do something
				done <- struct{}{}
			}()
		}
	}()

	<-waitForAll
}

// 场景4: Goroutine 间的等待与同步

func scenario4() {
	go func() {

	}()

	go func() {

	}()
}

//

// 场景5： 在规定时间被等待一个异步Gotoutine返回结果

func scenario5() {
	go func() {

	}()

	select {}
}

// 场景6： 在规定的时间内定时做事情

func scenario6() {

}
