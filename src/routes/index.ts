
import Elysia from "elysia";
import usersContreoller from "./routeControllers/usersController";
import liabilitiesController from "./routeControllers/liabilitiesController";

export default (app: Elysia) => app
            .use(usersContreoller)
            .use(liabilitiesController)

