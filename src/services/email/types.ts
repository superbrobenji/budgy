import { TResult } from "services/types";

export interface TMailer {
    sendEmail(email: string, subject: string, message: string): Promise<TResult>;
}
