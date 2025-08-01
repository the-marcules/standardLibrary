import styles from './notification.module.css';
import { useNotification } from './notificationProvider';

export type NotificationType = 'error' | 'info' | 'success' | 'warning';

export type NotificationProps = {
  type: NotificationType;
  title: string;
  message: string;
  id: string;
};

export default function Notification(props: NotificationProps): JSX.Element {
  const notification = useNotification();

  let cssType;

  switch (props.type) {
    case 'error':
      cssType = styles.error;
      break;
    case 'success':
      cssType = styles.success;
      break;
    case 'warning':
      cssType = styles.warning;
      break;
    default:
    case 'info':
      cssType = styles.info;
      break;
  }

  return (
    <div className={`${styles.notification} ${cssType}`} id={props.id}>
      <h4>{props.title}</h4>
      <p className={styles.message}>{props.message}</p>
      <button onClick={() => notification?.removeNotification(props.id)} className={styles.closeButton}>
        🅧
      </button>
    </div>
  );
}
