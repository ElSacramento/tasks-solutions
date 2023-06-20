package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("wrong args")
		os.Exit(1)
	}

	fmt.Printf("called args: %+v\n", args)

	i := 0
	var err error
	reader := bufio.NewReader(os.Stdin)
	for err == nil {
		_, err = reader.ReadByte()
		i += 1

		if i%10 == 0 {
			fmt.Printf("read %d bytes\n", i)
		}

		switch args[1] {
		case "timeout":
			time.Sleep(time.Millisecond * 100)
		case "exit":
			if i == 35 {
				fmt.Println("need to exit")
				os.Exit(1)
			}
		case "die":
			if i == 35 {
				os.Stdin.Close()
				return
			}
		}
	}

	if err != nil && err != io.EOF {
		fmt.Printf("failed to read byte: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("done")
}
