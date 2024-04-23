package main

import (
	"WaveToMP3Converter/internal/ffmpeg"
	"flag"
	"fmt"
)

func main() {
	// コマンドラインオプションの定義
	name := flag.String("name", "World", "a name to say hello to")
	flag.Parse() // オプションの解析

	// オプションの使用
	fmt.Printf("Hello, %s!\n", *name)

	inputPath := "path/to/input.wav"
	outputPath := "path/to/output.mp3"

	if err := ffmpeg.ConvertToMP3(inputPath, outputPath); err != nil {
		fmt.Printf("Error converting file: %v\n", err)
		return
	}

	fmt.Println("File converted successfully!")
}
