import { useContext, ReactElement } from 'react';
import { Ctx } from '../../context/notificationContext/notificationContext';

export function RandomComponent(): ReactElement {
  const notificationCtx = useContext(Ctx);

  const onclick = () => {
    notificationCtx.addNotification([
      {
        title: 'Hello',
        msg: 'this is just a small reminder to shut the fuck up.',
        id: `id_1`,
      },
      {
        title: 'Sponsoring',
        msg: 'Visit our Partners at http://www.google.de',
        id: `id_2`,
        action: {
          action: () => {
            window.open('http://google.de');
          },
          btnLabel: 'Visit',
        },
      },
    ]);
  };

  return (
    <button type="button" onClick={onclick}>
      add notification
    </button>
  );
}
