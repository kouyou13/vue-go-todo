import pluginVue from "eslint-plugin-vue"
import { defineConfigWithVueTs, vueTsConfigs } from "@vue/eslint-config-typescript"
import js from "@eslint/js"
import globals from "globals"
import pluginPrettier from "eslint-plugin-prettier"
import eslintConfigPrettier from "eslint-config-prettier/flat"
import pluginImport from "eslint-plugin-import"
import pluginUnusedImports from "eslint-plugin-unused-imports"

export default defineConfigWithVueTs({
  files: ["**/*.{ts,vue}"],

  languageOptions: {
    globals: {
      ...globals.browser,
    },
  },

  extends: [
    js.configs.recommended,
    pluginVue.configs["flat/recommended"],
    vueTsConfigs.recommended,
    eslintConfigPrettier,
  ],

  plugins: {
    prettier: pluginPrettier,
    import: pluginImport,
    "unused-imports": pluginUnusedImports,
  },

  rules: {
    "import/order": [
      "error",
      {
        groups: ["builtin", "external", "internal", "parent", "sibling", "index"],
        "newlines-between": "always",
        alphabetize: {
          order: "asc",
          caseInsensitive: true,
        },
      },
    ],

    "prettier/prettier": [
      "warn",
      {
        semi: false,
        singleQuote: false,
        printWidth: 100,
      },
    ],

    "@typescript-eslint/consistent-type-imports": "error",
    "unused-imports/no-unused-imports": "warn",
  },
})
