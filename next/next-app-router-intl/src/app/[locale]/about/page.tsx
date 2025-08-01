import {useTranslations} from 'next-intl';
import {Link} from '@/i18n/navigation';
import { NextPage } from 'next';
 
const AboutPage: NextPage = () => {
  const t = useTranslations('AboutPage');
  return <div>
      <h1>{t('title')}</h1>
      <Link href="/">{t('home')}</Link>
    </div>
  ;
};

export default AboutPage;