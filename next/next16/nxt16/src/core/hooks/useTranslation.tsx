'use client'
import { createContext, ReactElement, useContext, useEffect, useState } from 'react'
import { getDictionary } from '@/core/translation/dictionaries'
import { useParams } from 'next/navigation'
import { getLocale, isLocaleSupported, translate, translateToTsx } from '@/core/translation/translation'

export interface TranslationContextType {
  t: (key: string) => string
  ttsx: (key: string, param?: TranslationParameter) => ReactElement
  dict: Dict
  locale: Locale
}

const TranslationContext = createContext<TranslationContextType>({} as TranslationContextType)
const useTranslation = () => useContext(TranslationContext)

export const TranslationProvider = ({ children }: { children: React.ReactNode }) => {
  const { lang: localeFromParams } = useParams<{ lang: Locale }>()
  const [dict, setDict] = useState<Dict | null>(null)
  const [locale, setLocale] = useState<Locale>(localeFromParams)

  useEffect(() => {
    const fetchDictionary = async () => {
      const newDict = await getDictionary(localeFromParams)
      setDict(newDict)
    }

    const setStateLocale = async () => {
      setLocale(isLocaleSupported(localeFromParams) ? localeFromParams : getLocale())
    }

    if (localeFromParams !== locale || dict === null) {
      setStateLocale()
      fetchDictionary()
    }
  }, [localeFromParams, locale, dict])

  return (
    <TranslationContext.Provider
      value={{
        t: (key: string) => translate(dict, key),
        ttsx: (key: string, param?: TranslationParameter) => translateToTsx(dict, key, param),
        dict,
        locale,
      }}
    >
      {children}
    </TranslationContext.Provider>
  )
}

export default useTranslation
