from urllib.parse import urlparse, parse_qs
from pyfiglet import figlet_format


def extract_video_id(url):
    parsed = urlparse(url)
    if parsed.hostname in ("www.youtube.com", "youtube.com"):
        return parse_qs(parsed.query).get("v", [None])[0]
    if parsed.hostname == "youtu.be":
        return parsed.path.lstrip("/")
    return url


def main():
    # Show banner using figlet
    try:
        print(figlet_format("YouTube Transcript Summariser"))
    except ImportError:
        print("YouTube Transcript Summariser")  # fallback if pyfiglet not installed

    # Prompt FIRST (instant)
    video_url = input("Enter YouTube URL: ").strip()
    if not video_url:
        print("No URL provided. Exiting.")
        return

    format_input = (
        input("Choose format (text/json/srt) [default: text]: ").strip().lower()
    )
    if format_input not in ["text", "json", "srt", ""]:
        print("Invalid format. Defaulting to text.")
        format_input = "text"
    elif format_input == "":
        format_input = "text"

    # Import heavy libs AFTER prompt
    from youtube_transcript_api import YouTubeTranscriptApi
    from yaspin import yaspin
    from rich.console import Console
    from rich.markdown import Markdown
    import os

    console = Console()

    video_id = extract_video_id(video_url)

    with yaspin(text="Fetching transcript...", color="cyan") as spinner:
        try:
            transcript = YouTubeTranscriptApi.get_transcript(
                video_id, languages=("en", "en-US")
            )
            spinner.ok("✔")
        except Exception as e:
            spinner.fail("✖")
            print(f"Error fetching transcript: {e}")
            return

    content = format_transcript(transcript, format_input)

    # Ensure output folder exists
    output_dir = "output"
    os.makedirs(output_dir, exist_ok=True)

    # Save inside output folder
    extension = "txt" if format_input == "text" else format_input
    filename = os.path.join(output_dir, f"transcript_{video_id}.{extension}")
    with open(filename, "w", encoding="utf-8") as f:
        f.write(content)

    console.print(f"[green]Transcript saved to {os.path.abspath(filename)}[/green]")

    with yaspin(text="Generating summary...", color="yellow") as spinner:
        summary = summarise_transcript(content)
        spinner.ok("✔")

    console.print(Markdown(f"### Summary\n\n{summary}"))


def format_transcript(transcript, fmt):
    import json

    if fmt == "json":
        return json.dumps(transcript, indent=2, ensure_ascii=False)
    elif fmt == "srt":
        lines = []
        for i, entry in enumerate(transcript, start=1):
            start = entry["start"]
            end = start + entry.get("duration", 0)
            lines.append(
                f"{i}\n{format_time(start)} --> {format_time(end)}\n{entry['text']}\n"
            )
        return "\n".join(lines)
    else:
        return "\n".join(
            f"{entry['start']:.2f}s: {entry['text']}" for entry in transcript
        )


def format_time(seconds):
    hrs = int(seconds // 3600)
    mins = int((seconds % 3600) // 60)
    secs = int(seconds % 60)
    ms = int((seconds - int(seconds)) * 1000)
    return f"{hrs:02}:{mins:02}:{secs:02},{ms:03}"


def summarise_transcript(video_transcript: str) -> str:
    import requests

    try:
        system_prompt = "Summarize the following transcript in 3 concise bullet points using Markdown list format (- point):"
        user_prompt = f"{system_prompt}\n\n{video_transcript}"
        url = "https://micro24.vercel.app/api/genai"
        headers = {
            "Accept": "application/json",
            "Content-Type": "application/json",
        }
        payload = {
            "model": "deepseek/deepseek-r1:free",
            "messages": [{"role": "user", "text": user_prompt}],
        }
        response = requests.post(url, headers=headers, json=payload, timeout=60)
        response.raise_for_status()
        data = response.json()
        return data.get("output", "No summary.")
    except Exception as e:
        print(f"Unexpected error: {str(e)}")
        return "No summary."


if __name__ == "__main__":
    main()
