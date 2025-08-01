import styles from './header.module.css';

export default function Layout() {
  return (
    <header className={styles.header}>
      <div>Grid 'n' 💩</div>
      <div className={styles.navElement}>some | menu | items</div>
    </header>
  );
}
