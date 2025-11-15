const removeShorts = () => {
  // Homepage shelves
  document
    .querySelectorAll('ytd-reel-shelf-renderer')
    .forEach((el) => el.remove());

  // Sidebar "Shorts" link
  document.querySelectorAll('a[title="Shorts"]').forEach((el) => el.remove());

  // Shorts videos across YouTube
  [
    'ytd-reel-shelf-renderer',
    'ytd-rich-item-renderer',
    'ytd-grid-video-renderer',
    'ytd-video-renderer',
  ].forEach((type) => {
    document.querySelectorAll(`${type} a[href^="/shorts"]`).forEach((el) => {
      el.closest(type)?.remove();
    });
  });
};

// Run immediately
removeShorts();

// Keep watching
const observer = new MutationObserver(removeShorts);
observer.observe(document.body, { childList: true, subtree: true });
