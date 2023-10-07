import { Elysia, t } from "elysia";
import { createUser, getUserByEmail } from "resolvers";
import validateEmailService from "services/validateEmail.service";
import { VerificationCode } from "services/verificationCode.service";
import { hashPassword } from "utils/bcrypt";
import validateEmail from "utils/validateEmail";

export default (app: Elysia) =>
    app.post("/signup",
        async ({ body, set }) => {
            const { email, name, surname, password, verifyPassword, verificationCode } = body;
            if (!verificationCode) {
                const isValidEmail = validateEmail(email);
                if (!isValidEmail) {
                    set.status = 400;
                    return {
                        success: false,
                        data: null,
                        message: "Invalid email address.",
                    };
                }
                const emailExists = await getUserByEmail(email);
                if (emailExists.length > 0) {
                    set.status = 400;
                    return {
                        success: false,
                        data: null,
                        message: "Email address already in use.",
                    };
                }

                //verify password
                if (password !== verifyPassword) {
                    set.status = 400;
                    return {
                        success: false,
                        data: null,
                        message: "Passwords do not match.",
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
                // handle password
                const { hash, salt } = await hashPassword(password);

                const newUser = await createUser(name, surname, email, hash, salt);
                return {
                    success: true,
                    message: "Account created",
                    data: {
                        user: newUser,
                    },
                };
            }
        },
        {
            body: t.Object({
                name: t.String(),
                email: t.String(),
                surname: t.String(),
                password: t.String(),
                verifyPassword: t.String(),
                verificationCode: t.Optional(t.String({ default: "" })),
            }),
        }
    )

