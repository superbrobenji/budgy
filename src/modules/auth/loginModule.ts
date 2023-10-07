import { Elysia, t } from "elysia";
import { getUserLoginDetails } from "resolvers";
import { comparePassword } from "utils/bcrypt";
import validateEmail from "utils/validateEmail";
import validateEmailService from "services/validateEmail.service";
import { VerificationCode } from "services/verificationCode.service";

export default (app: Elysia) =>
    app.post("/login",
        //@ts-ignore
        async ({ body, set, jwt, setCookie }) => {
            const { email, password, verificationCode } = body;
            if (!verificationCode) {
                // verify email
                const isValidEmail = validateEmail(email);
                if (!isValidEmail) {
                    set.status = 400;
                    return {
                        success: false,
                        data: null,
                        message: "Invalid credentials",
                    };
                }
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

                const verifyEmail = await validateEmailService(email);
                if (verifyEmail.success) {
                    return {
                        success: true,
                        data: null,
                        message: "Email verification code sent.",
                    }
                } else {
                    set.status = 400;
                    return {
                        success: false,
                        data: null,
                        message: "Email verification failed. Please try again.",
                    }
                }
            } else {
                if (verificationCode !== VerificationCode.getInstance().verificationCode) {
                    const verifyEmail = await validateEmailService(email);
                    if (verifyEmail.success) {
                        set.status = 400;
                        return {
                            success: false,
                            data: null,
                            message: "Verification failed. A new email verification code sent.",
                        }
                    } else {
                        set.status = 400;
                        return {
                            success: false,
                            data: null,
                            message: "Email verification failed. Please try again.",
                        }
                    }
                }
                const userArray = await getUserLoginDetails(email);
                const user = userArray[0];
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
            }
        },
        {
            body: t.Object({
                email: t.String(),
                password: t.String(),
                verificationCode: t.Optional(t.String({ default: "" })),
            }),
        }
    )

