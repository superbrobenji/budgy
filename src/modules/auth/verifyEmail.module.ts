import { Elysia, t } from "elysia";
import { LoginService } from "services/auth/login.service";
import EmailVerification from "services/email/validateEmail.service";
import { TResult } from "services/types";

export default (app: Elysia) =>
        app.post("*/verifyEmail",
                //@ts-ignore
                async ({ body, set, jwt, setCookie }) => {
                    const { verificationCode, email } = body;
                    let res: TResult;
                    const loginService = new LoginService(email)
                    const emailService = new EmailVerification(email);
                    const verificationData = await emailService.verifyCode(verificationCode)
                    if (verificationData.success) {
                        res = await loginService.loginUser(setCookie, jwt)
                        set.status = res.status
                        return {
                            success: res.success,
                            data: res.data,
                            message: res.message
                        }
                    } else {
                        return {
                            ...verificationData,
                            success: false
                        }
                    }
                },
                {
                    body: t.Object({
                        verificationCode: t.String(),
                        email: t.String(),
                    }),
                }
    );
