export class VerificationCode {
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
    public generateVerificationCode(): string {    
        const chars = "0123456789";
        const string_length = 6;
        let randomstring = "";
        for (let i = 0; i < string_length; i++) {
            const rnum = Math.floor(Math.random() * chars.length);
            randomstring += chars.substring(rnum, rnum + 1);
        }
        this._verificationCode = randomstring;
        console.log("Verification Code: ", randomstring);
        return randomstring;
    }
    get verificationCode(): string {
        return this._verificationCode;
    }
}
