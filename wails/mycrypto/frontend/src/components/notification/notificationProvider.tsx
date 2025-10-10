import { createContext, useContext, useEffect, useState } from "react";
import { NotificationProps, NotificationType } from "./notification";
import styles from "./notification.module.css";
import { v4 as uuidv4 } from "uuid";

export type NotificationContextType = {
  addNotification: (
    type: NotificationType,
    title: string,
    message: string,
    ttl?: number
  ) => void;
  removeNotification: (id: string) => void;
  notifications: NotificationMap | undefined;
};

export enum NotificationTTL {
  info = 5000,
  success = 5000,
  warning = 7000,
}

export function SetTTL(ttl: NotificationTTL): number {
  return Date.now() + ttl;
}

export type NotificationContextProps = NotificationProps & {
  ttl?: number;
};

export type NotificationMap = Map<string, NotificationContextProps>;

const notificationContext = createContext<NotificationContextType | undefined>(
  undefined
);
export const useNotification = () => useContext(notificationContext);

export function NotificationProvider({
  children,
}: {
  children: JSX.Element;
}): JSX.Element {
  const [notifications, setNotifications] = useState<NotificationMap>();

  useEffect(() => {
    if (!notifications || notifications.size == 0) {
      return;
    }

    const ttlCheckInterval = setInterval(() => {
      const date = Date.now();
      let needToUpdate = false;
      if (notifications) {
        const newNotifications = new Map(notifications);
        newNotifications.forEach((notification, key) => {
          if (notification.ttl && notification.ttl < date) {
            needToUpdate = true;
            newNotifications.delete(key);
          }
        });
        if (needToUpdate) {
          setNotifications(newNotifications);
        }
        if (newNotifications.size == 0) {
          clearInterval(ttlCheckInterval);
        }
      }
    }, 2000);

    return () => clearInterval(ttlCheckInterval);
  }, [notifications]);

  const removeNotification = (id: string) => {
    if (notifications && notifications?.has(id)) {
      const newNotifications = new Map(notifications);
      newNotifications.delete(id);
      setNotifications(newNotifications);
    }
  };

  const addNotification = (
    type: NotificationType,
    title: string,
    message: string,
    ttl?: number
  ) => {
    const uuid = uuidv4();
    const newNotifications = new Map(notifications);

    newNotifications.set(uuid, { title, message, id: uuid, ttl, type });
    setNotifications(newNotifications);
  };

  return (
    <notificationContext.Provider
      value={{ addNotification, removeNotification, notifications }}
    >
      {children}
    </notificationContext.Provider>
  );
}
