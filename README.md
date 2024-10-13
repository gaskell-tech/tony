# tony
tony is a utils application.

## encryptAES

Reads and encrypts file content; printing ciphertext to stdout.

The file can be provided as either a relative or absolute path.

The ciphertext format: 'AES256:<base64-encoded-data>`

```bash
tony.exe e -p \<password\> -f \<file\>
```

## decryptAES

Decrypts ciphertext; printing cleartext to stdout.

The ciphertext format: 'AES256:<base64-encoded-data>`

```bash
tony.exe d -p \<password\> \<ciphertext\>
```
