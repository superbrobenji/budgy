import Elysia from "elysia";
import db from "../../database";
import { isAuthenticated } from "middlewares/auth";
// TODO find a solution for the types in the auth middleware
export default (app: Elysia) =>
    app.group("/users", app => app
            .get("/", async () => await db.query("SELECT * FROM users", []), { beforeHandle: [isAuthenticated] })
            .get("/:user_id", async ({ params: { user_id } }) => await db.query("SELECT * FROM users WHERE $1=users.id", [user_id]))
    )

