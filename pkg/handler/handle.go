package handler

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const bufferSize = 4 * 1024 * 1024

// MakeHandler creates a handler function that will serve files from the specified directory
func MakeHandler(baseDir string) func(http.ResponseWriter, *http.Request) {
	pathToIndex := filepath.Join(baseDir, "/index.html")
	return func(resp http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		if strings.HasSuffix(path, "/") {
			serveFile(pathToIndex, req, resp)
			return
		}

		path = filepath.Join(baseDir, path)
		fileStat, err := os.Stat(path)

		if os.IsNotExist(err) {
			serveFile(pathToIndex, req, resp)
			return
		}

		if err != nil {
			internalError(req, resp, err)
			return
		}

		if fileStat.IsDir() {
			serveFile(pathToIndex, req, resp)
			return
		}

		serveFile(path, req, resp)
	}
}

func internalError(req *http.Request, resp http.ResponseWriter, err error) {
	log.Printf("Error: %s\n", err)
	resp.WriteHeader(http.StatusInternalServerError)
	resp.Write([]byte(fmt.Sprintf("Sorry, something went wrong: '%s'\n", err)))
}

func serveFile(path string, req *http.Request, resp http.ResponseWriter) {
	log.Printf("%s %s", req.URL.Path, path)
	file, _ := os.Open(path) // File must exist, ignoring error here

	resp.Header().Add("content-type", mime.TypeByExtension(filepath.Ext(file.Name())))
	resp.WriteHeader(http.StatusOK)

	defer file.Close()

	buffer := make([]byte, bufferSize)
	for {
		bytesRead, readError := file.Read(buffer)

		if readError != nil && readError != io.EOF {
			log.Fatal("Error while reading file.", readError)
		}

		if bytesRead == 0 {
			break
		}

		resp.Write(buffer[:bytesRead])
	}
}
