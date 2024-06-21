"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const dotenv_1 = __importDefault(require("dotenv"));
const app = (0, express_1.default)();
const config = dotenv_1.default.config();
const port = process.env.PORT;
app.use(express_1.default.json());
app.get("/", (request, response) => {
    console.log("got request on /", request.body);
    response.status(200);
    response.send('ok');
});
app.listen(port, () => console.log("server listening on port: ", port));
//# sourceMappingURL=index.js.map