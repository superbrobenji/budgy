interface Callback {
    (error: any, results: any): void;
}

export default interface DB {
    query(sql: string, params: string[]): {};
}


