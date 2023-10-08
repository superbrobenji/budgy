import { TResult } from "services/types";

interface IVerificationCode {
    verifyCode(clientCode: string): TResult
}

export class VerificationCode implements IVerificationCode {
    private static instance: VerificationCode;
    private _verificationCode: string = "";

    private constructor() {
    }

    public static getInstance(): VerificationCode {
        if (!VerificationCode.instance) {
            VerificationCode.instance = new VerificationCode();
        }

        return VerificationCode.instance;
    }

    private generateVerificationCode(): void {
        const chars = "325323432344345";
        const string_length = 6;
        let randomstring = "";
        for (let i = 0; i < string_length; i++) {
            const rnum = Math.floor(Math.random() * chars.length);
            randomstring += chars.substring(rnum, rnum + 1);
        }
        this._verificationCode = randomstring;
    }

    public verifyCode(clientCode: string): TResult {
        const isValid = clientCode == this._verificationCode
        if (!isValid) {
            this.generateVerificationCode();
            return {
                status: 400,
                success: false,
                data: null,
                message: "Email code not verified."
            }
        }
            return {
                status: 200,
                success: true,
                data: null,
                message: 'Email successfully verified'
            }
    }

    get verificationCode() {
        if(this._verificationCode === '') {
            this.generateVerificationCode()
        }
        return this._verificationCode;
    }
}
