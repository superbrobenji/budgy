import Elysia from "elysia";
import db from "../../database";

export default (app: Elysia) =>
    app.group("/investments", app => app
            .get("/", async () => await db.query("SELECT * FROM owned_resources.investments", []))
            .get("/:user_id", async ({ params: { user_id } }) => await db.query("SELECT * FROM owned_resources.investments INNER JOIN owned_resources.assets ON owned_resources.assets.id=owned_resources.investments.asset_id WHERE owned_resources.assets.user_id=$1", [user_id]))
    )
