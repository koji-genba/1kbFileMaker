package main

import (
    "fmt"
    "os"
    "strconv"
    "sync"
	"time"
	"flag"
)

func createFiles(threadNum int, wg *sync.WaitGroup, fileNum int) {
    defer wg.Done()
	fmt.Println("thread" + strconv.Itoa(threadNum))
    for i := 0; i < fileNum; i++ {
        fileName := "thread_" + strconv.Itoa(threadNum) + "_file_" + strconv.Itoa(i) + ".txt"
        file, err := os.Create(fileName)
        if err != nil {
            fmt.Println("Error creating file:", err)
            return
        }
        data := make([]byte, 1024)
        file.Write(data)
        file.Close()
    }
}

func main() {
	start := time.Now()

	var (
		f = flag.Int("f", 1, "file/thread")
		t = flag.Int("t", 1, "thread")
	)
	flag.Parse()
	fmt.Println("param f : ", *f)
	fmt.Println("param t : ", *t)

    var wg sync.WaitGroup
    for i := 0; i < *t; i++ {
        wg.Add(1)
        go createFiles(i, &wg, *f)
    }
    wg.Wait()

	fmt.Printf("process time: %s\n", time.Since(start))
    fmt.Println("All files created.")
}
