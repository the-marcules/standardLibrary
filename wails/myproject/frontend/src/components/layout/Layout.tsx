import React from 'react'
import { HashRouter, Routes, Route } from "react-router-dom";
import styles from './Layout.module.scss'
import App from '../../App'
import ShowLog from '../../pages/showLog/ShowLog'
import Menu from "../../components/menu/Menu";
import OpenFile from "../../pages/open/OpenFile";

export default function Layout(): JSX.Element {
    return (
        <HashRouter basename={"/"}>
            <Menu />
            <div className={styles.content}>
                <Routes>
                    <Route path="/" element={<App/>}/>
                    <Route path="/showLog" element={<ShowLog/>}/>
                    <Route path="/open" element={<OpenFile/>}/>
                </Routes>
            </div>
        </HashRouter>
    )
}