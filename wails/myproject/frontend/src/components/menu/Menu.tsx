import {Link} from "react-router-dom";
import styles from './Menu.module.scss'
import {ReactElement} from "react";

export default function Menu():  ReactElement {
    return (
        <nav className={styles.mainNav}>
            <ul>
                <li>
                    <Link to={'/'}>Home</Link>
                </li>
                <li>
                    <Link to={'/showLog'}>View Logfile</Link>
                </li>   <li>
                    <Link to={'/open'}>Open File</Link>
                </li>
            </ul>
        </nav>
    )
}