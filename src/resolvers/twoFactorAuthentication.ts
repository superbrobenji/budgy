import db from "database";
import SQLSanatizer from 'utils/SQLSanatizer';

const createToken = (email: string, token: string) => db.query(
    "IF NOT EXISTS ( SELECT 1 FROM two_factor_authentication WHERE two_factor_authentication.email =$1) BEGIN INSERT INTO two_factor_authentication (email, token) VALUES ($1, $2) END ELSE BEGIN UPDATE two_factor_authentication SET token=$2 WHERE email=$1 END",
    [SQLSanatizer(email), SQLSanatizer(token)]
)
const validateToken = (email: string, token: string) => {
    const tokenExists = db.query("SELECT 1 FROM two_factor_authentication WHERE two_factor_authentication.email =$1 AND two_factor_authentication.token=$2", [SQLSanatizer(email), SQLSanatizer(token)]) 
   return db.query(
    "IF EXISTS ( SELECT 1 FROM two_factor_authentication WHERE two_factor_authentication.email =$1 AND two_factor_authentication.token=$2) BEGIN DELETE FROM two_factor_authentication WHERE email=$1 END ELSE BEGIN RAISERROR('Invalid token', 16, 1) END",
    [email, token])
}
export { validateToken, createToken };
