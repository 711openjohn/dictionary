export type Example = {
  text: string;
  translation: string;
};

export type Meaning = {
  text: string;
  translation: string;
  examples: Example[];
};

export type Definition = {
  pos: string;
  pronunciationLink: string;
  meaning: Meaning;
};

export type DictionaryPage = {
  symbol: string;
  definitions: Definition[];
};
