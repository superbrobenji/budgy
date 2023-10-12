import { Elysia, t } from "elysia";
import { LoginService } from "services/auth/login.service";

export default (app: Elysia) =>
    app.post("/login",
        //@ts-ignore
        async ({ body, set, setCookie }) => {
            const { email, password } = body;
            const loginServie = new LoginService(email);
            const res = await loginServie.verifyLoginDetails(password, setCookie)
            set.status = res.status
            return {
                success: res.success,
                data: res.data,
                message: res.message
            }
        },
        {
            body: t.Object({
                email: t.String(),
                password: t.String(),
            }),
        }
    );
