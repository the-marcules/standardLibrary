import Header from '@/components/header/header'
import './styles/globals.css'

export const metadata = {
  title: 'My Next.js App',
  description: 'A Next.js app using the App Router',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>
        <Header />
        {children}
      </body>
    </html>
  )
}
