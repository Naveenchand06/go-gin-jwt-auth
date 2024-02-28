# Go Gin JWt Auth

### ECDSA Public & Private Keys using OpenSSL

To generate ECDSA keys pairs, We have to create private key first

```sh
openssl ecparam -name secp256k1 -genkey -noout -out <output_filename>
```

This will generate a file with .pem extension

Now, From this private key we can generate public key

```sh
openssl ec -in <private_key_file> -pubout > <output_filename>
```

This will output public key file
