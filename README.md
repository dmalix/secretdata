# Secret Data

Easy the encrypting and decrypting a dataset.

### Supported Go versions

Our support of Go versions is aligned with Go's [version release policy](https://golang.org/doc/devel/release#policy).
So we will support a major version of Go until there are two newer major releases.

---

* [Install](#install)
* [Algorithm](#algorithm)
* [Quick Start](#quick-start)
* [Issue Reporting](#issue-reporting)
* [Secure](#secure)
* [Author](#author)
* [License](#license)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/dmalix/secretdata
```

## Algorithm

When encrypting and decrypting a dataset, the secret key is forced hashed to generate a 32-byte key for use with AES-256.
The **Secret Data** library uses AEAD, an encryption mode that provides authenticated encryption of associated data. 
For a description of the methodology, see [https://en.wikipedia.org/wiki/Authenticated_encryption](https://en.wikipedia.org/wiki/Authenticated_encryption)

## Quick Start

Create a new **Secret Data** instance and set a secret key:

```go
secretDataInstance := secretdata.NewSecretData("secret")
```

Encrypt a dataset:

```go
dataset := []byte("dataset")
encryptedData, err := secretDataInstance.Encrypt(dataset)
if err != nil {
   log.Fatal(err)
}
```

And decrypt a dataset:

```go
decryptedData, err := secretDataInstance.Decrypt(encryptedData)
if err != nil {
   log.Fatal(err)
}
```

## Issue Reporting
If you have found a bug or if you have a feature request, please report them at this repository issues section. Please do not report security vulnerabilities on the public GitHub issue tracker.

## Secure
If you discover any security related issues, please email [dmalix@yahoo.com](mailto:dmalix@yahoo.com) instead of using the issue tracker.

## Author
[DmAlix](mailto:dmalix@yahoo.com)

## License
This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for more info.
