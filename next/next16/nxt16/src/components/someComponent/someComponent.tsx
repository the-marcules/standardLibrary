"use client";

import useTranslation from "@/core/hooks/useTranslation";

export default function SomeComponent() {
  const { t } = useTranslation();
  return (
    <div>
      <ul>
        <li>{t("common.title")}</li>
        <li>{t("common.description")}</li>
        <li>{t("common.welcome")}</li>
      </ul>
      {t("common.loading")}
    </div>
  );
}
