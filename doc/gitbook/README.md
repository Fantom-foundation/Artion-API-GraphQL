# Getting Started

Artion is an NFT marketplace built on Fantom. It provides a platform for minting and trading NFTs - Non-Fungible tokens, as they are defined in ERC-721 spec.

Artion use EVM smart-contracts to allow to trade ERC-721 tokens for ERC-20 tokens.

Purpose of the Artion GraphQL API is to provide high performance datasource for the Artion Client app, however it is open to use or even to develop by anybody. It aggregates data from on-chain [Artion contracts](https://github.com/Fantom-foundation/Artion-Contracts/) and provides them the frontend app in form of standard GraphQL API.

The API server is published on [Artion-API-GraphQL](https://github.com/Fantom-foundation/Artion-API-GraphQL) GitHub repository. Feel free to download and deploy your copy of the application. We are actively developing the project and adding new features so it's worth checking the recent development even if you already checked the repository before.

## Technology

We use several pieces of technology to make the API server work. If your intention is to run your own copy of it, please check our setup and make sure you are familiar with using these technologies at least on rudimentary level before you continue.

First, you will need access to **go-opera** full node RPC interface. You can use properly configured remote one, but it would significantly affect performance and potentially also security of your deployment. Please consider security implications of opening go-opera RPC to outside access. Especially if you enable "personal" commands on your node while keeping your account keys in the Opera key store. We recommend using local IPC channel for communication between an Opera node and the Artion API server. The IPC interface is available by default. Please check the [launch instructions](https://github.com/Fantom-foundation/lachesis\_launch) as well as [go-opera repository](https://github.com/Fantom-foundation/go-opera) for configuration details.

Data aggregated from the chain are kept in [**mongoDB**](https://www.mongodb.com) database. This includes NFTs minted on the observed ERC-721 contracts (collections), listings, offers and auctions created in the Artion marketplace and auction contracts and more.

The database is initialized by the API server, populated with historical data and kept in sync with the full node by internal subscriptions. You don't have to deal with extra database management, but you need to prepare well configured and reliable database connection address to the API server.

Except the per-node local mongoDB database, where the apiserver aggregates data from the chain, you will also need one mongoDB database shared across all apiservers of your marketplace. The shared database can be actually decentralized cluster of mongoDB nodes, however it needs to serve one database accessible by all apiserver nodes.

The shared database is used to store user profiles data, list of approved (or in-review) collections or unlockable contents attached to NFTs.

Except on-chain data, part of NFT tokens are also metadata files and images representing the tokens. Most of the files are stored in IPFS. To acces them, you need to install [local IPFS node](https://github.com/ipfs/go-ipfs) or use IPFS gateway like [Pinata](https://www.pinata.cloud). Pinata is also used to store metadata and images of newly minted tokens, so the apiserver needs valid Pinata JWT access token to be able to help the frontend with minting new tokens.

For sending email notifications to registered users you need to setup [Sendgrid API](https://sendgrid.com) a configure email templates there.
