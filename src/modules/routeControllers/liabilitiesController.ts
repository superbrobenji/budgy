import Elysia from "elysia";
import db from "../../database";

export default (app: Elysia) =>
    app.group("/liabilities", app => app
            .get("/", async () => await db.query("SELECT * FROM owned_resources.liabilities", []))
            .get("/:user_id", async ({ params: { user_id } }) => await db.query("SELECT * FROM owned_resources.liabilities WHERE $1=owned_resources.liabilities.user_id", [user_id]))
    )

