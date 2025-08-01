import styles from './sign.module.css';
import { useState } from 'react';
import { Sign as CryptokitSign } from '../../../../wailsjs/go/cryptokit/Client';
import Result from '../../result/result';
import { useNotification, NotificationTTL, SetTTL } from '../../notification/notificationProvider';
import Button from '../../Button/Button';

export function Sign() {
  const [result, setResult] = useState<Response>();
  const [payload, setPayload] = useState<string>('');
  const updateName = (e: any) => setPayload(e.target.value);
  const notificationService = useNotification();

  function handleSign() {
    if (payload === '') {
      notificationService?.addNotification(
        'warning',
        'Warning',
        'Payload cannot be empty. Fill out the payload field and try again, you must.',
        SetTTL(NotificationTTL.warning)
      );
      return;
    }

    CryptokitSign(payload).then((response) => {
      const responseObject = JSON.parse(response) as Response;
      setResult(responseObject);
      if (responseObject.error.isError) {
        notificationService?.addNotification(
          'error',
          'Error',
          responseObject.error.message || responseObject.description || ''
        );
      } else {
        notificationService?.addNotification(
          'success',
          'Success',
          'Got a result',
          Date.now() + NotificationTTL.success
        );
      }
    });
  }

  const reset = () => {
    const textarea = document.querySelector<HTMLTextAreaElement>('#signInput');
    if (textarea) {
      textarea.value = '';
    }
    setPayload('');
  };

  return (
    <>
      <div id="input" className={`inputBox`}>
        <h3>Payload</h3>
        <label title="Payload">
          <textarea className={styles.input} onChange={updateName} id="signInput" placeholder="your payload" />
        </label>
        <div className={'buttonContainer'}>
          <Button onClick={() => reset()} variant="secondary">
            Cancel
          </Button>
          <Button onClick={handleSign}>Sign &raquo;</Button>
        </div>
      </div>

      <Result response={result} title="Response from CryptoKit" />
    </>
  );
}
