import PostgresAdapter from './databaseAdapters/postgresAdapter';
export default class DbFactory {
    private static instance: DBList;

    private constructor() { }

    private static createDb() {
        switch (process.env.DB_TYPE) {
            case 'postgres':
                return new PostgresAdapter(
                    process.env.POSTGRES_USER as string,
                    process.env.POSTGRES_HOST as string,
                    process.env.POSTGRES_DB as string,
                    process.env.POSTGRES_PASSWORD as string,
                    process.env.POSTGRES_PORT as unknown as number
                );
            default:
                throw new Error('Invalid db type');
        }
    }

    public static getInstance(): DBList {
        if (!DbFactory.instance) {
            DbFactory.instance = this.createDb();
        }
        return DbFactory.instance;
    }
}

type DBList = PostgresAdapter;
