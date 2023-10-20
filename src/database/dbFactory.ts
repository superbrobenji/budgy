import PostgresAdapter from './databaseAdapters/postgresAdapter';

//TDOO create interface for DbFactory
export default class DbFactory {
    private static instance: DBAdapterList;

    private constructor() { }

    private static createDb() {
        switch (process.env.DB_TYPE) {
            case 'postgres':
                return new PostgresAdapter(
                    process.env.POSTGRES_USER,
                    process.env.POSTGRES_HOST,
                    process.env.POSTGRES_DB,
                    process.env.POSTGRES_PASSWORD,
                    process.env.POSTGRES_PORT,
                );
            default:
                throw new Error('Invalid db type');
        }
    }

    public static getInstance(): DBAdapterList {
        if (!DbFactory.instance) {
            DbFactory.instance = this.createDb();
        }
        return DbFactory.instance;
    }
}

type DBAdapterList = PostgresAdapter;
