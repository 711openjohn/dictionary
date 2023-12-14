chrome.commands.onCommand.addListener(
  (command: string, _tab: chrome.tabs.Tab) => {
    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
      chrome.tabs.sendMessage(tabs[0].id!, { action: command });
    });
  }
);
