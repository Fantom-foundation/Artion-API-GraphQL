# Artion-API-GraphQL
GraphQL API server for Fantom Artion v2 - backend for [Artion-Client-V2](https://github.com/Fantom-foundation/Artion-Client-V2).

Build using make:
```bash
make
```

Create JSON config file by `doc/config.example.json` example.

Requirements for run:
* Locally running go-opera - configure IPC file in `node.url` in config file.
* Local MongoDB for data scanned from chain - configure in `db` section of config file.
* MongoDB shared by all Artion nodes for storing users data - `shared_db` section.
* IPFS node for loading token images and token JSON metadata file - two options:
  * Local IPFS node configured in `ipfs.url`
  * Pinata gateway configured in `ipfs.gateway`
* For uploading images into IPFS, Pinata bearer needs to be configured in `ipfs.gateway_bearer` (even when local IPFS node is used otherwise!)
* For sending email notifications, Sendgrid API domain and key needs to be configured in `notification.sendgrid` section.

Before first start you need to initialize the MongoDB database.
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

When configured, run the Artion api server:

```bash
build/artionapi -cfg my-config-file.json
```

For production deployment check systemd example in `doc/systemd` to install the api server as systemd service.

As soon as the api server is started, you can access GraphiQL testing interface at http://localhost:7373/graphi.

To connect Artion-Client-V2 update the providers list in `app.config.js` to use `http://localhost:7373/graphql`.
