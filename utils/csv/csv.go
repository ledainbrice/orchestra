package csv

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

//Read use csv.Read("assets/people.csv")
func Read(filename string) [][]string {
	var output [][]string
	f, err := os.Open(filename)
	checkError("error opening file= ", err)
	defer f.Close()
	// Create a new reader.

	r := csv.NewReader(bufio.NewReader(f))
	for {
		line, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		output = append(output, line)
		log.Println(line)
	}
	return output
}

//Write use csv.Write([][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}, "assets/people.csv")
func Write(data [][]string, filename string) {

	file, err := os.Create(filename)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func GenerateUpperCaseLetters(inputString string) chan string {
	// Create a channel where to send output
	outputChannel := make(chan string)
	// Launch an (anonymous) function in another thread, that does
	// the actual processing.
	go func() {
		// Loop over the letters in inputString
		for _, letter := range inputString {
			// Send an uppercased letter to the output channel
			outputChannel <- strings.ToUpper(string(letter))
		}
		// Close the output channel, so anything that loops over it
		// will know that it is finished.
		close(outputChannel)
	}()
	return outputChannel
}

func Download(c *gin.Context, fileName string, fileOut string) {
	targetPath := filepath.Join("tmp/", fileName)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileOut)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}

func checkError(message string, err error) {
	if err != nil {
		log.Println(message, err)
	}
}
