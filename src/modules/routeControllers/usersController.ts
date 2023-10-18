import Elysia from "elysia";
import db from "../../database";
import { isAuthenticated } from "middlewares/auth";

export default (app: Elysia) =>
    app.group("/users", app => app
              //@ts-ignore
            .get("/", async () => await db.query("SELECT * FROM users", []), { beforeHandle: [isAuthenticated] })
            .get("/:user_id", async ({ params: { user_id } }) => await db.query("SELECT * FROM users WHERE $1=users.id", [user_id]))
    )

