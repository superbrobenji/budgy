import { TTfa } from './types';
import { createToken, validateToken } from 'resolvers/twoFactorAuthentication';
import { getUserByEmail } from 'resolvers/users'
import { Mailer } from 'services/email/mailer.service';
import { TResult } from 'services/types';
import Security from 'utils/bcrypt';

export class Tfa implements TTfa {
    private readonly security: Security;

    public constructor() {
        this.security = new Security();
    }
    public async createAndSendToken(setCookie: any, email: string): Promise<TResult> {
        const chars = "239287321905";
        const string_length = 6;
        let randomstring = "";
        for (let i = 0; i < string_length; i++) {
            const rnum = Math.floor(Math.random() * chars.length);
            randomstring += chars.substring(rnum, rnum + 1);
        }
        try {
            await createToken(email, randomstring);
            const res = await this.sendEmail(randomstring, email);

            const emailCrypt = this.security.encrypt(email);
            setCookie("tfa", emailCrypt, {
                // TODO set cookie age to env var
                maxAge: 5 * 60, // 5 minutes
                path: "/",
            });
            return res;
        } catch (err) {
            console.error(err)
            return {
                success: false,
                status: 500,
                data: err,
                message: "Failed to create Token"
            }
        }
    }
    public async verifyToken(clientToken: string, email: string, cookie: any, setCookie: any): Promise<TResult> {
        if (!cookie!.tfa) {
            setCookie("tfa", "", {
                maxAge: -4,
                path: "/",
            })
            return {
                status: 400,
                success: false,
                data: null,
                message: "Token expired. Please try again"
            }
        }
        const emailHash = cookie!.tfa
        const serverEmail = this.security.decrypt(emailHash) as string;
        if (serverEmail !== email) {
            setCookie("tfa", "", {
                maxAge: -4,
                path: "/",
            })
            return {
                status: 400,
                success: false,
                data: null,
                message: "Incorrect token. Please try again"
            }
        }
        try {
            const isValid = await validateToken(email, clientToken);
            if (isValid) {
                setCookie("tfa", "", {
                    maxAge: -4,
                    path: "/",
                })
                return {
                    status: 200,
                    success: true,
                    data: null,
                    message: "Email successfully verified"
                }
            }
        } catch (err) {
            console.error(err)
            setCookie("tfa", "", {
                maxAge: -4,
                path: "/",
            })
            return {
                status: 400,
                success: false,
                data: err,
                message: "Email not verified."
            }
        }
        setCookie("tfa", "", {
            maxAge: -4,
            path: "/",
        })
        return {
            status: 500,
            success: false,
            data: null,
            message: "Error verifying token"
        }
    }
    private async sendEmail(token: string, email: string): Promise<TResult> {
        const mailer = new Mailer();
        const userArray = await getUserByEmail(email);
        const name = userArray[0].name;
        const subject = `Your verification Token - ${token}`
        const body = `<h1> Hello ${name}!</h1> <p> Your verification Token is ${token}. </p>`
        return await mailer.sendEmail(email, subject, body)
    }
} 
