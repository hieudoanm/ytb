# YouTube Transcript Summariser

A CLI tool to fetch and format YouTube transcripts (captions) into text, JSON, or SRT formats.
This version is packaged as a **standalone binary** using PyInstaller — no Python installation required to run the built binary.

---

## Features

- Interactive prompts for **YouTube URL** and **output format**
- Output formats: `text` (default), `json`, `srt`
- Option to save transcript to a file
- Loading spinner while fetching transcript

---

## Build Binary

### 1. Install Dependencies

```bash
pip install -r requirements.txt
pip install pyinstaller
```

### 2. Build with PyInstaller

```bash
pyinstaller --onefile yt_transcript.pybashbash
```

The binary will be created in the dist/ folder:

```txt
dist/yt_transcript   # macOS / Linux
dist/yt_transcript.exe  # Windows
```
