'use client'

import useTranslation from '@/core/hooks/useTranslation'
import Link from 'next/link'
import React from 'react'
import styles from './AppHeader.module.css'

const AppHeader: React.FC = () => {
  const { locale } = useTranslation()

  return (
    <div className={styles.header}>
      <nav>
        <Link href={`/${locale}`}>Home</Link>
        <Link href={`/${locale}/about`} style={{ marginLeft: '1rem' }}>
          About
        </Link>
        <Link href={`/${locale}/upload`} style={{ marginLeft: '1rem' }}>
          Upload
        </Link>
      </nav>
    </div>
  )
}

export default AppHeader
