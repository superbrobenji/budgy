import DB from '../dbInterface';

const Pool = require('pg').Pool;

export default class PostgresAdapter implements DB {
    private readonly pool: typeof Pool;

    public constructor(user: string, host: string, database: string, password: string, port: number) {
        this.pool = new Pool({
            user: user,
            host: host,
            database: database,
            password: password,
            port: port,
        });
    }
    async query(sql: string, params: any) {
        const res = await this.pool.query(sql, params);
        return res.rows;
    }
}
