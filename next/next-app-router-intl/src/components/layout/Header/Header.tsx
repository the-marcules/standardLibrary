import { useTranslations } from "next-intl";
import Link from "next/link";
import styles from "./Header.module.css";
import LanguageSwitcher from "../LanguageSwitcher/LanguageSwitcher";
const Header = () => {
  const t = useTranslations("navigation");
  return (
    <header className={styles.header}>
      <div className={styles.logo}>
        <svg
          width="48"
          height="36"
          viewBox="0 0 48 36"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
          
        >
          {/* Smooth blend of Alpha (α) and Omega (Ω) */}
          <path
            d="M12 28
               Q16 10 24 10
               Q32 10 36 28"
            stroke="#ff6600"
            strokeWidth="3"
            fill="none"
            strokeLinecap="round"
          />
          <path
            d="M16 28
               Q18 18 24 18
               Q30 18 32 28"
            stroke="#222"
            strokeWidth="2"
            fill="none"
            strokeLinecap="round"
          />
          <ellipse
            cx="24"
            cy="22"
            rx="3"
            ry="2"
            fill="#ff6600"
            opacity="0.7"
          />
        </svg>
        <span>
          Standard Library
        </span>
      </div>
      <nav>
        <Link href="/" className="nav-link">
          {t("home")}
        </Link>
        <Link href="/about" className="nav-link">
          {t("about")}
        </Link>
        <Link href="/support" className="nav-link">
          {t("support")}
        </Link>
        <Link href="/documentation" className="nav-link">
          {t("documentation")}
        </Link>
      </nav>
      <LanguageSwitcher />
    </header>
  );
};

export default Header;
