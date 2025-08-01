import { NextIntlClientProvider } from "next-intl";
import React from "react";

export default function Layout({ children }: { children: React.ReactNode }) {

  // const fixedHeaderScrollBeahavior = (e) => {
  //   e.preventDefault();
  //   const targetId = this.getAttribute("href");
  //   const targetElement = document.querySelector(targetId);
  //   if (targetElement) {
  //     targetElement.scrollIntoView({ behavior: "smooth" });
  //   }
  // };

  // useEffect(() => {
  //   document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
  //     anchor.addEventListener("click", fixedHeaderScrollBeahavior);
  //   });

  //   return () => {
  //     document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
  //       anchor.removeEventListener("click", fixedHeaderScrollBeahavior);
  //     });
  //   }
  // }, []);

  return (
    <html lang="de">
      <body>
        <NextIntlClientProvider locale="de">{children}</NextIntlClientProvider>
      </body>
    </html>
  );
}
