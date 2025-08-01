import { useLoaderData } from 'react-router';
import Button from '../../Button/Button';
import { useNotification } from '../../notification/notificationProvider';

export default function Base64() {
  const notificationService = useNotification();

  const decodeBase64 = () => {
    const textAreaInput = document.querySelector('textarea#base64Input') as HTMLTextAreaElement;
    const textAreaOutput = document.querySelector('#base64Output') as HTMLTextAreaElement;

    if (textAreaInput && textAreaOutput) {
      const input = textAreaInput.value;
      try {
        textAreaOutput.value = atob(input);
      } catch (error) {
        notificationService?.addNotification('error', 'Error', 'Invalid Base64 string');
      }
    }
  };

  const reset = () => {
    const textAreaInput = document.querySelector('textarea#base64Input') as HTMLTextAreaElement;
    const textAreaOutput = document.querySelector('#base64Output') as HTMLTextAreaElement;

    if (textAreaInput && textAreaOutput) {
      textAreaInput.value = '';
      textAreaOutput.value = '';
    }
  };

  return (
    <>
      <h1>Decode Base64</h1>
      <div className="inputBox">
        <h4>Input</h4>
        <textarea id="base64Input" placeholder="Enter Base64 encoded text here"></textarea>
        <div className={'buttonContainer'}>
          <Button variant="secondary" onClick={reset}>
            Clear
          </Button>
          <Button onClick={() => decodeBase64()}>Decode</Button>
        </div>
        <h4>Output</h4>
        <textarea id="base64Output" placeholder="Decoded text will appear here"></textarea>
      </div>
    </>
  );
}
