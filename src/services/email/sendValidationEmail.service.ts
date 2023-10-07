import { emailBody, emailSubject } from 'templates/verificationEmail.template'
import { VerificationCode } from './verificationCode.service'
const nodemailer = require('nodemailer')

export type InboxResults = {
    email_sent: boolean
}

export const sendValidationEmail = async (emailInbox: string): Promise<InboxResults> => {
    const verificationCode = VerificationCode.getInstance().generateVerificationCode();
    const result = { email_sent: false }
    const transporter = nodemailer.createTransport({
        service: 'gmail',
        auth: {
            user: process.env.EMAIL_VALIDATION_SENDER_USERNAME,
            pass: process.env.EMAIL_VALIDATION_SENDER_PASSWORD
        },
    })
    let mailDetails = {
        from: process.env.EMAIL_VALIDATION_SENDER,
        to: emailInbox,
        subject: emailSubject(verificationCode),
        html: emailBody(verificationCode),
    }
    try {
        const info = await transporter.sendMail(mailDetails)
        console.log(info)
        result.email_sent = true
        return {
            email_sent: true
        };
    } catch (err) {
        console.log(err)
        return result
    }
}
