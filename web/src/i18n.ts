import { createI18n } from 'vue-i18n'
import en from './locals/en.json'
import zh from './locals/zh.json'

const messages = {
  en,
  zh,
}

const i18n = createI18n({
  leagcy: false,
  locale: 'zh',
  fallbackLocale: 'en',
  messages,
})

export default i18n
