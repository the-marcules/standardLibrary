import { getDictionary } from './dictionaries'
import { config as translationConfig } from '@/config/translation.config'
import { ReactElement } from 'react'

export default async function getTranslations(params: Promise<{ lang: Locale }>): Promise<Dict> {
  const { lang: locale } = await params
  const dict: Dict = await getDictionary(locale)
  return (key: string) => translate(dict, key)
}

export function translate(dict: Dict, key: string): string {
  if (!dict) {
    return key
  }
  return (
    key.split('.').reduce((obj, key) => (typeof obj === 'object' && key in obj ? obj[key] : undefined), dict) || key
  )
}

export function translateToTsx(dict: Dict, key: string, param?: TranslationParameter): ReactElement {
  let result = translate(dict, key)

  if (result === key) {
    return <span>{key}</span>
  }

  if (param) {
    for (const [paramKey, paramValue] of Object.entries(param)) {
      result = result.replaceAll('{{' + paramKey + '}}', paramValue)
    }
  }

  const sanitized = result.replace(/<(?!br\/?|strong\b)[^>]+>/gi, '')
  return <span dangerouslySetInnerHTML={{ __html: sanitized }} />
}

export function isLocaleSupported(locale: string): boolean {
  return translationConfig.supportedLocales.includes(locale as Locale)
}

export function getUsersPreferredLocale(): Locale {
  if (typeof navigator === 'undefined') {
    return getDefaultLocale()
  }

  const navigatorLocale = navigator.language.split('-')[0]

  if (isLocaleSupported(navigatorLocale)) {
    return navigatorLocale as Locale
  }

  return getDefaultLocale()
}

export function getDefaultLocale(): Locale {
  return translationConfig.defaultLocale
}

export function getLocale(): Locale {
  return getUsersPreferredLocale() || getDefaultLocale()
}

export function pathContainsValidLocale(pathname: string): boolean {
  return translationConfig.supportedLocales.some(
    (locale) => pathname.startsWith(`/${locale}/`) || pathname === `/${locale}`
  )
}
