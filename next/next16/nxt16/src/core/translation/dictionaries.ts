export const getDictionary = async (locale: Locale): Promise<Dict> =>
  import(`../../dictionaries/${locale}.json`).then((module) => module.default);
