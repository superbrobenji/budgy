import { TTfa } from './types';
import { createToken, validateToken } from 'resolvers/twoFactorAuthentication';
import { getUserByEmail } from 'resolvers/users'
import { Mailer } from 'services/email/mailer.service';
import { TResult } from 'services/types';

export class Tfa implements TTfa {
    private readonly email: string;

    public constructor(email: string) {
        this.email = email;
    }
    public async createAndSendToken(): Promise<TResult> {
        const chars = "239287321905";
        const string_length = 6;
        let randomstring = "";
        for (let i = 0; i < string_length; i++) {
            const rnum = Math.floor(Math.random() * chars.length);
            randomstring += chars.substring(rnum, rnum + 1);
        }
        try {
            await createToken(this.email, randomstring);
            return await this.sendEmail(randomstring);
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
    public async verifyToken(clientToken: string, email: string): Promise<TResult> {
        try {
            const isValid = await validateToken(email, clientToken);
            if (isValid) {
                return {
                    status: 200,
                    success: true,
                    data: null,
                    message: "Email successfully verified"
                }
            }
        } catch (err) {
            try {
                await this.createAndSendToken()
            } catch (err) {
                return {
                    status: 500,
                    success: false,
                    data: err,
                    message: "Error creating new token"
                }
            }
            console.error(err)
            return {
                status: 400,
                success: false,
                data: err,
                message: "Email not verified. A new token has been sent."
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
        const userArray = await getUserByEmail(this.email);
        const name = userArray[0].name;
        const subject = `Your verification Token - ${token}`
        const body = `<h1> Hello ${name} Your verification Token is ${token} </h1>`
        return await mailer.sendEmail(this.email, subject, body)
    }
} 
