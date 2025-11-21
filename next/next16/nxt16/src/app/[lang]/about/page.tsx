import getTranslations from "@/core/translation/translation";

export default async function AboutPage({
  params,
}: {
  params: Promise<{ lang: Locale }>;
}) {
  const t = await getTranslations(params);

  return (
    <div>
      <h1>{t("about.title")}</h1>
      <p>{t("about.description")}</p>
    </div>
  );
}
