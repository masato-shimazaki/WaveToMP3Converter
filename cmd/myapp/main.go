package main

import (
	"WaveToMP3Converter/internal/ffmpeg"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// ConvertedFilesDir は、変換後のファイルを保存するディレクトリのパスです。
	ConvertedFilesDir = "./../../build/convert"
)

func main() {
	// コマンドラインオプションの定義
	greetingName := flag.String("name", "World", "a name to greet")
	flag.Parse() // オプションの解析

	// オプションの使用
	fmt.Printf("Hello, %s!\n", *greetingName)

	sourceDir := "./../../convert"
	fmt.Println("Source directory:" + sourceDir)

	if err := os.MkdirAll(ConvertedFilesDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create directory:", err)
		return
	}

	if err := convertFilesInDirectory(sourceDir); err != nil {
		fmt.Println("Failed to convert files:", err)
	}
}

func convertFilesInDirectory(directory string) error {
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %q: %w", path, err)
		}
		if !info.IsDir() {
			fmt.Println("Processing file:", path)
			outputFile := buildOutputFilePath(path)
			if err := ffmpeg.ConvertToMP3(path, outputFile); err != nil {
				return fmt.Errorf("error converting %q to %q: %w", path, outputFile, err)
			}
			fmt.Println("File converted successfully:", outputFile)
		}
		return nil
	})
}

func buildOutputFilePath(inputPath string) string {
	filenameWithoutExtension := trimFileExtension(filepath.Base(inputPath))
	return filepath.Join(ConvertedFilesDir, filenameWithoutExtension+".mp3")
}

func trimFileExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
