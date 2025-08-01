import { test, expect } from '@playwright/test'

test.describe('navigation', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('http://localhost:3000')
  })

  test('main page has loaded', async ({ page }) => {
    await expect(page).toHaveURL('http://localhost:3000/')
  })

  test('page has 3 links', async ({ page }) => {
    const header = page.locator(`Header`)
    const links = header.locator('a')
    await expect(links).toHaveCount(3)
  })

  test('click on "About" opens next page', async ({ page }) => {
    const aboutLink = page.getByText('About')
    expect(aboutLink).toBeVisible()
    await aboutLink.click()

    await page.waitForLoadState('networkidle')

    await expect(page.url()).toContain('about')
  })
})
