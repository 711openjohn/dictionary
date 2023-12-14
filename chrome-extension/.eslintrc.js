module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: [
    "@nuxtjs/eslint-config-typescript",
    "plugin:vue/vue3-recommended",
    "prettier",
  ],
  parserOptions: {
    ecmaVersion: 13,
    sourceType: "module",
  },
  plugins: ["prettier"],
  rules: {
    "prettier/prettier": "error",
  },
  overrides: [
    {
      files: [
        "**/pages/**/*.{js,ts,vue}",
        "**/layouts/**/*.{js,ts,vue}",
        "**/app.{js,ts,vue}",
        "**/error.{js,ts,vue}",
      ],
      rules: {
        "vue/no-multiple-template-root": 2,
        "vue/multi-word-component-names": "off",
        "vue/component-name-in-template-casing": ["error", "kebab-case"],
        "vue/prop-name-casing": ["error", "camelCase"],
        "vue/attribute-hyphenation": ["error", "always"],
        "vue/v-on-event-hyphenation": ["error", "always"],
      },
    },
  ],
};
