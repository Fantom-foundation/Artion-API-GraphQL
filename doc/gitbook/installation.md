---
description: How to build and configure you own Artion GraphQL API server.
---

# Installation

## Requirements

Before Artion API server installation, you needs to have following stack installed:

* Go 1.16 or newer and `make`
* Locally running go-opera (remote node provides very poor performance, but can be sufficient for testing/development)
* Locally running MongoDB database for data scanned from the chain
* Single MongoDB database shared by all apiservers of your marketplace, where will be approved collections and user profiles stored
* IPFS access point - local IPFS node or IPFS gateway like Pinata (paid access token is necessary to be able to download sufficient amount of files while scanning)
* If you want to be able to mint new tokens, you will need Pinata access token for uploading metadata files and images into IPFS.
* If you want to send notifications to registered users, you will need Sendgrid API access.

## API server installation

When you have all requirements fulfilled, you can install the API server.

First clone the server sources from GitHub:

```
https://github.com/Fantom-foundation/Artion-API-GraphQL.git
```

Now you can build the binary:

```
make
```

The built binary can be found at `build/artionapi`.

Before starting the API server you need to create configuration file, which will connect the API server with everthing in the requirements section.
Feel free to start with config file from the example in `doc` directory - the `config.example.json`.

In the config file:

* Set path to the go-opera IPC file in `node.url`.
* Set connection to the local MongoDB in `db` section.
* Set connection to the MongoDB shared by all nodes in `shared_db` section.
* If you are using local IPFS node, set `ipfs.url` to the host and port of your IPFS node.
* Otherwise, set `ipfs.gateway` to your Pinata HTTP-to-IPFS gateway domain and `ipfs.gateway_bearer` to your Pinata JWT token.
* Set `ipfs.skip_http_gateways` to `true` to use your IPFS node or gateway even for URLs directed to known public HTTP-to-IPFS gateways (like `https://ipfs.io/ipfs/`) - this is recommended to avoid exceeding public gateways limits while scanning the chain.
* Set `ipfs.gateway_bearer` to your Pinata JWT even when you are using your local IPFS node, if you want to mint new tokens using Artion - it will be used for pinning new metadata files and token images. It is neccessary also for uploading user profile pictures.
* Set `notification.sendgrid.key` to your Sendgrid API key if you want to let the server send email notifications to your users. If you are running multiple API server nodes, you should configure this only on one node to avoid sending every notification multiple times.

Before the first start you need to initialize the MongoDB database.
If you want to use other than the official contracts on mainnet, you will need to update `observed.json` appropriately first.

```bash
mongoimport --db=artion --collection=observed --file=doc/db/observed.json
mongoimport --db=artion --collection=status --file=doc/db/status.json
```

For the shared MongoDB database:

```bash
mongoimport --db=artionshared --collection=colcats --file=doc/db/colcats.json
mongoimport --db=artionshared --collection=collections --file=doc/db/collections.json
```

When configured, run the Artion API server:

```bash
build/artionapi -cfg my-config-file.json
```

For production deployment check systemd example in `doc/systemd` to install the api server as systemd service.

As soon as the api server is started, you can access GraphiQL testing interface at http://localhost:7373/graphi.

To connect [Artion-Client-V2](https://github.com/Fantom-foundation/Artion-Client-V2) update the providers list in `app.config.js` to use `http://localhost:7373/graphql`.

