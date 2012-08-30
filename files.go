package fun

import (
	"io/ioutil"
	"log"
	"os"
)

func OpenAndRead(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func MaybeFatalAt(where string, err error) {
    if err != nil {
		log.Fatalf("Error near %s: %v\n", where, err)
	}
	return
}
