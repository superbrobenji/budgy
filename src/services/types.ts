export type TInboxResults = {
    email_sent: boolean
}
export type TResult =  {
    status: number;
    success: boolean;
    data: TInboxResults | null | any;
    message: string;
}
