import express, { Response, Request, Application } from "express"
import dotenv from "dotenv"

const config = dotenv.config()
const port  = process.env.PORT


const app: Application = express()
app.use(express.json())
app.post("/", (request: Request, response: Response) => {

    console.log("got request on /", JSON.stringify(request.body))
    response.status(200)
    response.send('ok')

})

app.listen(port, () => console.log("server listening on port: ", port))