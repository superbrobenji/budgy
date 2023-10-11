import { Elysia, t } from "elysia";
import { SignupService } from "services/auth/signup.service";

export default (app: Elysia) =>
    app.post("/signup",
        async ({ body, set }) => {
            const { email, name, surname, password, verifyPassword } = body;
            const signupService = new SignupService(email);
            const verifyData = await signupService.verifySignupDetails(password, verifyPassword)
            if (verifyData.success) {
                const res = await signupService.signupUser(name, surname, password)
                set.status = res.status
                return {
                    message: res.message,
                    data: res.data,
                    success: res.success
                }
            } else {
                set.status = 400
                return {
                    ...verifyData,
                    success: false
                }
            }
        },
        {
            body: t.Object({
                name: t.String(),
                email: t.String(),
                surname: t.String(),
                password: t.String(),
                verifyPassword: t.String(),
            }),
        }
    )

