import { emailBody, emailSubject } from 'templates/verificationEmail.template'
const nodemailer = require('nodemailer')
import { VerificationCode } from './verificationCode.service'
import { TInboxResults, TResult } from 'services/types'

interface IValidateEmailService {
    sendVerificationEmail(): Promise<TResult>,
    verifyCode(verificationCode: string, email: string): Promise<TResult>
}

export default class EmailVerification implements IValidateEmailService{
    private result: TResult 
    public constructor() {
        this.result = { status: 400, success: false, data: null, message: '' }
    }

    private async sendEmail(): Promise<TInboxResults> {
        const verificationInstance = VerificationCode.getInstance()
    const verificationCode = verificationInstance.verificationCode;
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
        to: VerificationCode.getInstance().verificationEmail,
        subject: emailSubject(verificationCode),
        html: emailBody(verificationCode),
    }
    try {
        await transporter.sendMail(mailDetails)
        result.email_sent = true
        return {
            email_sent: true
        };
    } catch (err) {
        console.error(err)
        return result
    }
    }

    public async verifyCode(verificationCode: string, email: string): Promise<TResult> {
        const res = VerificationCode.getInstance().verifyCode(verificationCode, email)
        if (res.success) {
            return res
        } else {
            const res = await this.sendVerificationEmail();
            return {
                ...res,
                success: false
            }
        }
    }

    public async sendVerificationEmail(): Promise<TResult> {
    let smptResult = await this.sendEmail();
    if (!smptResult.email_sent) {
        for (let i = 0; i < process.env.MAX_EMAIL_VERIFICATION_TRIES; i++) {
            smptResult = await this.sendEmail();
            if (smptResult.email_sent) {
                break;
            }
        }
        this.result = {
            ...this.result,
            data: smptResult,
            message: "Email verification failed, pelase try again."
        }
        return this.result;
    }
    this.result = {
        status: 200,
        data: smptResult,
        success: true,
        message: "Email verification sent."
    }
    return this.result;
    }
}
