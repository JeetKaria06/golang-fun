package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func fooHandler(
	response http.ResponseWriter,
	request *http.Request,
) {
	log.Println("fooHandler called")
	response.Header().Set("Content-Type", "text/event-stream")
	response.Write([]byte("hello damn world!"))
}

func readFile(filename string) ([]byte, []byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	buf := make([]byte, 62)

	start := int64(0)
	bytesRead, err := file.ReadAt(buf, start)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}

	fmt.Print(string(buf))
	start = int64(bytesRead)
	buf1 := make([]byte, 62)
	bytesRead1, err := file.ReadAt(buf1, start)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}
	fmt.Println(err) // EOF in case it has reached end of file.
	fmt.Print(string(buf[:bytesRead]))
	return buf[:bytesRead], buf1[:bytesRead1], nil
}

func readFileHandler(
	response http.ResponseWriter,
	request *http.Request,
) {
	// flusher to send the chunk even though writer's
	// internal buffer is not full
	flusher, ok := response.(http.Flusher)
	if !ok {
		log.Println("expected http.ResponseWriter to be an http.Flusher")
	}

	filename := "test.txt"
	data, data1, err := readFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("File Read")

	response.Header().Set("Content-Type", "text/plain; charset=us-ascii")
	response.Header().Set("Connection", "Keep-Alive")
	response.Header().Set("Cache-Control", "no-cache")
	// response.Header().Set("Transfer-Encoding", "chunked")
	response.Header().Set("X-Content-Type-Options", "nosniff") // search more into this

	ticker := time.NewTicker(time.Second)
	go func() {
		cnt := 0
		for t := range ticker.C {
			sendData := data
			if cnt%2 == 1 {
				sendData = data1
			}
			io.WriteString(response, string(sendData))
			flusher.Flush()
			cnt += 1

			if cnt == 2 {
				break
			}

			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()

	fmt.Println("Finished: should return Content-Length: 0 here")
	// response.Header().Set("Content-Length", "0")

	// response.Write(data)
}

func main() {
	// logger := log.Logger{}
	http.Handle("/foo", http.HandlerFunc(fooHandler))
	http.Handle("/read", http.HandlerFunc(readFileHandler))
	// http.Handle()
	log.Println("Printing Logs")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
