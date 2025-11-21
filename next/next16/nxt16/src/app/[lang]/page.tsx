import styles from "./page.module.css";
import getTranslations from "@/core/translation/translation";
import SomeComponent from "@/components/someComponent/someComponent";

export default async function Home({
  params,
}: {
  params: Promise<{ lang: Locale }>;
}) {
  const t = await getTranslations(params);

  return (
    <div className={styles.page}>
      <main className={styles.main}>
        <div className={styles.intro}>
          <h1>{t("common.welcome")}</h1>
          <SomeComponent></SomeComponent>
        </div>
      </main>
    </div>
  );
}
