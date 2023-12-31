import { createUser, getUserByEmail } from "resolvers/users";
import { Tfa } from "services/twoFactorAuthentication/tfa.service";
import { TResult } from "services/types";
import Security from "utils/bcrypt";
import validateEmail from "utils/validateEmail";

//TODO add interface to types file
interface ISignupService {
    verifySignupDetails(password: string, verifyPassword: string): Promise<TResult>
    signupUser(name: string, surname: string, password: string, setCookie: any): Promise<TResult>
}
export class SignupService implements ISignupService {
    private readonly email: string
    private readonly security: Security;

    public constructor(email: string) {
        this.email = email
        this.security = new Security()
    }

    public async signupUser(name: string, surname: string, password: string, setCookie: any): Promise<TResult> {
        try {
            const { hash, salt } = await this.security.hashPassword(password);
            const newUser = await createUser(name, surname, this.email, hash, salt);
            const tfa = new Tfa()
            const verificationEmail = await tfa.createAndSendToken(setCookie, this.email);
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
