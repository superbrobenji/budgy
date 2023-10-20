import { randomBytes, pbkdf2, createCipheriv, createDecipheriv } from 'crypto';

// TODO create an interface for Security
export default class Security {
    encoding: BufferEncoding = 'hex';

    // process.env.CRYPTO_KEY should be a 32 BYTE key
    key: string = process.env.CRYPTO_KEY || '12345678901234567890123456789012'; 

    encrypt(plaintext: string) {
        try {
            const iv = randomBytes(16);
            const cipher = createCipheriv('aes-256-cbc', this.key, iv);

            const encrypted = Buffer.concat([
                cipher.update(
                    plaintext, 'utf-8'
                ),
                cipher.final(),
            ]);

            return iv.toString(this.encoding) + encrypted.toString(this.encoding);

        } catch (e) {
            console.error(e);
        }
    };

    decrypt(cipherText: string) {
        const {
            encryptedDataString,
            ivString,
        } = this.splitEncryptedText(cipherText);

        try {
            const iv = Buffer.from(ivString, this.encoding);
            const encryptedText = Buffer.from(encryptedDataString, this.encoding);

            const decipher = createDecipheriv('aes-256-cbc', this.key, iv);

            const decrypted = decipher.update(encryptedText);
            return Buffer.concat([decrypted, decipher.final()]).toString();
        } catch (e) {
            console.error(e);
        }
    }

    hashPassword(
        password: string
    ): Promise<{ hash: string; salt: string }> {
        const salt = randomBytes(16).toString("hex");
        return new Promise((resolve, reject) => {
            pbkdf2(password, salt, 1000, 64, "sha512", (error, derivedKey) => {
                if (error) {
                    return reject(error);
                }
                return resolve({ hash: derivedKey.toString("hex"), salt });
            });
        });
    }

    comparePassword(
        password: string,
        salt: string,
        hash: string
    ): Promise<boolean> {
        return new Promise((resolve, reject) => {
            pbkdf2(password, salt, 1000, 64, "sha512", (error, derivedKey) => {
                if (error) {
                    return reject(error);
                }
                return resolve(hash === derivedKey.toString("hex"));
            });
        });
    }
    private splitEncryptedText(encryptedText: string) {
        return {
            ivString: encryptedText.slice(0, 32),
            encryptedDataString: encryptedText.slice(32),
        }
    }
}
