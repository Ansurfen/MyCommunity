import crypto from 'crypto'
let iv = [0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d,
    0x0e, 0x0f];
let key = "It's front  key."
export const EncodeAES = (message: string): string => {
    let md5 = crypto.createHash('md5').update(key).digest('hex');
    const cipher = crypto.createCipheriv(
        'aes-128-cbc',
        Buffer.from(md5, 'hex'),
        Buffer.from(iv)
    );
    // cipher.setAutoPadding(true);
    var encrypted = cipher.update(message, 'utf8', 'base64');
    encrypted += cipher.final('base64');
    return encrypted;
}

export const DecodeAES = (message: string): string => {
    let md5 = crypto.createHash('md5').update(key).digest('hex');
    const decipher = crypto.createDecipheriv(
        'aes-128-cbc',
        Buffer.from(md5, 'hex'),
        Buffer.from(iv)
    );
    var decrypted = decipher.update(message, 'base64', 'utf8');
    decrypted += decipher.final('utf8');
    return decrypted;
}