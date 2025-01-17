package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateTgz(inputFileName string) (string, error) {
	// Open the input file
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return "", err
	}
	defer inputFile.Close()

	dir, _ := os.Getwd()

	// Extract the file name without extension
	base := filepath.Base(inputFileName)
	outputFileName := strings.TrimSuffix(base, filepath.Ext(base)) + ".tgz"
	outputPath := fmt.Sprintf("%s/%s/%s", dir, TMP_RUN_FILE, outputFileName)

	// Create a new .tgz file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	// Create a tar writer
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// Create a new header for the file in the tar
	fileInfo, err := inputFile.Stat()
	if err != nil {
		return "", err
	}

	header, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		return "", err
	}

	// Set the file name in the header
	header.Name = base

	// Write the header to the tar
	if err := tarWriter.WriteHeader(header); err != nil {
		return "", err
	}

	// Copy the file content to the tar
	_, err = io.Copy(tarWriter, inputFile)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}

func CreateTestFile(filename string, code string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("%s", "failed to find working directory: "+err.Error())
	}

	path := fmt.Sprintf("%s/%s/%s", dir, TMP_TESTS_FILE, filename)

	file, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("%s", "failed to create test file: "+err.Error())
	}

	os.WriteFile(path, []byte(code), 0644)

	defer file.Close()

	return path, nil
}

func IsErrorOutput(output string) bool {
	return strings.HasPrefix(output, ERROR_PREFIX)
}

func GetOutPath(filePath string) string {
	chunks := strings.Split(filePath, "/")
	outPath := strings.Join(chunks[:len(chunks)-1], "/") + "/app"

	return outPath
}
