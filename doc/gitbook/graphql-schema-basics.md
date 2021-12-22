---
description: Description of the API structure and ideas.
---

# GraphQL Schema Basics

### Basics

The Artion API provides information about the marketplace, about tokens available for trade and about the marketplace users. Key concepts necessary for work with the Artion API are following:

* **Token** or NFT (Non-Fungible Token) is the tradable item.
* **Collection** is an ERC-721 contract - it can host a lot of various tokens.
* **Category** is a set of collections. One collection can be assigned into multiple categories. Categories are defined off-chain.
* **Listing** is an record of the marketplace contract, which allows an owner of a token to offer the token for immediate sale by any buyer for given price.
* **Offer** is an record of the marketplace contract, which can create anybody to offer some amount for given token. The trade will take place if the owner of the token will accept the offer.
* **Auction** is an record of the auction contract, which allows an owner of a token to let potential buyers to compete and sell the token for highest offered price.
* **Bid** is an offer in the auction, given by somebody who wants to buy the auctioned token.
* **Activity** is some action on token done in the chain - created a listing, sold, transfered, etc.
* **PayToken** is an ERC-20 token, which can be used as a currency in trades.

### Data types

The API uses several different scalar values besides the default GraphQL set. Most of them encode big values, or byte arrays of hashes and addresses. If your client application is build using JavaScript, or TypeScript, you can benefit from using excellent [Ethereum Web3.js](https://github.com/ethereum/web3.js/) library, which can help you deal with decoding and validating process of most data types used here.

{% hint style="info" %}
Transaction amounts and basically all amount related data are transferred as big integers encoded in prefixed hexadecimal format. It represents the smallest indivisible amount of value inside the block chain, so called WEI.

One FTM token contains 10¹⁸ WEIs.

One USD is in Artion represented as 10^6.
{% endhint %}

Following scalar values are used by the API:

* **Address** is a 20 bytes Opera address, represented by 0x prefixed hexadecimal number.
* **BigInt** is a large integer value. Input is accepted as either a JSON number, or a hexadecimal number alternatively prefixed with 0x. Output is 0x prefixed hexadecimal.
* **Long** is a 64 bit unsigned integer value represented by 0x prefixed hexadecimal number.
* **Cursor** is an unspecified string value representing position in a sequential list of edges. Please see following section for details about Cursor usage.
* **Time** is string representation of date and time in RFC3339 format.

### Pagination

The API uses cursor based pagination. We didn't use well known and extensively used \<offset, count>, or very similar \<from, to> architecture, where you deal with numeric offset of a collection member from the top of the collection. In our use case the collection top can be changing constantly and keeping track of your position in the collection with ever changing anchor is nearly impossible.

The solution we use is based on a cursor. Each member of the collection has a unique identifier called cursor. When you want to obtain a slice of the collection, you simply specify the cursor value of the collection member and a count of how many neighbors of that anchor you need.

If you do not specify the cursor, we assume you mean either top, or bottom of the collection. The relative anchor is determined by the value of your neighbors count.

For example, this query provides 10 tokens consecutive after the one with cursor `"JQAAABJpbmRleAABwMAAGAAAAAljcmVhdGVkABiXYcV9AQAAAA=="`:

```
{
  tokens(first: 10, after:"JQAAABJpbmRleAABwMAAGAAAAAljcmVhdGVkABiXYcV9AQAAAA==") {
    edges {
      node {
        name
        contract
        tokenId
      }
      cursor
    }
    totalCount
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
  }
}
```

Please note that each response will contain several important information which will help you navigate through the collection.

First is the **pageInfo** structure, defined by **PageInfo** type. It contains the first and last **cursor** of your current slice so you can use these values and ask for next **X** edges, or previous **-X** edges.

Also each **edge** of the list contains not only the desired data element, but also a **cursor**, which can be used to create a link, if desired, to the same position in the collection regardless of the absolute distance of the element from the top, or bottom.
