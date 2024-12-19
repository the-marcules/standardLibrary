import { defineConfig } from 'cypress';
import { addCucumberPreprocessorPlugin } from "@badeball/cypress-cucumber-preprocessor";
import createBundler from '@bahmutov/cypress-esbuild-preprocessor'
import { createEsbuildPlugin } from '@badeball/cypress-cucumber-preprocessor/esbuild'


async function setupNodeEvents(on, config) {
  await addCucumberPreprocessorPlugin(on, config)

  on(
    'file:preprocessor',
    createBundler({
      plugins: [createEsbuildPlugin(config)]
    }))

    return config
}

export default defineConfig({
  e2e: {
    setupNodeEvents,
    specPattern: "cypress/e2e/features/*.feature",
    watchForFileChanges:false,
    chromeWebSecurity: false,
    experimentalModifyObstructiveThirdPartyCode: true,
    defaultCommandTimeout: 10000
 },
})