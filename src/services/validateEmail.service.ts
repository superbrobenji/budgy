
import { InboxResults, sendValidationEmail} from "./sendValidationEmail.service";
import validateEmail from "utils/validateEmail";

interface validateEmailService {
    status: number;
    success: boolean;
    data: InboxResults | null;
    message: string;
}
export default async (email: string): Promise<validateEmailService> => {
    let result: validateEmailService = { status: 400, success: false, data: null, message: '' }
    const isValidEmail = validateEmail(email);
    if (!isValidEmail) {
        result.status = 400;
        result.data = null;
        result.success = false
        result.message = "Invalid email address."
        return result;
    }
    const smptResult = await sendValidationEmail(email);
    console.log('smpt result: ', smptResult);
    if(!smptResult.email_sent){
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
