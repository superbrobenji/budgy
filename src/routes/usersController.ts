import Elysia from "elysia";
import DbFactory from "../database/dbFactory";

const db = DbFactory.getInstance();

export default (app: Elysia) =>
    app.group("/users", app =>
        app
            .get("/", async () => await db.query("SELECT * FROM users", []))
            .get("/:user_id", async ({ params: { user_id } }) => await db.query("SELECT * FROM users WHERE $1=users.id", [user_id]))
    )

