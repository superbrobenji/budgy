import Elysia from "elysia";
import db from "database";

export default (app: Elysia) =>
    app.group("/assets", app => app
            .get("/", async () => await db.query("SELECT * FROM owned_resources.assets", []))
            .get("/:user_id", async ({ params: { user_id } }) => await db.query("SELECT * FROM owned_resources.assets WHERE $1=owned_resources.assets.user_id", [user_id]))
    )

