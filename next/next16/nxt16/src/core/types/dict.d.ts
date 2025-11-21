type Locale = 'en' | 'de'
type Dict = Record<string, string | Dict>
type TranslationConfig = {
  supportedLocales: Locale[]
  defaultLocale: Locale
}
type TranslationParameter = Record<string, string>
