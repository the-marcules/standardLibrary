import React, { useState } from 'react';

import { Decrypt as CryptoKitDecrypt } from '../../../../wailsjs/go/cryptokit/Client';
import Result from '../../result/result';
import Button from '../../Button/Button';

export default function Decrypt(): JSX.Element {
  const [payload, setPayload] = useState<string>('');
  const [result, setResult] = useState<Response>();

  const handleEncrypt = () => {
    CryptoKitDecrypt(payload).then((response) => {
      setResult(JSON.parse(response));
    });
  };

  const reset = () => {
    const textarea = document.querySelector<HTMLTextAreaElement>('#decryptInput');
    if (textarea) {
      textarea.value = '';
    }
    setPayload('');
    setResult(undefined);
  };
  return (
    <>
      <div className="inputBox">
        <h3>Encrypted payload</h3>
        <textarea
          id="decryptInput"
          title="payload"
          onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => {
            setPayload(e.target.value);
          }}
        />
        <div className={'buttonContainer'}>
          <Button onClick={() => reset()} variant="secondary">
            Cancel
          </Button>
          <Button onClick={handleEncrypt}>Decrypt &raquo;</Button>
        </div>
      </div>
      {result && <Result title="Decryption result" response={result as Response} />}
    </>
  );
}
