
import Elysia from "elysia";
import usersContreoller from "./routeControllers/usersController";
import liabilitiesController from "./routeControllers/liabilitiesController";
import assetsController from "./routeControllers/assetsController";
import investmentsController from "./routeControllers/investmentsController";
import incomeController from "./routeControllers/incomeController";

// TODO rethink this breakdown, we need to get each item by user id so more along the lines of /users/:user_id/assets or something similiar. 
// These so far are just examples of the sql. we should break the logic out of each endpoint into a different file. 
// I also need to rethink the endpoints and think about the use cases on an app.
export default (app: Elysia) => app
            .use(usersContreoller)
            .use(liabilitiesController)
            .use(assetsController)
            .use(investmentsController)
            .use(incomeController)

