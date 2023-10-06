import Elysia from "elysia";
import db from "../../database";

export default (app: Elysia) =>
    app.group("/income", app => app
            .get("/", async () => await db.query("SELECT * FROM monthly_tracking.income", []))
            .get("/:user_id", async ({ params: { user_id } }) => await db.query("SELECT * FROM monthly_tracking.income WHERE monthly_tracking.income.user_id=$1", [user_id]))
    )

