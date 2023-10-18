import { TResult } from 'services/types'
import { TMailer } from './types'
const nodemailer = require('nodemailer')

export class Mailer implements TMailer {
    public constructor() { }
    public async sendEmail(recipientEmail: string, subject: string, body: string): Promise<TResult> {
        const transporter = nodemailer.createTransport({
            service: 'gmail',
            auth: {
                user: process.env.EMAIL_VALIDATION_SENDER_USERNAME,
                pass: process.env.EMAIL_VALIDATION_SENDER_PASSWORD
            },
        })
        let mailDetails = {
            from: process.env.EMAIL_VALIDATION_SENDER,
            to: recipientEmail,
            subject: subject,
            html: body,
        }
        for (let i = 0; i < process.env.MAX_EMAIL_VERIFICATION_ATTEMPTS; i++) {
            try {
                await transporter.sendMail(mailDetails)
                return {
                    status: 200, 
                    success: true,
                    data: null,
                    message: "Email sent successfully"
                }
            } catch (err) {
                console.error(err)
                if (i === process.env.MAX_EMAIL_VERIFICATION_ATTEMPTS - 1) {
                return {
                    status: 500, 
                    success: false,
                    data: err,
                    message: "Email failed to send"
                }
                }
            }
        }
                return {
                    status: 500, 
                    success: false,
                    data: null,
                    message: "Something went wrong on the mailer"
                }
    }
}
