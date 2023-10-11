import { TResult } from "services/types";

export interface TTfa {
    createAndSendToken(): Promise<TResult>
    verifyToken(clientToken: string, email: string): Promise<TResult>
}
