import styles from './result.module.css';

export type ResultProps = {
  response?: Response;
  title?: string;
};

export default function Result(props: ResultProps): JSX.Element {
  if (!props.response) {
    return <></>;
  }

  const copyToClipboard = () => {
    navigator.clipboard.writeText(JSON.stringify(props.response, null, 2));
  };

  return (
    <div className={styles.resultContainer}>
      {props.title && <h3>{props.title}</h3>}
      <pre className={styles.resultBox}>
        {JSON.stringify(props.response, null, 2)}
        <button className={styles.copyButton} onClick={() => copyToClipboard()}>
          Copy
        </button>
      </pre>
    </div>
  );
}
