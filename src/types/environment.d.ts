export { }

declare global {
    namespace NodeJS {
        interface ProcessEnv {
            DB_TYPE: 'postgres',
            POSTGRES_USER: string,
            POSTGRES_HOST: string,
            POSTGRES_DB: string,
            POSTGRES_PASSWORD: string,
            POSTGRES_PORT: number
        }
    }
}
