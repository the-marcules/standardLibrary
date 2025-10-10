import React, { useState, useEffect, ChangeEvent } from "react";

import {
  Decrypt as CryptoKitDecrypt,
  GetKeysString,
} from "../../../../wailsjs/go/cryptokit/Client";
import Result from "../../result/result";
import Button from "../../Button/Button";

export default function Decrypt(): JSX.Element {
  const [payload, setPayload] = useState<string>("");
  const [result, setResult] = useState<Response>();
  const [selectedKey, setSelectedKey] = useState<string>("");
  const [keys, setKeys] = useState<KeyData[]>([]);

  useEffect(() => {
    const fetchKeys = async () => {
      const tmpKeys = await GetKeysString();
      const keys = JSON.parse(tmpKeys) as KeyData[];
      setKeys(keys);
      setSelectedKey(keys[0].kid);
    };
    fetchKeys();
  }, []);

  const handleEncrypt = () => {
    CryptoKitDecrypt(payload, selectedKey).then((response) => {
      setResult(JSON.parse(response));
    });
  };

  const reset = () => {
    const textarea =
      document.querySelector<HTMLTextAreaElement>("#decryptInput");
    if (textarea) {
      textarea.value = "";
    }
    setPayload("");
    setResult(undefined);
  };
  return (
    <>
      <div className="inputBox">
        <h3>Key name</h3>
        <select
          id="publicKey"
          name="publicKey"
          title="public key"
          onChange={(e: ChangeEvent<HTMLSelectElement>) => {
            setSelectedKey(e.target.value);
          }}
        >
          {keys.length > 0 &&
            keys.map((key: KeyData, index: number) => {
              return (
                <option key={index} value={key.kid}>
                  {key.kid}
                </option>
              );
            })}
        </select>
        <h3>Encrypted payload</h3>
        <textarea
          id="decryptInput"
          title="payload"
          onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => {
            setPayload(e.target.value);
          }}
        />
        <div className={"buttonContainer"}>
          <Button onClick={() => reset()} variant="secondary">
            Cancel
          </Button>
          <Button onClick={handleEncrypt}>Decrypt &raquo;</Button>
        </div>
      </div>
      {result && (
        <Result title="Decryption result" response={result as Response} />
      )}
    </>
  );
}
