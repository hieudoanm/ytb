browser.runtime.onMessage.addListener((request) => {
  if (request.action === 'GET_TRANSCRIPT') {
    return getTranscript();
  }
});

async function getTranscript(): Promise<string | null> {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const playerData = (window as any).ytInitialPlayerResponse;
  if (!playerData.captions) return null;

  const captionTracks =
    playerData.captions.playerCaptionsTracklistRenderer.captionTracks;
  if (!captionTracks || captionTracks.length === 0) return null;

  const captionUrl = captionTracks[0].baseUrl;
  const res = await fetch(captionUrl);
  const xml = await res.text();

  const parser = new DOMParser();
  const doc = parser.parseFromString(xml, 'text/xml');

  return Array.from(doc.getElementsByTagName('text'))
    .map((node) => node.textContent?.replace(/&#39;/g, "'") ?? '')
    .join(' ');
}
