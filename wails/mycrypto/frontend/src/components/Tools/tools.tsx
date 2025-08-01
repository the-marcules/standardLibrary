import { NavLink } from 'react-router';
import styles from './tools.module.css';
import { isActiveCheck } from '../header/header';

export default function Tools(): JSX.Element {
  return (
    <div className={styles.toolsContainer}>
      <div>TOOLS</div>
      <ul>
        <li>
          <NavLink to={'/tools/base64'} className={({ isActive }) => isActiveCheck(isActive, styles.isActive)}>
            Base64
          </NavLink>
        </li>
      </ul>
    </div>
  );
}
