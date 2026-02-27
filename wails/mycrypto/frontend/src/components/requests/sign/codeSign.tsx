import styles from './sign.module.css'
import { ChangeEvent, useState } from 'react'
import { Sign as CryptokitSign } from '../../../../wailsjs/go/cryptokit/Client'
import Result from '../../result/result'
import {
  useNotification,
  NotificationTTL,
  SetTTL,
} from '../../notification/notificationProvider'
import Button from '../../Button/Button'

export function CodeSign() {
  const [result, setResult] = useState<Response>()
  const [payload, setPayload] = useState<string>('')

  const updateName = (e: ChangeEvent<HTMLInputElement>) => {
    console.log(e.target.files)
    const file = e.target.files ? e.target.files[0] : null
    if (file) {
      const reader = new FileReader()
      reader.onload = function () {
        console.log(reader.result)
        setPayload(btoa(reader.result as string))
      }
      reader.readAsText(file)
    }
  }

  const notificationService = useNotification()

  function handleSign() {
    if (payload === '') {
      notificationService?.addNotification(
        'warning',
        'Warning',
        'Input cannot be empty. Fill out the payload/file field and try again, you must.',
        SetTTL(NotificationTTL.warning)
      )
      return
    }

    CryptokitSign(payload, true).then((response) => {
      const responseObject = JSON.parse(response) as Response
      setResult(responseObject)
      if (responseObject.error.isError) {
        notificationService?.addNotification(
          'error',
          'Error',
          responseObject.error.message || responseObject.description || ''
        )
      } else {
        notificationService?.addNotification(
          'success',
          'Success',
          'Got a result',
          Date.now() + NotificationTTL.success
        )
      }
    })
  }

  const reset = () => {
    setPayload('')
  }

  return (
    <>
      <div id="input" className={`inputBox`}>
        <h3>Payload</h3>

        <label title="Select File for Code Signing">
          <input
            type="file"
            className={styles.input}
            id="codeSignInput"
            name="file"
            onChange={updateName}
          />
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
  )
}

function codeSigningInput({
  onChange,
}: {
  onChange: (e: ChangeEvent<HTMLInputElement>) => void
}) {
  const [fileName, setFileName] = useState<string>('')

  return (
    <label title="Select File for Code Signing">
      <input
        type="file"
        className={styles.input}
        id="codeSignInput"
        name="file"
      />
    </label>
  )
}
