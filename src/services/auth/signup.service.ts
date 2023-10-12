import { createUser, getUserByEmailHash } from "resolvers/users";
import { Tfa } from "services/twoFactorAuthentication/tfa.service";
import { TResult } from "services/types";
import { hashEmail, hashPassword } from "utils/bcrypt";
import validateEmail from "utils/validateEmail";

interface ISignupService {
    verifySignupDetails(password: string, verifyPassword: string): Promise<TResult>
    signupUser(name: string, surname: string, password: string, setCookie: any): Promise<TResult>
}
export class SignupService implements ISignupService {
    private readonly email: string

    public constructor(email: string) {
        this.email = email
    }

    public async signupUser(name: string, surname: string, password: string, setCookie: any): Promise<TResult> {
        try {
            const { hash, salt } = await hashPassword(password);
            const emailCrypt = await hashEmail(this.email);
            const newUser = await createUser(name, surname, emailCrypt.hash, emailCrypt.salt, hash, salt);
            const tfa = new Tfa(this.email)
            const verificationEmail = await tfa.createAndSendToken(setCookie);
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
        const emailCrypt = await hashEmail(this.email);
        const emailExists = await getUserByEmailHash(emailCrypt.hash);
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
