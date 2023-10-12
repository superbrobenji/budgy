import { getUserIdByEmailHash, getUserLoginDetails } from "resolvers/users";
import { Tfa } from "services/twoFactorAuthentication/tfa.service";
import { TResult } from "services/types";
import { comparePassword, hashEmail } from "utils/bcrypt";
import validateEmail from "utils/validateEmail";

interface ILoginService {
    verifyLoginDetails(password: string, setCookie: any): Promise<TResult>
    loginUser(setCookie: any, jwt: any): Promise<TResult>
}
export class LoginService implements ILoginService {
    private readonly email: string

    public constructor(email: string) {
        this.email = email
    }

    public async loginUser(setCookie: any, jwt: any): Promise<TResult> {
        try {
            const emailCrypt = await hashEmail(this.email);
            const userArray = await getUserIdByEmailHash(emailCrypt.hash);
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

    public async verifyLoginDetails(password: string, setCookie: any): Promise<TResult> {
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
        const userArray = await getUserLoginDetails(emailCrypt.hash);
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
        const tfa = new Tfa(this.email) 
        return await tfa.createAndSendToken(setCookie);
    }
}
