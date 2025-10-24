import { useState, useEffect } from "react";
import { getDictionary } from "../translation/dictionaries";
import { useParams } from "next/navigation";
import { translate } from "../translation/translation";

export default function useTranslation() {
  const { lang: localeFromParams } = useParams<{ lang: Locale }>();
  const [dict, setDict] = useState<Dict | null>(null);
  const [locale, setLocale] = useState<Locale | undefined>(localeFromParams);

  useEffect(() => {
    const fetchDictionary = async () => {
      const newDict = await getDictionary(localeFromParams);
      setDict(newDict);
    };

    const setStateLocale = async () => {
      setLocale(localeFromParams);
    };

    if (localeFromParams !== locale || dict === null) {
      setStateLocale();
      fetchDictionary();
    }
  }, [localeFromParams, locale, dict]);

  return { t: (key: string) => translate(dict, key), dict, locale };
}
