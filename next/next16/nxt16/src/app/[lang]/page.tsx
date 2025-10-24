import styles from "./page.module.css";
import { getDictionary } from "@/core/translation/dictionaries";
import SomeComponent from "@/components/someComponent/someComponent";

export default async function Home({
  params,
}: {
  params: Promise<{ lang: Locale }>;
}) {
  const { lang: locale } = await params;
  const dict = await getDictionary(locale);

  return (
    <div className={styles.page}>
      <main className={styles.main}>
        <div className={styles.intro}>
          <h1>{dict.common.welcome}</h1>
          <SomeComponent></SomeComponent>
        </div>
      </main>
    </div>
  );
}
