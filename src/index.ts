import { Elysia } from "elysia";
import usersContreoller from "./routes/usersController";
import liabilitiesController from "./routes/liabilitiesController";

const app = new Elysia();

app.get("/", () => "Hello Elysia");

app.use(usersContreoller);
app.use(liabilitiesController);

app.listen(3000, () => {
    console.log(
        `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
    );
})

