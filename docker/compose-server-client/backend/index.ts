import express, { Response, Request, Application } from "express"
import dotenv from "dotenv"

const app: Application = express()
const config = dotenv.config()

const port  = process.env.PORT

app.use(express.json())
app.get("/", (request: Request, response: Response) => {

    console.log("got request on /", request.body)
    response.status(200)
    response.send('ok')

})

app.listen(port, () => console.log("server listening on port: ", port))