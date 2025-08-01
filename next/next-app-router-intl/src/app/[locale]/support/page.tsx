import { NextPage } from "next";
import  { useTranslations } from "next-intl";

const SupportPage: NextPage = () => {
  const t = useTranslations("SupportPage");
  return (
    <div>
      <h1>{t('title')}</h1>
      <p>{t('description')}</p>
    </div>
  );
}

export default SupportPage;