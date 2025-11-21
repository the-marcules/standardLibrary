import '@/styles/globals.css'
import { TranslationProvider } from '@/core/hooks/useTranslation'
import AppHeader from '@/components/layout/appHeader/AppHeader'
import AppFooter from '@/components/layout/appFooter/AppFooter'

export default async function RootLayout({
  children,
  params,
}: Readonly<{
  children: React.ReactNode
  params: { lang: string }
}>) {
  const { lang: locale } = await params

  return (
    <html lang={locale}>
      <TranslationProvider>
        <body>
          <AppHeader />
          <main>{children}</main>
          <AppFooter />
        </body>
      </TranslationProvider>
    </html>
  )
}
