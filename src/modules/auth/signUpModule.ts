
import { Elysia, t } from "elysia";
import { createUser, getUserByEmail } from "resolvers";
import { hashPassword } from "utils/bcrypt";

export default (app: Elysia) =>
    app.post("/signup",
        async ({ body, set }) => {
            const { email, name, surname, password } = body;
            const emailExists = await getUserByEmail(email);
            if (emailExists.length > 0) {
                set.status = 400;
                return {
                    success: false,
                    data: null,
                    message: "Email address already in use.",
                };
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
        },
        {
            body: t.Object({
                name: t.String(),
                email: t.String(),
                surname: t.String(),
                password: t.String(),
            }),
        }
    )

