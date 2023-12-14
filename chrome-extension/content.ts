import { DictionaryClient } from "./src/clients/dictionary";
import { createApp, ref } from "vue";
import Notification from "./src/components/Notification.vue";
import { DictionaryPage } from "./src/types/dictionary";
import s from "./src/index.css?inline";

const notificationId = "nnnotification";
if (!document.querySelector(`#${notificationId}`)) {
  const div = document.createElement("div");
  div.id = notificationId;
  document.body.insertBefore(div, document.body.firstChild);
  const shadowRoot = div.attachShadow({ mode: "open" });
  shadowRoot.innerHTML = `<style>${s}</style><div id='nnnotificationShadow'></div>`;
}
const page = ref<DictionaryPage>({ symbol: "init" } as DictionaryPage);
const open = ref<boolean>(false);
const notification = createApp(Notification, {
  page: page,
  open: open,
});

let element = document
  .querySelector(`#${notificationId}`)!
  .shadowRoot!.querySelector("#nnnotificationShadow") as Element;

notification.mount(element);

function handleCopy() {
  const selection = window.getSelection()?.toString() || "";
  open.value = false;
  if (selection === "") {
    open.value = false;
    return;
  }
  navigator.clipboard.writeText(selection);
  DictionaryClient.getInstance()
    .lookup(selection)
    .then((newPage) => {
      if (newPage) {
        if (newPage.definitions && newPage.definitions.length > 0) {
          const audio = document.createElement("audio");
          audio.src = newPage.definitions[0].pronunciationLink;

          audio.play();
        }
        open.value = true;
        page.value = newPage;
      } else {
        open.value = false;
      }
    });
}

chrome.runtime.onMessage.addListener(
  (message, _1: chrome.runtime.MessageSender, _2: any) => {
    const action = message.action;
    switch (action) {
      case "copy":
        handleCopy();
        break;
    }
  }
);
