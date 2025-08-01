import {
  Before,
  After,
  BeforeAll,
  AfterAll,
  setDefaultTimeout,
} from '@cucumber/cucumber'
import { chromium } from 'playwright'
import type { Browser, Page } from 'playwright'
import { spawn } from 'child_process'

setDefaultTimeout(60 * 1000)

let serverProcess: any
let browser: Browser
let page: Page

BeforeAll(async function () {
  console.log('🚀 Starting the development server')
  serverProcess = spawn('npm', ['run', 'dev'], {
    shell: true,
  })

  await developmentServerIsReady()
})

AfterAll(async function () {
  console.log('🛑 Stopping the development server...')
  if (serverProcess) {
    serverProcess.kill()
    console.log('Development server stopped.')
  }
})

Before(async function () {
  browser = await chromium.launch({ headless: false }) // Launch the browser
  const context = await browser.newContext() // Create a new browser context
  page = await context.newPage() // Open a new page
  this.browser = browser // Attach browser to the Cucumber world
  this.page = page // Attach page to the Cucumber world
})

After(async function () {
  if (browser) {
    await browser.close() // Close the browser after each scenario
  }
})

async function developmentServerIsReady(): Promise<void> {
  const serverUrl = 'http://localhost:3000'
  const maxRetries = 10
  const retryInterval = 1500

  let retries = 0
  while (retries < maxRetries) {
    try {
      await fetch(serverUrl)
      console.log('✅ Development server is ready')
      break
    } catch (error) {
      retries++
      console.log(`⏱️  Waiting for the server... (${retries}/${maxRetries})`)
      await new Promise((resolve) => setTimeout(resolve, retryInterval))
    }
  }

  if (retries === maxRetries) {
    console.error('Failed to start the development server.')
    process.exit(1) // Exit the process if the server doesn't start
  }
}

export { browser, page }
