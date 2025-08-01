import { Link, NavLink } from 'react-router';
import logo from '../../assets/images/cryptokit-logo-white.svg';
import styles from './header.module.css';
export default function Header() {
  return (
    <div className={styles.header}>
      <img src={logo} id="logo" alt="logo" />
      <nav>
        <ul>
          <li>
            <NavLink to={'/sign'} className={({ isActive }) => isActiveCheck(isActive, styles.isActive)}>
              <span>Sign</span>
            </NavLink>
          </li>
          <li>
            <NavLink to="/verify" className={({ isActive }) => isActiveCheck(isActive, styles.isActive)}>
              <span>Verify</span>
            </NavLink>
          </li>
          <li>
            <NavLink to="/encrypt" className={({ isActive }) => isActiveCheck(isActive, styles.isActive)}>
              <span>Encrypt</span>
            </NavLink>
          </li>
          <li>
            <NavLink to="/decrypt" className={({ isActive }) => isActiveCheck(isActive, styles.isActive)}>
              <span>Decrypt</span>
            </NavLink>
          </li>
        </ul>
      </nav>
    </div>
  );
}

export const isActiveCheck = (isActive: boolean, style: string) => {
  return isActive ? style : '';
};
