
import { createUser, getUserByEmail } from "resolvers/users";
import EmailVerification from "services/email/validateEmail.service";
import { TResult } from "services/types";
import { hashPassword } from "utils/bcrypt";
import validateEmail from "utils/validateEmail";

interface ISignupService {
    verifySignupDetails(password: string, verifyPassword: string): Promise<TResult>
    signupUser(name: string, surname: string, password: string): Promise<TResult>
}
export class SignupService implements ISignupService {
    private readonly email: string
    private readonly emailVerification: EmailVerification

    public constructor(email: string) {
        this.email = email
        this.emailVerification = new EmailVerification()
    }

    public async signupUser(name: string, surname: string, password: string): Promise<TResult> {
        try {
            const { hash, salt } = await hashPassword(password);

            const newUser = await createUser(name, surname, this.email, hash, salt);
            const verificationEmail = await this.emailVerification.sendVerificationEmail();
            if (verificationEmail.success) {
                return {
                    status: 200,
                    success: true,
                    message: "Account created",
                    data: {
                        user: newUser,
                    },
                }
            } else {
                return verificationEmail
            }
        } catch (err) {
            console.error(err)
            return {
                status: 500,
                success: false,
                data: null,
                message: "Failed to create user. Please try again."
            }
        }
    }

    public async verifySignupDetails(password: string, verifyPassword: string): Promise<TResult> {
        const isValidEmail = validateEmail(this.email);
        if (!isValidEmail) {
            return {
                status: 400,
                success: false,
                data: null,
                message: "Invalid credentials",
            };
        }
        const emailExists = await getUserByEmail(this.email);
        if (emailExists.length > 0) {
            return {
                status: 400,
                success: false,
                data: null,
                message: "Email address already in use.",
            };
        }
        // verify password
        if (password !== verifyPassword) {
            return {
                status: 400,
                success: false,
                data: null,
                message: "Passwords do not match.",
            };
        }
        return {
            status: 200,
            success: true,
            data: null,
            message: "Valid credentials",
        }
    }
}
