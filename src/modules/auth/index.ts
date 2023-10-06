import { Elysia } from "elysia";
import loginModule from "modules/auth/loginModule";
import signUpModule from "modules/auth/signUpModule";
export default (app: Elysia) =>
    app.group("/auth", (app) =>
        app
            .use(signUpModule)
            .use(loginModule)
    );
