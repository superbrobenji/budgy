import { Elysia, t } from "elysia";
import { LoginService } from "services/auth/login.service";
import { Tfa } from "services/twoFactorAuthentication/tfa.service";
import { TResult } from "services/types";

export default (app: Elysia) =>
    app.post("/verify-email",
        //@ts-ignore
        async ({ body, set, jwt, setCookie, cookie }) => {
            const { clientToken, email } = body;
            let res: TResult;
            const loginService = new LoginService(email)
            const tfa = new Tfa();
            const verificationData = await tfa.verifyToken(clientToken, email, cookie, setCookie)
            if (verificationData.success) {
                res = await loginService.loginUser(setCookie, jwt)
                set.status = res.status
                return {
                    success: res.success,
                    data: res.data,
                    message: res.message
                }
            } else {
                set.status = 400
                return {
                    ...verificationData,
                    success: false
                }
            }
        },
        {
            body: t.Object({
                clientToken: t.String(),
                email: t.String(),
            }),
        }
    );
