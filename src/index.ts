import { Elysia } from "elysia";
import DbFactory from "./database/dbFactory";

const db = DbFactory.getInstance();
const app = new Elysia();

app.get("/", () => "Hello Elysia");
app.get("/users", async () => await db.query("SELECT * FROM users", []));

app.listen(3000, () => {
    console.log(
        `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
    );
})

