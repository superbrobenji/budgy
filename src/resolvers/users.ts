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
const getUserById = (userId: string): Promise<TUser[]> => db.query("SELECT id, name, surname, email_hash FROM users WHERE $1=users.id LIMIT 1", [SQLSanatizer(userId)])
const getUserByEmailHash = (email: string): Promise<TUser[]> => db.query("SELECT id, name, surname, email_hash FROM users WHERE $1=users.email_hash LIMIT 1", [SQLSanatizer(email)])
const getUserIdByEmailHash = (email: string): Promise<{id: string}[]> => db.query("SELECT id FROM users WHERE $1=users.email_hash LIMIT 1", [SQLSanatizer(email)])
const getUserLoginDetails = (email: string): Promise<TUserPassword[]> => db.query("SELECT id, hash, salt FROM users WHERE $1=users.email_hash LIMIT 1", [SQLSanatizer(email)])

const createUser = (name: string, surname: string, email_hash: string, email_salt: string, hash: string, salt: string): Promise<TUser[]> => 
    db.query(
        "INSERT INTO users (name, surname, email_hash, email_salt, hash, salt) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, surname, email_hash",
        [SQLSanatizer(name), SQLSanatizer(surname), SQLSanatizer(email_hash), SQLSanatizer(email_salt), SQLSanatizer(hash), SQLSanatizer(salt)]
    )

export { getUserById, getUserByEmailHash, getUserIdByEmailHash, getUserLoginDetails, createUser };
