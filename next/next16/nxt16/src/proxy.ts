import { NextResponse } from "next/server";
import { NextRequest } from "next/server";

const locales: Locale[] = ["en", "de"];

// Get the preferred locale, similar to the above or using a library
function getLocale(): string {
  return "de";
}

export function proxy(request: NextRequest) {
  console.log("Middleware executed for:", request.url);
  // Check if there is any supported locale in the pathname
  const { pathname } = request.nextUrl;
  const pathnameHasLocale = locales.some(
    (locale) => pathname.startsWith(`/${locale}/`) || pathname === `/${locale}`
  );

  if (pathnameHasLocale) return;

  // Redirect if there is no locale
  const locale = getLocale();
  request.nextUrl.pathname = `/${locale}${pathname}`;
  // e.g. incoming request is /products
  // The new URL is now /en-US/products
  return NextResponse.redirect(request.nextUrl);
}

export const config = {
  matcher: [
    // Skip all internal paths (_next)
    "/((?!_next|api|.*\\.svg).*)",
    // Optional: only run on root (/) URL
    // '/'
  ],
};
