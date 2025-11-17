document.getElementById('getTranscript')?.addEventListener('click', () => {
  browser.tabs.query({ active: true, currentWindow: true }).then((tabs) => {
    if (!tabs[0].id) return;
    browser.tabs
      .sendMessage(tabs[0].id, { action: 'GET_TRANSCRIPT' })
      .then((response) => {
        const transcriptElem = document.getElementById('transcript');
        if (transcriptElem) {
          transcriptElem.textContent = response || 'No transcript found.';
        }
      });
  });
});
