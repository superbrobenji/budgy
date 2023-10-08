import { Elysia } from "elysia";
import loginModule from "modules/auth/login.module";
import signUpModule from "modules/auth/signUp.module";
import logoutModule from "modules/auth/logout.module";
import verifyEmailModule from "modules/auth/verifyEmail.module";

export default (app: Elysia) =>
    app.group("/auth", (app) =>
        app
            .use(signUpModule)
            .use(loginModule)
            .use(verifyEmailModule)
            .use(logoutModule)
    );
