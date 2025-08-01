import { ChangeEvent, TextareaHTMLAttributes, useState } from 'react';
import { useNotification, NotificationTTL, SetTTL } from '../../notification/notificationProvider';
import Button from '../../Button/Button';
import { Verify as CryptokitVerify } from '../../../../wailsjs/go/cryptokit/Client';
import Result from '../../result/result';

export function Verify() {
  const [result, setResult] = useState<Response>();
  const [format, setFormat] = useState<string>('application/mose');
  const [signature, setSignature] = useState<string>('');

  const notificationService = useNotification();

  function handleVerify() {
    if (signature === '') {
      notificationService?.addNotification(
        'warning',
        'Warning',
        'Payload cannot be empty. Fill out the payload field and try again, you must.',
        SetTTL(NotificationTTL.warning)
      );
      return;
    }
    CryptokitVerify(format, signature).then((response) => {
      const verificationResponse: Response = JSON.parse(response);
      setResult(verificationResponse);
      if (verificationResponse.verification.result == 'good') {
        notificationService?.addNotification(
          'success',
          'Valid',
          'Signature is valid and verified.',
          SetTTL(NotificationTTL.success)
        );
      } else {
        notificationService?.addNotification('error', 'Invalid', 'Signature is invalid.');
      }
    });
  }

  const reset = () => {
    const textarea = document.querySelector<HTMLTextAreaElement>('#signInput');
    if (textarea) {
      textarea.value = '';
    }
    setFormat('');
    setSignature('');
  };

  const updateFormat = (e: ChangeEvent<HTMLSelectElement>) => setFormat(e.target.value);
  const updateSignature = (e: ChangeEvent<HTMLTextAreaElement>) => setSignature(e.target.value);

  return (
    <>
      <div id="input" className={`inputBox`}>
        <h3>Format</h3>
        <label title="Payload">
          {/* <input onChange={updateFormat} id="formatInput" placeholder="Sign format" /> */}
          <select onChange={updateFormat} id="formatInput">
            <option value="application/mose">application/mose</option>
          </select>
        </label>
        <h3>Signature</h3>
        <label title="Payload">
          <textarea onChange={updateSignature} id="signInput" placeholder="Your signature" />
        </label>
        <div className={'buttonContainer'}>
          <Button onClick={() => reset()} variant="secondary">
            Cancel
          </Button>
          <Button onClick={handleVerify}>Verify &raquo;</Button>
        </div>
      </div>

      <Result response={result} title="Response from CryptoKit" />
    </>
  );
}
