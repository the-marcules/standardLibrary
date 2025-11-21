'use server'

import fs from 'fs'

export const getDictionary = async (locale: Locale): Promise<Dict> => {
  const importDirectory = `./src/dictionaries/${locale}/`

  const files = fs.readdirSync(importDirectory).filter((file) => file.endsWith('.json'))

  const dict: Dict = {}

  await Promise.all(
    files.map((file) =>
      import(`../../dictionaries/${locale}/${file}`).then((module) =>
        Object.assign(dict, { [file.replace('.json', '')]: module.default })
      )
    )
  )

  return dict
}
