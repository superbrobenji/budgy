import db from "database";
import SQLSanatizer from 'utils/SQLSanatizer';
type TUser = {
    id: string,
    name: string,
    surname: string,
    email: string
}
type TUserPassword = {
    id: string,
    hash: string,
    salt: string
}
const getUserById = (userId: string): Promise<TUser[]> => db.query("SELECT id, name, surname, email FROM users WHERE $1=users.id LIMIT 1", [SQLSanatizer(userId)])
const getUserByEmail = (email: string): Promise<TUser[]> => db.query("SELECT id, name, surname, email FROM users WHERE $1=users.email LIMIT 1", [SQLSanatizer(email)])
const getUserIdByEmail = (email: string): Promise<{id: string}[]> => db.query("SELECT id FROM users WHERE $1=users.email LIMIT 1", [SQLSanatizer(email)])
const getUserLoginDetails = (email: string): Promise<TUserPassword[]> => db.query("SELECT id, hash, salt FROM users WHERE $1=users.email LIMIT 1", [SQLSanatizer(email)])

const createUser = (name: string, surname: string, email: string, hash: string, salt: string): Promise<TUser[]> => 
    db.query(
        "INSERT INTO users (name, surname, email, hash, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, surname, email",
        [SQLSanatizer(name), SQLSanatizer(surname), SQLSanatizer(email), SQLSanatizer(hash), SQLSanatizer(salt)]
    )

export { getUserById, getUserByEmail, getUserIdByEmail, getUserLoginDetails, createUser };
