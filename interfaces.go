package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/waditya/go-io-operations/shape"
)

func main() {
	s := shape.Rectangle{Length: 10.0, Width: 6.0}
	fmt.Println("Area of rectangle is : ", s.Area())
	fmt.Println("Perimeter of rectangle is :", s.Perimeter())

	r := strings.NewReader("Learning Go is fun")

	// Create a slice of size 4 and of type bytes using make function
	buf := make([]byte, 4)

	// Read the values in buf using the Read method of NewReader
	numberOfBytesRead, err := r.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Bytes Read from buf (In ASCII): ", buf[:numberOfBytesRead])
	fmt.Println("Bytes Read from buf : ", string(buf[:numberOfBytesRead]))

	for {
		// Read the values in buf using the Read method of NewReader
		numberOfBytesRead, err := r.Read(buf)
		fmt.Println("Bytes Read from buf (In ASCII): ", buf[:numberOfBytesRead])
		fmt.Println("Bytes Read from buf : ", string(buf[:numberOfBytesRead]))

		if err != nil {
			fmt.Println("Breaking out due to error : \n", err)
			break
		}
	}

	// Exploring the Copy method of io package to print the data provided by NewReader on Standard output
	stringToBeCopied := strings.NewReader("My sample string to be read and passed to standard output")

	// Copy copies from src to dst until either EOF is reached on src or an error occurs.
	// It returns the number of bytes copied and the first error encountered while copying, if any.
	// os.Stdout will write to the Standard Output. Standard Outout File implements the Write method of Writer interface in IO package
	// Hence, we are able to copy and WRITE to os.StdOut
	if _, err := io.Copy(os.Stdout, stringToBeCopied); err != nil {
		log.Fatal(err)
	}
}
