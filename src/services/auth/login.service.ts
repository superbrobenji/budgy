import { getUserIdByEmail, getUserLoginDetails } from "resolvers/users";
import EmailVerification from "services/email/validateEmail.service";
import { VerificationCode } from "services/email/verificationCode.service";
import { TResult } from "services/types";
import { comparePassword } from "utils/bcrypt";
import validateEmail from "utils/validateEmail";

interface ILoginService {
    verifyLoginDetails(password: string): Promise<TResult>
    loginUser(setCookie: any, jwt: any): Promise<TResult>
}
export class LoginService implements ILoginService {
    private readonly email: string
    private readonly emailVerification: EmailVerification

    public constructor(email: string) {
        this.email = email
        this.emailVerification = new EmailVerification()
    }

    public async loginUser(setCookie: any, jwt: any): Promise<TResult> {
        try {
            const userArray = await getUserIdByEmail(this.email)
            const userId = userArray[0].id
            // generate access 
            const accessToken = await jwt.sign({
                userId: userId,
            });

            setCookie("access_token", accessToken, {
                maxAge: 15 * 60, // 15 minutes
                path: "/",
            });
        } catch (err) {
            console.error(err)
            return {
                status: 500,
                success: false,
                data: null,
                message: "Failed to fetch user. Please try again."
            }
        }
        return {
            status: 200,
            success: true,
            data: null,
            message: "Account login successfully",
        };
    }

    public async verifyLoginDetails(password: string): Promise<TResult> {
        const isValidEmail = validateEmail(this.email);
        if (!isValidEmail) {
            return {
                status: 400,
                success: false,
                data: null,
                message: "Invalid credentials",
            };
        }
        const userArray = await getUserLoginDetails(this.email);
        const user = userArray[0];
        if (!user) {
            return {
                status: 400,
                success: false,
                data: null,
                message: "Invalid credentials",
            };
        }
        // verify password
        const match = await comparePassword(password, user.salt, user.hash);
        if (!match) {
            return {
                status: 400,
                success: false,
                data: null,
                message: "Invalid credentials",
            };
        }
        const verificationInstance = VerificationCode.getInstance()
        verificationInstance.setVerificationEmail = this.email
        return await this.emailVerification.sendVerificationEmail();
    }
}
