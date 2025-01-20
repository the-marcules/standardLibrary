import MagicComponent from '@/components/magicComponent';
import styles from '../styles/Layout.module.css';

export default function Layout() {
  return (
    <div className={styles.basicLayout}>
      <header className={styles.header}>
        <p>header</p>
        <div className={styles.nestedGridElement}>Hello</div>
      </header>
      <div className={styles.sideNav}>side</div>
      <main className={styles.content}>
        <div>World</div>
      </main>
      <MagicComponent></MagicComponent>
    </div>
  );
}
