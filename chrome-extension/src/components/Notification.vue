<template>
  <div
    v-if="open"
    class="fixed text-blue-950 float-right z-[9999] right-0 bottom-0 bg-white w-64 rounded h-screen overflow-y-auto text-base font-sans"
  >
    <h1 class="font-semibold p-4 bg-blue-950 text-amber-300">
      {{ page.symbol }}
    </h1>
    <div
      v-for="(definition, idx) in page.definitions"
      :key="idx"
      class="p-2 border-b"
    >
      <div class="">
        <div class="flex flex-col px-2">
          <div
            class="flex flex-row items-center text-gray-500 gap-2 cursor-pointer"
          >
            <div
              class="text-sm rounded border border-blue-300 bg-blue-100 text-blue-950 px-2"
              @click="
                () => {
                  handleAudioPlay(`audio-${page.symbol}-${idx}`);
                }
              "
            >
              {{ definition.pos }}
            </div>
            <audio :id="`audio-${page.symbol}-${idx}`">
              <source type="audio/mpeg" :src="definition.pronunciationLink" />
            </audio>
          </div>
          <div v-if="definition.meaning">
            <div>{{ definition.meaning.text }}</div>
            <div class="text-blue-500">
              {{ definition.meaning.translation }}
            </div>
            <div class="pl-2">
              <ul class="list-disc">
                <li
                  class="text-blue-500"
                  v-for="(example, eidx) of definition.meaning.examples"
                  :key="eidx"
                >
                  {{ example.text }}
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { defineProps, ref } from "vue";
import { DictionaryPage } from "../types/dictionary";
const props = defineProps<{ page: DictionaryPage; open: boolean }>();
const page = ref(props.page);
const open = ref(props.open);

const handleAudioPlay = (audioId: string) => {
  const audio = document
    .getElementById("nnnotification")
    ?.shadowRoot?.getElementById(audioId) as any;
  audio!.play();
};
</script>
