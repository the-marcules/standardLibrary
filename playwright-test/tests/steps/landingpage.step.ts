import { Given, When, Then, Given as And } from '@cucumber/cucumber'
import { expect } from '@playwright/test'
import { page } from '../setup/cucumberTestSetup.ts'

Given('I am on the landing page', function () {
  page.goto('http://localhost:3000')
})

And('I see the Navigation bar', () => {
  const navbar = page.locator(`nav`)
  expect(navbar).toBeVisible()
})

When('I click on the {string} link', (linkText: string) => {
  expect(page.getByText(linkText)).toBeVisible()
  page.getByText(linkText).click()
})

Then('i should see the {string} page', async (pageName: string) => {
  await page.waitForLoadState('networkidle')

  await expect(page.url()).toContain(pageName.toLowerCase())
})
