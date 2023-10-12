import db from "database";
import SQLSanatizer from 'utils/SQLSanatizer';
// TODO fix the syntax error on IF statement 
const createToken = async (emailHash: string, token: string) => {
    await db.query(
        "CREATE OR REPLACE FUNCTION create_token(email_var varchar(30), token_var varchar(30)) RETURNS VOID LANGUAGE plpgsql AS $$ BEGIN if NOT EXISTS ( SELECT 1 FROM two_factor_authentication WHERE two_factor_authentication.email = email_var) then INSERT INTO two_factor_authentication (email, token) VALUES (email_var, token_var); else UPDATE two_factor_authentication SET token=token_var WHERE email=email_var; END if; END $$;",
        []
    )
    return db.query(
        "SELECT create_token($1, $2);",
        [SQLSanatizer(emailHash), SQLSanatizer(token)]
    )
}
const validateToken = async (emailHash: string, token: string) => {
    await db.query(
        "CREATE OR REPLACE FUNCTION validate_token(email_var varchar(30), token_var varchar(30)) RETURNS VOID LANGUAGE plpgsql AS $$ BEGIN if EXISTS (SELECT 1 FROM two_factor_authentication WHERE two_factor_authentication.email = email_var AND two_factor_authentication.token=token_var) then DELETE FROM two_factor_authentication WHERE email=email_var; else RAISE EXCEPTION 'Invalid token for email'; END if; END $$;",
        []
    )
    return db.query(
        "SELECT validate_token($1, $2);",
        [SQLSanatizer(emailHash), SQLSanatizer(token)])
}
export { validateToken, createToken };
