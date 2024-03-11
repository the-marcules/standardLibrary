import {useState, ReactElement} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet, Calculate} from "../wailsjs/go/main/App";
import {Write} from "../wailsjs/go/storage/Storage";
import Menu from "./components/menu/Menu";

function App(): ReactElement {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [lengthOfName, setLengthOfName] = useState("");
    const [feedback, setFeedback] = useState("");

    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    
    function greet() {
        Greet(name).then(updateResultText);
        Write(name).then((feedback)=>{
            setFeedback(feedback)
            setTimeout(()=>{
                setFeedback("")
            }, 5000)
        })
    }
    function len() {
        Calculate(name).then((res) => {
            setLengthOfName(res)
        });
    }

    return (
        <div id="App">
            <div id="result" className="result">{resultText} {lengthOfName}</div>
            <div id="result1" className="result">{feedback}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={greet}>Greet</button>
                <button className="btn" onClick={len}>Length</button>
            </div>
        </div>
    )
}

export default App
