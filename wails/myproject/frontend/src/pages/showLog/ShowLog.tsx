import React, {ReactElement, useEffect, useState} from "react";
import {Read} from "../../../wailsjs/go/storage/Storage";
import styles from './Page.module.scss'

export default function ShowLog(): ReactElement{
    const [contents, setContents] = useState([''])

    useEffect( ():void => {
        Read().then((logContent) => {
            const lines = logContent.split('\n')
            setContents(lines)
        })
    }, [])

    return (
        <div>
            <h1>Log File Content:</h1>
            <ul className={styles.logList}>
                {
                    contents.map((line: string, lineNo: number) => {
                        return (
                            <li key={lineNo}>
                                {line}
                            </li>

                        )
                    })
                }
            </ul>


        </div>

    )
}