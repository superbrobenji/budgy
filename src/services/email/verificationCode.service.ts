import { TResult } from "services/types";

interface IVerificationCode {
    verifyCode(clientCode: string, email: string): TResult
}
//TODO change this to normal class instead of singleton
//TODO add expiration time
//TODO in generateVerification code, sae code to table with email and expiration time
//TODO in verifyCode, check if code and email exists in table and if it does, check if it is expired
//TODO if it is expired, generate new code and save it to table and throw error to trigger new email
//TODO if it is not expired, check if code matches and if it does, delete row from table and return success
export class VerificationCode implements IVerificationCode {
    private static instance: VerificationCode;
    private _verificationCode: string = "";
    private email: string = "";

    private constructor() {
    }

    public static getInstance(): VerificationCode {
        if (!VerificationCode.instance) {
            VerificationCode.instance = new VerificationCode();
        }

        return VerificationCode.instance;
    }

    private generateVerificationCode(): void {
        const chars = "239287321905";
        const string_length = 6;
        let randomstring = "";
        for (let i = 0; i < string_length; i++) {
            const rnum = Math.floor(Math.random() * chars.length);
            randomstring += chars.substring(rnum, rnum + 1);
        }
        this._verificationCode = randomstring;
    }

    public verifyCode(clientCode: string, email: string): TResult {
        const isValid = clientCode == this._verificationCode && email == this.email;
        if (!isValid) {
            this.generateVerificationCode();
            return {
                status: 400,
                success: false,
                data: null,
                message: "Email code not verified."
            }
        }
        this._verificationCode = ''
        this.email = ''
        return {
            status: 200,
            success: true,
            data: null,
            message: 'Email successfully verified'
        }
    }
    get verificationEmail() {
        return this.email;
    }
    set setVerificationEmail(email: string) {
        this.email = email;
    }
    get verificationCode() {
        if (this._verificationCode === '') {
            this.generateVerificationCode()
        }
        return this._verificationCode;
    }
}
