import dotenv from "dotenv"

const config = dotenv.config()
const port  = process.env.BACKEND_PORT


const int = setInterval(() => {
    const url = `http://backend:${port}`
    console.log(`trying to fetch from ${url}`)
    fetch(url, {method: "POST", body:JSON.stringify({"greeting": "Hello, world!"})})
    .then(result => result.text())
    .then(data => console.log(`got data: ${data}`))
    .catch(err => console.error(`got err: ${err}`)) 
}, 5000)