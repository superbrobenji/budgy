import { Elysia } from "elysia";
import authModules from "modules/auth";
import usersConroller from "modules/routeControllers/usersController";

export default (app: Elysia) =>
    app
        .use(authModules)
        .use(usersConroller)


