import { TResult } from "services/types";

export interface TTfa {
    createAndSendToken(setCookie: any): Promise<TResult>
    verifyToken(clientToken: string, email: string, cookie: any, setCookie:any): Promise<TResult>
}
