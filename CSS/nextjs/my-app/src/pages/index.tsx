import styles from '../styles/Layout.module.css';
import Header from '@/components/Header/header';
import Footer from '@/components/Footer/footer';

export default function Layout() {
  return (
    <div className={styles.basicLayout}>
      <Header></Header>
      <main className={styles.content}>
        <div contentEditable>The content.</div>
      </main>
      <Footer />
    </div>
  );
}
