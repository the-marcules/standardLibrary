import styles from './notification.module.css';
import { useNotification } from './notificationProvider';
import Notification, { NotificationProps } from './notification';

export function NotificationScreen(): JSX.Element {
  const notificationProvider = useNotification();

  return (
    <div className={styles.notificationScreen}>
      {notificationProvider &&
        notificationProvider.notifications &&
        Array.from(notificationProvider.notifications.entries()).map(
          ([id, notification]: [string, NotificationProps]) => {
            return <Notification key={id} {...notification} />;
          }
        )}
    </div>
  );
}
