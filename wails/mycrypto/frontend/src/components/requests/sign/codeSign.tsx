import styles from './sign.module.css'
import { ChangeEvent, useState } from 'react'
import { EngageCodeSignProcess } from '../../../../wailsjs/go/cryptokit/Client'
import Result from '../../result/result'
import { useNotification, NotificationTTL } from '../../notification/notificationProvider'
import Button from '../../Button/Button'

export function CodeSign() {
  const [result, setResult] = useState<Response>()

  const notificationService = useNotification()

  function handleSign() {
    EngageCodeSignProcess().then((response) => {
      const responseObject = JSON.parse(response) as Response
      setResult(responseObject)
      if (responseObject.error.isError) {
        notificationService?.addNotification(
          'error',
          'Error',
          responseObject.error.message || responseObject.description || ''
        )
      } else {
        notificationService?.addNotification('success', 'Success', 'Got a result', Date.now() + NotificationTTL.success)
      }
    })
  }

  return (
    <>
      <div>
        <h3>Payload</h3>
        <Button onClick={handleSign} variant="primary">
          Select file for signing
        </Button>
      </div>

      <Result response={result} title="Response from CryptoKit" />
    </>
  )
}
