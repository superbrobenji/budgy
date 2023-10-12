import { TTfa } from './types';
import { createToken, validateToken } from 'resolvers/twoFactorAuthentication';
import { getUserByEmailHash } from 'resolvers/users'
import { Mailer } from 'services/email/mailer.service';
import { TResult } from 'services/types';
import { hashEmail } from 'utils/bcrypt';

export class Tfa implements TTfa {
    private email: string;

    public constructor(email: string) {
        this.email = email;
    }
    public async createAndSendToken(setCookie: any): Promise<TResult> {
        const chars = "239287321905";
        const string_length = 6;
        let randomstring = "";
        for (let i = 0; i < string_length; i++) {
            const rnum = Math.floor(Math.random() * chars.length);
            randomstring += chars.substring(rnum, rnum + 1);
        }
        try {
            const emailCrypt = await hashEmail(this.email);
            await createToken(emailCrypt.hash, randomstring);
            const res = await this.sendEmail(randomstring);
            
            setCookie("tfa", emailCrypt.hash, {
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
    //TODO change this to return the result type
    public async verifyToken(clientToken: string, email: string, cookie: any, setCookie: any): Promise<TResult> {
        console.log("email cookie: ", cookie.tfa)
        if(!cookie!.tfa){
            return {
                status: 400,
                success: false,
                data: null,
                message: "Token expired. Please try again"
            }
        }
       const emailHash = cookie!.tfa
       const clientEmailCrypt = await hashEmail(email)
         if(emailHash !== clientEmailCrypt.hash){
            return {
                status: 400,
                success: false,
                data: null,
                message: "Incorrect token. Please try again"
            }
         }
        try {
            const isValid = await validateToken(clientEmailCrypt.hash, clientToken);
            console.log(isValid)
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
            return {
                status: 400,
                success: false,
                data: err,
                message: "Email not verified."
            }
        }
            return {
                status: 500,
                success: false,
                data: null,
                message: "Error verifying token"
            }
    }
    private async sendEmail(token: string): Promise<TResult> {
        const mailer = new Mailer();
        const emailCrypt = await hashEmail(this.email);
        const userArray = await getUserByEmailHash(emailCrypt.hash);
        const name = userArray[0].name;
        const subject = `Your verification Token - ${token}`
        const body = `<h1> Hello ${name}!</h1> <p> Your verification Token is ${token}. </p>`
        console.log(subject, body, this.email)
        return await mailer.sendEmail(this.email, subject, body)
    }
} 
