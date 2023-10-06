import db from "database";

//TODO add validation on all the parameters
const getUserById = (userId: string) => db.query("SELECT id, name, surname, email FROM users WHERE $1=users.id LIMIT 1", [userId])
const getUserByEmail = (email: string) => db.query("SELECT id, name, surname, email FROM users WHERE $1=users.email LIMIT 1", [email])
const getUserLoginDetails = (email: string) => db.query("SELECT id, hash, salt FROM users WHERE $1=users.email LIMIT 1", [email])
const createUser = (name: string, surname: string, email: string, hash: string, salt: string) => db.query("INSERT INTO users (name, surname, email, hash, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, surname, email LIMIT 1", [name, surname, email, hash, salt])

export { getUserById, getUserByEmail, getUserLoginDetails, createUser };
