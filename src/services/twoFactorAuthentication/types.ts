import { TResult } from "services/types";

export interface TTfa {
    createAndSendToken(setCookie: any, emailCrypt: string): Promise<TResult>
    verifyToken(clientToken: string, email: string, cookie: any, setCookie:any): Promise<TResult>
}
