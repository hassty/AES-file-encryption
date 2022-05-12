## About

Cli utility to encrypt or decrypt files using aes-cbc

## Usage

### Help

To get help type

```sh
./aes --help
```

or

```sh
./aes [command] --help
```

### Encryption

```sh
./aes encrypt --iv <initialization vector> --keyfile <path to aes key> [file]
```

If no file is specified, it reads from stdin

### Decryption

```sh
./aes decrypt --iv <initialization vector> --keyfile <path to aes key> [file]
```

If no file is specified, it reads from stdin

### Example

You can find usage example in [run.sh](./run.sh)
