#!/bin/bash

# ffmpegが既にインストールされているか確認
if ! command -v ffmpeg &> /dev/null
then
    echo "ffmpeg could not be found, installing..."
    # Homebrewを使ってffmpegをインストール
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        sudo apt update
        sudo apt install ffmpeg
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        brew install ffmpeg
    fi
else
    echo "ffmpeg is already installed."
fi
