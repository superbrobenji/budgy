import DB from '../dbInterface';

const Pool = require('pg').Pool;

export default class PostgresAdapter implements DB {
    private pool: typeof Pool;
    
    private readonly user: string;
    private readonly host: string;
    private readonly database: string;
    private readonly password: string;
    private readonly port: number;

    public constructor(user: string, host: string, database: string, password: string, port: number) {
        this.user = user;
        this.host = host;
        this.database = database;
        this.password = password;
        this.port = port;

        this.connectDb();
    }
    private connectDb() {
        this.pool = new Pool({
            user: this.user,
            host: this.host,
            database: this.database,
            password: this.password,
            port: this.port,
        });
    }
    async query(sql: string, params: any) {
        const res = await this.pool.query(sql, params);
        return res.rows;
    }
}
