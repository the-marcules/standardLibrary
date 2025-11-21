import { NextResponse } from "next/server";
import { NextRequest } from "next/server";
import {
  getLocale,
  pathContainsValidLocale,
} from "./core/translation/translation";

export function proxy(request: NextRequest) {
  // Check if there is any supported locale in the pathname
  const { pathname } = request.nextUrl;
  const pathHasSupportedLocale = pathContainsValidLocale(pathname);

  if (pathHasSupportedLocale) return;

  // add locale to path and redirect
  const locale = getLocale();
  request.nextUrl.pathname = `/${locale}${pathname}`;
  return NextResponse.redirect(request.nextUrl);
}

export const config = {
  matcher: [
    // Skip following paths: _next, api , .svg
    "/((?!_next|api|.*\\.svg).*)",
  ],
};
