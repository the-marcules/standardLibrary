# Spring App to sign and verify


## what is this?

This is a simple Spring Boot application that provides endpoints to sign and verify messages using RSA keys. It uses the `java.security` package for cryptographic operations.

## How to run
### using Maven
```bash
mvn clean install
mvn spring-boot:run
```
### using IDE
start the application using the start application button in upper right corner.
## how to build
### using Maven

```bash 
mvn clean package
```

then start the application using the following command:
```bash
java -jar target/app.jar
```
> ℹ️ The name of the jar file is defined in the `pom.xml` file in the `<finalName>` Tag within the `<build>` section.

## How to use
### Sign a message
```bash
curl -X POST http://localhost:8080/sign -H "Content-Type: application/json" -d '{"message": "Hello, World!"}'
```
### Verify a signature
you can use the response body from the sign endpoint as request body to verify the signature (just copy & paste).
```bash
curl -X POST http://localhost:8080/verify -H "Content-Type: application/json" -d '{"signedPayload": "<signedPayload>", "certificate": "<certificate>", "payload": "<payload>"}'
```

## how does it work?

in the static folder of the resources, there are two files:
- `sign.key`: contains the private key used to sign messages.
- `sing.cert`: contains the cert with pub key and identity information used to verify signatures.
the application reads these files at startup and uses them to sign and verify messages.

> ⚠️ **Important** ⚠️: The private key is neither encrypted nor in a safe place, so this is not a production ready application. **Keys should never be commited and pushed to vcs**. In this case this is done for the sake of simplicity and testing purposes only.

