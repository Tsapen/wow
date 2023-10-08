# Word Of Wisdom

## Overview

The "Word of Wisdom" TCP Server is a robust and secure server designed to provide a collection of inspirational quotes to clients while ensuring protection against Distributed Denial of Service (DDoS) attacks using a [Proof of Work (PoW)](https://en.wikipedia.org/wiki/Proof_of_work) challenge-response protocol. This project aims to deliver a reliable and secure solution for serving inspirational quotes to clients over a TCP connection. 


## Table of Contents
- [Details](#details)
- [Getting Started](#getting-started)
- [Usage](#usage)
  - [Common Instructions](#common-instructions)
  - [Server](#server)
  - [Client](#client)
- [Contributing](#contributing)
- [License](#license)

## Details
"Word of Wisdom" is a secure and efficient TCP Server based on the SHA-256 algorithm.   
**Advantages:**

- **Security Verification:** PoW can be used to verify that the client has performed some computational work before allowing access to a service. This can help mitigate certain types of attacks, such as resource abuse or spam.

- **Resource Costs:** PoW tasks require the client to invest computational resources (CPU time) to compute the hash. This cost can deter malicious users from overwhelming the server with requests.

- **Fairness:** PoW can ensure a fair allocation of resources by requiring clients to prove their commitment or interest in accessing a service, as opposed to simply making requests without any effort.

---


## Getting Started

1. Clone the repository:

   ```shell
   git clone https://github.com/Tsapen/wow.git
   ```

2. Change to the project directory:

   ```shell
   cd wow
   ```

## Usage

1. **Client Connection:**
   - A client establishes a connection with the server.

2. **Challenge Generation:**
   - The server generates a random challenge string.

3. **Challenge Transmission:**
   - The server sends the challenge string to the client.

4. **Client Computes SHA-256:**
   - The client computes the SHA-256 hash of the challenge string.

5. **Client Sends Solution:**
   - The client sends the computed SHA-256 hash (solution) back to the server.

6. **Verification:**
   - The server verifies the received solution.

7. **Response Generation:**
   - If the solution is correct, the server responds with a quote from the "Word of Wisdom" book.

### Common Instructions

Before running the server or client, make sure to follow these common instructions:

#### Run tests
Execute tests to verify the correct installation of all components:

```shell
make test
```

### Server

To run the server in a Docker container:

```shell
make run-server
```

This will build and run the server using the provided configuration.

To stop the server container:

```shell
make stop-server
```

### Client

To initiate a single session for the client in a Docker container:

```shell
make run-client
```

This will build and run the client using the provided configuration. It can be repeated any number of times to simulate multiple interactions.

To stop the client container:

```shell
make stop-client
```
