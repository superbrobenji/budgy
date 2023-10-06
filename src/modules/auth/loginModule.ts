import { Elysia, t } from "elysia";
import { getUserLoginDetails } from "resolvers";
import { comparePassword } from "utils/bcrypt";

export default (app: Elysia) =>
    app.post("/login",
        //@ts-ignore
        async ({ body, set, jwt, setCookie }) => {
            const { email, password } = body;
            // verify email/username
            const userArray = await getUserLoginDetails(email);
            const user = userArray[0];
            if (!user) {
                set.status = 400;
                return {
                    success: false,
                    data: null,
                    message: "Invalid credentials",
                };
            }
            // verify password
            const match = await comparePassword(password, user.salt, user.hash);
            if (!match) {
                set.status = 400;
                return {
                    success: false,
                    data: null,
                    message: "Invalid credentials",
                };
            }

            // generate access 
            const accessToken = await jwt.sign({
                userId: user.id,
            });

            setCookie("access_token", accessToken, {
                maxAge: 15 * 60, // 15 minutes
                path: "/",
            });


            return {
                success: true,
                data: null,
                message: "Account login successfully",
            };
        },
        {
            body: t.Object({
                email: t.String(),
                password: t.String(),
            }),
        }
    )

