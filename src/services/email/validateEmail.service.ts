
import { InboxResults, sendValidationEmail} from "./sendValidationEmail.service";

interface validateEmailService {
    status: number;
    success: boolean;
    data: InboxResults | null;
    message: string;
}
export default async (email: string): Promise<validateEmailService> => {
    let result: validateEmailService = { status: 400, success: false, data: null, message: '' }
    let smptResult = await sendValidationEmail(email);
    if(!smptResult.email_sent){
        for(let i = 0; i < process.env.MAX_EMAIL_VERIFICATION_TRIES; i++){
            smptResult = await sendValidationEmail(email);
            if(smptResult.email_sent){
                break;
            }
        }
        result.status = 400;
        result.data = smptResult;
        result.success = false
        result.message = "Email not sent."
        return result;
    }
    result.status = 200;
    result.data = smptResult;
    result.success = true
    result.message = "Email sent."
    return result;
}
