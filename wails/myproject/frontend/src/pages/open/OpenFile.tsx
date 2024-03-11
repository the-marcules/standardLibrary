import styles from './OpenFile.module.scss'
import React, {ReactElement, useEffect, useState} from "react";
import {OpenFileDialog} from "../../../wailsjs/go/fileOperation/FileOperation";
import {BrowserOpenURL, LogError} from "../../../wailsjs/runtime";


export default function OpenFile(): ReactElement {
    const [files, setFiles] = useState<any[]>([])

    return(<div>
        <button onClick={()=>{
            OpenFileDialog()
                .then(res => JSON.parse(res))
                .then(response => {
                setFiles(response)
            })
                .catch(err => {
                    LogError(err)
                })

        }}>Open File(s)</button>
        <div>
            {
                files.map((file, i) => {
                    return  <ExifDisplay fileInfo={file.ExifInfo} itemNo={"ef_parent_"+i} title={file.FilePath}></ExifDisplay>
                })
            }
        </div>
    </div>)
}

interface Location {
    Latitude: string
    Longitude: string
}

function ExifDisplay(props: {fileInfo: string, itemNo: string, title?: string}): ReactElement {
    if ( !props.fileInfo ) return <></>
    return (
        <div key={props.itemNo}>
        {props.title && <h2 className={styles.h2}>{props.title}</h2>}
            <ul>
                {
                    Object.entries(props.fileInfo)
                    .map(([key, value], index) => {
                        switch (typeof value) {
                            case 'object':
                                if (key == "Location") {
                                    const loc = value as Location
                                    return <li key={'ed_'+index}>{`${key}: `}<MapsLink lat={loc.Latitude} lon={loc.Longitude}></MapsLink></li>
                                }
                                return (
                                    <li key={index}>
                                        {key}
                                        <ExifDisplay fileInfo={value} itemNo={"ef_child_"+index}></ExifDisplay>
                                    </li>
                                    )
                            default:
                            case 'string':
                                return <li key={'ed_'+index}>{`${key}: ${value}`}</li>
                        }
                    })
                }
            </ul>
        </div>
    )
}

function MapsLink(props: {lat: string, lon: string, title?: string}): ReactElement {
    if ( props.lat != '0' && props.lon != '0') {
        return <button onClick={()=>BrowserOpenURL(`https://www.google.com/maps/search/${props.lat},${props.lon}`)}>{props.title?props.title:"See on Google Maps"}</button>
    }
    return <span>No Location found.</span>
}