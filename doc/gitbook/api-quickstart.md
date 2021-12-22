---
description: Quick introduction into the Artion GraphQL API usage.
---

# Quickstart

You can use Artion GraphQL API to watch activity on yours tokens, to browse any tokens you are interested in or to get any other information from Artion.

This quickstart will guide you through first queries to the Artion GraphQL API. In the first example you will obtain status of auction running on given token, in the second example you will browse list of tokens, filtered and sorted by given fields.

Until you want to run you own Artion based marketplace, you don't need to worry about the API server setup - you can use our public API points:

```
https://artion-api-a.fantom.network/graphql
https://artion-api-b.fantom.network/graphql
```

{% hint style="info" %}
The API allows only to read data aggragated from the chain. If you want to do any operation writing to the chain (create an auction/listing, bid etc.), you will need to sign and send the transaction directly to the Opera RPC, for example using Web3.js library and Metamask.
{% endhint %}

## First example - the auction status

To check status of auction running on token `2937` (`0xb79` in hexa) from collection `0x5dbc2a8b01b7e37dfd7850237a3330c9939d6457`, you can just send following GraphQL API query:

```graphql
{
  token(contract:"0x5dbc2a8b01b7e37dfd7850237a3330c9939d6457", tokenId:"0xb79") {
    name           # name of the token
    auction {
      created      # when was the auction created
      payToken     # ERC-20 token used as the auction currency
      reservePrice # the reserve price of the auction in WEI
      lastBid      # the amount of the last bid in WEI
      lastBidder   # the user who have placed the last bid
    }
  }
}
```

You will get result in format defined by the query:

```json
{
  "data": {
    "token": {
      "name": "Ancestral Uman 2937",
      "auction": {
        "created": "2021-10-09T10:16:02+02:00",
        "payToken": "0x21be370d5312f44cb42ce377bc9b8a0cef1a4c83",
        "reservePrice": "0x61615021",
        "lastBid": null,
        "lastBidder": null
      }
    }
  }
}
```

**Try the example NOW in the integrated GraphiQL interface:**\
****[**https://artion-api-a.fantom.network/graphi**](https://artion-api-a.fantom.network/graphi)****

{% hint style="info" %}
Fast tip: The GraphiQL interface allows to browse the API documentation - just click on the "Docs" button in the right top corner of GraphiQL interface.

When you click on the "History" button, you can browse history of your last queries.
{% endhint %}

## Second example - tokens listing

To browse list of tokens you can use `tokens()` function, which allows to filter, sort and paginate list of all tokens known to Artion.

Following query filters only tokens more expensive than 1 USD, in ascending order, while sorting by creation time and limiting amount of results to five. In other words, it will find 5 oldest tokens above 1 USD.

```graphql
{
  tokens(
    filter: { priceMin: 1000000 }, # price above 1 USD (in 6 decimals fixed point)
    sortBy: CREATED, sortDir: ASC, # sort oldest first
    first: 5, # limit to only 5 results
  ) {
    edges {
      node {
        name
        contract
        tokenId
        created
      }
    }
    pageInfo {
      startCursor # cursor for obtaining the previous page
      endCursor # cursor for obtaining the next page
      hasPreviousPage
      hasNextPage
    }
  }
}
```

This will produce output similar to:

```json
{
  "data": {
    "tokens": {
      "edges": [
        {
          "node": {
            "name": "Bit Uman #34",
            "contract": "0x8c2fcd5d857ee9aa19d1a27ae81ab1129385e3ac",
            "tokenId": "0x22",
            "created": "2021-05-23T16:34:15+02:00"
          }
        },
        {
          "node": {
            "name": "Bit Uman #40",
            "contract": "0x8c2fcd5d857ee9aa19d1a27ae81ab1129385e3ac",
            "tokenId": "0x28",
            "created": "2021-05-23T16:44:06+02:00"
          }
        },
        {
          "node": {
            "name": "Bit Uman #50",
            "contract": "0x8c2fcd5d857ee9aa19d1a27ae81ab1129385e3ac",
            "tokenId": "0x32",
            "created": "2021-05-23T16:59:37+02:00"
          }
        },
        {
          "node": {
            "name": "Bit Uman #52",
            "contract": "0x8c2fcd5d857ee9aa19d1a27ae81ab1129385e3ac",
            "tokenId": "0x34",
            "created": "2021-05-23T17:01:06+02:00"
          }
        },
        {
          "node": {
            "name": "Bit Uman #54",
            "contract": "0x8c2fcd5d857ee9aa19d1a27ae81ab1129385e3ac",
            "tokenId": "0x36",
            "created": "2021-05-23T17:06:09+02:00"
          }
        }
      ],
      "pageInfo": {
        "startCursor": "JQAAABJpbmRleAAvMA0LBwAAAAljcmVhdGVkAFjepZl5AQAAAA==",
        "endCursor": "JQAAABJpbmRleAAM0LULBwAAAAljcmVhdGVkAOgSw5l5AQAAAA==",
        "hasPreviousPage": false,
        "hasNextPage": true
      }
    }
  }
}
```

For more information about pagination - like how to get following five results - check [GraphQL Schema Basics](graphql-schema-basics.md#pagination).

## Sending the GraphQL request

To send a GraphQL request you can use specialized library like [Apollo Client](https://www.apollographql.com/docs/) for JavaScript, or [Graphene](https://github.com/graphql-python/graphene) for Python.

However, you can also send the request manually. Check following example to load data from the API using plain JavaScript:

```javascript
fetch('https://artion-api-a.fantom.network/graphql', {
  method: 'POST',
  body: JSON.stringify({
    query:`
      {
          token(contract:\"0x5dbc2a8b01b7e37dfd7850237a3330c9939d6457\", tokenId:\"0xb79\") {
              name
              auction {
                  created
                  payToken
                  reservePrice
                  lastBid
                  lastBidder
              }
          }
      }
    `
  })
})
.then(response => response.json())
.then(json => {
    console.log(json); // print the result into the console
})
```

In the `json` variable you will get an object identical with the result above.

