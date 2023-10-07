import { Elysia } from "elysia";

export default (app: Elysia) =>
    app.post("/logout",
        //@ts-ignore
        async ({ setCookie }) => {
            setCookie("access_token", {});
            return {
                success: true,
                data: null,
                message: "Account logout successfully",
            };
        },
    )

