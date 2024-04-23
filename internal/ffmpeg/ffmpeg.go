package ffmpeg

import (
	"fmt"
	"os/exec"
)

// ConvertToMP3 は、指定された WAV ファイルを MP3 に変換します。
func ConvertToMP3(inputPath, outputPath string) error {
	fmt.Printf("Converting %s to %s...\n", inputPath, outputPath)

	cmd := exec.Command("ffmpeg", "-i", inputPath, "-codec:a", "libmp3lame", "-qscale:a", "2", outputPath)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg error: %w", err)
	}

	return nil
}
