import styles from './layout.module.css';

import Header from '../header/header';
import { Sign } from '../requests/sign/sign';
import { NotificationProvider } from '../notification/notificationProvider';
import { NotificationScreen } from '../notification/notification.screen';
import { BrowserRouter, Route, Routes } from 'react-router';
import { Verify } from '../requests/verify/verify';
import Encrypt from '../requests/encrypt/encrypt';
import Decrypt from '../requests/decrypt/decrypt';
import Tools from '../Tools/tools';
import Base64 from '../Tools/base64/base64';
import { GetKeysString } from '../../../wailsjs/go/cryptokit/Client';

export default function Layout() {
  return (
    <BrowserRouter>
      <NotificationProvider>
        <div id="App">
          <Header />
          <Tools />
          <div className={styles.content}>
            <Routes>
              <Route index path="/" element={<div>Choose your cryptographic operation form menu above.</div>} />
              <Route path="/sign" element={<Sign />} />
              <Route path="/verify" element={<Verify />} />
              <Route path="/encrypt" element={<Encrypt />} />
              <Route path="/decrypt" element={<Decrypt />} />
              <Route path="/tools/base64" element={<Base64 />} />
            </Routes>
          </div>
          <NotificationScreen />
        </div>
      </NotificationProvider>
    </BrowserRouter>
  );
}
