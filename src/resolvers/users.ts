import db from "database";
import SQLSanatizer from 'utils/SQLSanatizer';

const getUserById = (userId: string) => db.query("SELECT id, name, surname, email FROM users WHERE $1=users.id LIMIT 1", [SQLSanatizer(userId)])
const getUserByEmail = (email: string) => db.query("SELECT id, name, surname, email FROM users WHERE $1=users.email LIMIT 1", [SQLSanatizer(email)])
const getUserLoginDetails = (email: string) => db.query("SELECT id, hash, salt FROM users WHERE $1=users.email LIMIT 1", [SQLSanatizer(email)])
const createUser = (name: string, surname: string, email: string, hash: string, salt: string) => 
    db.query(
        "INSERT INTO users (name, surname, email, hash, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, surname, email LIMIT 1",
        [SQLSanatizer(name), SQLSanatizer(surname), SQLSanatizer(email), SQLSanatizer(hash), SQLSanatizer(salt)]
    )

export { getUserById, getUserByEmail, getUserLoginDetails, createUser };
