import { ReactElement, createContext, useContext, useState } from 'react';
import styles from './notification.module.scss';
import { v4 as uuidv4 } from 'uuid';

export type Action = {
  action: () => void;
  btnLabel: string;
};
export type TNotification = {
  id: string;
  title: string;
  msg: string;
  action?: Action;
};

export type NotificationContextType = {
  notifications: TNotification[];
  addNotification: (newNotification: TNotification[]) => void;
  closeNotification: (id: string) => void;
};

const initialCtx: NotificationContextType = {
  notifications: [],
  addNotification: () => {},
  closeNotification: () => {},
};

export const Ctx = createContext<NotificationContextType>(initialCtx);

export function NotificationContextProvider(props: { children: ReactElement | ReactElement[] }): ReactElement {
  const [notifications, setNotifications] = useState<TNotification[]>([]);

  const addNotification = (newNotifications: TNotification[]) => {
    newNotifications = newNotifications.map((item) => {
      return {
        id: uuidv4(),
        title: item.title,
        msg: item.msg,
        action: item.action,
      };
    });

    setNotifications([...notifications, ...newNotifications]);
  };

  const closeNotification = (id: string) => {
    const notificationElement = document.querySelector(`#id_${id}`);
    notificationElement?.classList.remove(styles.show);
    notificationElement?.classList.add(styles.hide);
    setTimeout(() => {
      setNotifications(notifications.filter((item) => item.id !== id));
    }, 480);
  };

  return (
    <Ctx.Provider value={{ notifications, addNotification, closeNotification }}>
      {notifications.length > 0 && (
        <NotificationLayouter>
          {notifications.map((item) => (
            <Notification {...item} key={item.id} />
          ))}
        </NotificationLayouter>
      )}
      {props.children}
    </Ctx.Provider>
  );
}

function Notification(props: TNotification): ReactElement {
  const NotificationContext = useContext(Ctx);
  return (
    <div className={`${styles.notification} ${styles.show}`} id={`id_${props.id}`}>
      <div className={styles.messageAndTitle}>
        <strong>{props.title}</strong>
        <p>{props.msg}</p>
      </div>
      <div className={styles.buttons}>
        <div>
          <button onClick={() => NotificationContext.closeNotification(props.id)}>Close</button>
        </div>
        {props.action && (
          <div>
            <button
              onClick={() => {
                props.action?.action();
                NotificationContext.closeNotification(props.id);
              }}
            >
              {props.action?.btnLabel}
            </button>
          </div>
        )}
      </div>
    </div>
  );
}

function NotificationLayouter(props: { children: JSX.Element | JSX.Element[] }) {
  return <div className={styles.notificationLayouter}>{props.children}</div>;
}
