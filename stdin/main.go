package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

const chunkSize = 10

func writeWithChunks(buffer *bytes.Buffer, w io.Writer) error {
	i := 0
	chunk := make([]byte, chunkSize)
	for {
		readN, err := buffer.Read(chunk)
		if readN == 0 || err != nil {
			// finished
			return nil
		}

		writeN, err := w.Write(chunk[:readN])
		if err != nil {
			return err
		}
		if readN != writeN {
			return fmt.Errorf("read %d bytes, wrote %d bytes", readN, writeN)
		}

		i += 1
		log.Printf("wrote %d chunk", i)
		time.Sleep(time.Millisecond * 100)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func readStdout(pout io.ReadCloser) {
	reader := bufio.NewReader(pout)

	var output string
	var err error
	for err == nil {
		output, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("failed to read stdout: %s", err)
		}

		if len(output) != 0 {
			log.Printf("[binary] %v", output)
		}
	}
}

func callBin(arg string) error {
	cmdCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	myBinary := "./binary/bin"
	cmd := exec.CommandContext(cmdCtx, myBinary, arg)
	log.Printf("cmd: %v", cmd.String())

	pin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to pipe stdin: %s", err)
	}

	pout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to pipe stdout: %s", err)
	}

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("unable to start: %s", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		readStdout(pout)
	}()

	data := randBytes(105)
	chunksNumber := len(data) / chunkSize
	if len(data)%chunkSize != 0 {
		chunksNumber += 1
	}
	log.Printf("data length: %d, chunks number: %d", len(data), chunksNumber)

	writeErr := writeWithChunks(bytes.NewBuffer(data), pin)
	if writeErr != nil {
		pin.Close()
		return fmt.Errorf("failed to write to stdin: %s", writeErr)
	}
	pin.Close()

	resultErr := cmd.Wait()
	if resultErr != nil {
		return fmt.Errorf("cmd wait failed: %s", resultErr)
	}

	wg.Wait()

	return nil
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Println("wrong args")
		return
	}

	if args[1] != "timeout" && args[1] != "exit" && args[1] != "die" && args[1] != "ok" {
		log.Println("unknown args")
		return
	}

	err := callBin(args[1])
	if err != nil {
		log.Fatal(err)
	}

	log.Println("done")
}
