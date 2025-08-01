import Link from 'next/link'
import styles from './header.module.css'

export default function Header() {
  return (
    <header className={styles.header}>
      <div>APP</div>
      <nav>
        <Link href="/">Home</Link>
        <Link href="/second">Second</Link>
        <Link href="/about">About</Link>
      </nav>
    </header>
  )
}
