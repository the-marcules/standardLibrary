"use client";

import { useRouter } from "next/navigation";
import styles from "./LanguageSwitcher.module.css";
import { useTranslations, useLocale } from "next-intl";
import { useEffect, useState } from "react";

const LanguageSwitcher = () => {
  const router = useRouter();
  const t = useTranslations("common");
  const currentLocale = useLocale();
  const [showDropdown, setShowDropdown] = useState(false);

  useEffect(() => {
    document.addEventListener("click", (e) => {
      const target = e.target as HTMLElement;
      if (!target.closest(`.${styles.container}`)) {
        setShowDropdown(false);
      }
    });
    return () => {
      document.removeEventListener("click", () => {});
    };
  }, []);

  const changeLanguage = (locale: string) => {
    setShowDropdown(false);
    router.push(window.location.pathname.replace(/\/[a-z]{2}/, `/${locale}`));
  };

  const toggleDropdown = () => {
    setShowDropdown(!showDropdown);
  };

  return (
    <div className={styles.container}>
      <button onClick={() => toggleDropdown()}>
        {t(`lng.${currentLocale}`).split(" ")[0]}
      </button>
      {showDropdown && (
        <div className={styles.languageButtonsMenu}>
          <button onClick={() => changeLanguage("en")}>{t("lng.en")}</button>
          <button onClick={() => changeLanguage("de")}>{t("lng.de")}</button>
        </div>
      )}
    </div>
  );
};

export default LanguageSwitcher;
