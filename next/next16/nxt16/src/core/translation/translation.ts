import { getDictionary } from "./dictionaries";

export default async function getTranslations(
  params: Promise<{ lang: Locale }>
): Promise<Dict> {
  const { lang: locale } = await params;
  const dict: Dict = await getDictionary(locale);
  return (key: string) => translate(dict, key);
}

export function translate(dict: Dict, key: string): string {
  if (!dict) {
    return key;
  }
  return key.split(".").reduce((obj, key) => obj[key], dict) || key;
}
