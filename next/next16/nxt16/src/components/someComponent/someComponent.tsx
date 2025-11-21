"use client";

import useTranslation from "@/core/hooks/useTranslation";

export default function SomeComponent() {
  const { t } = useTranslation();
  return (
    <div>
      <ul>
        <li>{t("common.does.not.exist")}</li>
        <li>{t("does.not.Exist.either")}</li>
      </ul>
      {t("common.loading")}
    </div>
  );
}
