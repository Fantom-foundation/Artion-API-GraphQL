---
description: Full schema definition files.
---

# Schema Structure

```graphql
# Root schema definition
schema {
    query: Query
    mutation: Mutation
    subscription: Subscription
}

# Entry points for querying the API
type Query {
    # version represents the API server version responding to your requests.
    version: String!

    # Get primary used Artion contracts addresses.
    contracts: Contracts!

    # Get list of token collections categories.
    categories: [Category!]!

    # Get collection by address.
    collection(contract: Address!): Collection

    # List collections (all, search name or filter mintable by given user).
    collections(search: String, mintableBy: Address, first: Int, after: Cursor, last: Int, before: Cursor): CollectionConnection!

    # Get token by address and id.
    token(contract: Address!, tokenId: BigInt!): Token

    # List all tokens (with defined filter and sorting).
    tokens(filter: TokenFilter, sortBy: TokenSorting, sortDir: SortingDirection, first: Int, after: Cursor, last: Int, before: Cursor): TokenConnection!

    # List of tokens supported for payments on the marketplace.
    payTokens: [PayToken!]!

    # Get user by address.
    user(address: Address!): User!

    # List or search users
    users(search: String, first: Int, after: Cursor, last: Int, before: Cursor): UserConnection!

    # randomTrade resolves a Random ERC-721 NFT Trade by address.
    randomTrade(contract: Address!): RandomTrade

    # Get user authenticated using bearer token.
    loggedUser: User

    # Get currently set shipping address for tokens redeem.
    loggedUserShippingAddress: ShippingAddress

    # Get notification settings of logged user.
    notificationSettings: NotificationSettings

    # Get un-lockable content attached to a NFT token (only token owner).
    unlockableContent(contract: Address!, tokenId: BigInt!): String

    # Estimate platform fee and gas needed for token minting.
    estimateMintFeeGas(user: Address!, contract: Address!, tokenUri: String!): MintFeeGas!

    # Check if the operator has ApprovedForAll permission to manipulate with tokens of given owner.
    isApprovedForAll(contract: Address!, owner: Address!, operator: Address!): Boolean!

    # Search collections, tokens, and users for the given phrase and return list of relevant results.
    textSearch(phrase: String!, maxLength: Int = 15): [TextSearchEdge]!

    # Check it the logged user is a moderator.
    isLoggedModerator: Boolean!

    # List collections to be reviewed by a moderator. (moderators only)
    collectionsInReview(search: String, first: Int, after: Cursor, last: Int, before: Cursor): CollectionConnection!

    # List collections which was banned by a moderator. (moderators only)
    bannedCollections(search: String, first: Int, after: Cursor, last: Int, before: Cursor): CollectionConnection!

    # List tokens which was banned by a moderator. (moderators only)
    bannedTokens(first: Int, after: Cursor, last: Int, before: Cursor): BannedNftConnection!
}

# Mutation endpoints for modifying the data
type Mutation {
    # Generate login challenge to be signed by private key and used to log-in
    initiateLogin: String!

    # Use private key signed login challenge to get bearer token.
    login(user: Address!, challenge: String!, signature: String!): String!

    # Update user profile of logged user
    updateUser(username: String, bio: String, email: String!): Boolean!

    # Update notification settings of logged user
    updateNotificationSettings(settings: InputNotificationSettings!): Boolean!

    # Add token into favourite tokens of logged user.
    likeToken(contract: Address!, tokenId: BigInt!): Boolean!

    # Remove token from favourite tokens of logged user.
    unlikeToken(contract: Address!, tokenId: BigInt!): Boolean!

    # Add user into users followed by the logged user.
    followUser(user: Address!): Boolean!

    # Remove user from users followed by the logged user.
    unfollowUser(user: Address!): Boolean!

    # Update shipping address for tokens redeem.
    updateShippingAddress(address: InputShippingAddress!): Boolean!

    # Set unlockable content attached to a NFT token (only token owner and only once)
    setUnlockableContent(contract: Address!, tokenId: BigInt!, content: String!): Boolean!

    # Increment amount of views of the token.
    incrementTokenViews(contract: Address!, tokenId: BigInt!): Boolean!

    # Approve the in-review collection (moderators only)
    approveCollection(contract: Address!): Boolean!

    # Decline the in-review collection (moderators only)
    declineCollection(contract: Address!): Boolean!

    # Ban the collection (moderators only)
    banCollection(contract: Address!): Boolean!

    # Unban the collection (moderators only)
    unbanCollection(contract: Address!): Boolean!

    # Ban the token (moderators only)
    banToken(contract: Address!, tokenId: BigInt!): Boolean!

    # Unban the token (moderators only)
    unbanToken(contract: Address!, tokenId: BigInt!): Boolean!
}

type Subscription {
    # Subscribe events relevant for given user (owned auction bid, etc.)
    watchUserEvents(user: Address!): Event!

    # Subscribe auction events
    watchAuction(contract: Address!, tokenId: BigInt!): Event!
}
```

## Activities

```graphql
# ActivityType represents type of Activity on a token.
enum ActivityType {
    UNKNOWN

    # Listing created (offered for sale for given price by the token owner)
    LISTING_CREATED

    # Listing updated (changed price)
    LISTING_UPDATED

    # Listing cancelled
    LISTING_CANCELLED

    # Token sold in a listing
    LISTING_SOLD

    # Offer created by a buyer
    OFFER_CREATED

    # Offer cancelled
    OFFER_CANCELLED

    # Token sold by an offer
    OFFER_SOLD

    # Auction created
    AUCTION_CREATED

    # New bid in the auction
    AUCTION_BID

    # Bid cancelled (only when the auction allows is)
    AUCTION_BID_WITHDRAW

    # Auction cancelled
    AUCTION_CANCELLED

    # Auction finished - token sold in the auction
    AUCTION_RESOLVED

    # Auction updated (changed time range or reserve price)
    AUCTION_UPDATED

    # Token transfer
    TRANSFER

    # Token minted (created)
    MINT

    # Token burned (redeemed)
    BURN
}

# Activity represents an event that happened on a market-sellable NFT token.
type Activity {
    # The time of the event
    time: Time!

    # Type of the activity
    type: ActivityType!

    # The token contract
    contract: Address!

    # The token id
    tokenId: BigInt!

    # The amount of tokens (always 1 for ERC-721)
    quantity: BigInt

    # The address of initiator of the activity (auction bidder, offer proposer, seller of listed item)
    from: Address!

    # The initiator of the activity (auction bidder, offer proposer, seller of a listed item)
    fromUser: User!

    # The address of receiver of the activity (buyer of a listed item, offer acceptor, item owner for bids)
    to: Address

    # The receiver of the activity (buyer of a listed item, offer acceptor, item owner for bids)
    toUser: User

    # The pay token of unitPrice
    payToken: Address

    # The price for one piece of the token (total price = unitPrice * quantity)
    unitPrice: BigInt

    # Start time (of the auction/listing)
    startTime: Time

    # End time (of auction/offer)
    endTime: Time

    # The related token
    token: Token

    # The FantomMarketplace contract (for offers/listings)
    marketplace: Address

    # The FantomAuction contract (for auctions)
    auctionHall: Address
}

type ActivityEdge {
    cursor: Cursor!
    node: Activity!
}

type ActivityConnection {
    # Edges contains provided edges of the sequential list.
    edges: [ActivityEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}

input ActivityFilter {
    types: [ActivityType!]
}

# Price history is aggregation of trades from one day
type PriceHistory {
    # Time of the day
    time: Time!

    # Average token price in USD, 6-decimals fixed point, hex
    price: Long!
}
```

## Auctions

```graphql
# Auction represents an auction being conducted on a NFT token.
type Auction {
    # Address of the token contract
    contract: Address!

    # ID of the token (in given token contract)
    tokenId: BigInt!

    # The seller of the item
    owner: Address!

    # The FantomAuction contract
    auctionHall: Address!

    # The auctioned amount of tokens (always 1 for ERC-721)
    quantity: BigInt!

    # The token used to pay for the auctioned item (zeros for native token)
    payToken: Address!

    # Starting price - minimal initial bid
    reservePrice: BigInt!

    # Minimal next bid (last bid + minimal increment, or the reserve price)
    minBidAmount: BigInt!

    # When was the auction created
    created: Time!

    # When the auction starts (it is not possible to bid before this time)
    startTime: Time!

    # When is the planned end of the auction
    endTime: Time!

    # When was the auction closed - resolved or cancelled (null if not closed)
    closed: Time

    # The last bid amount
    lastBid: BigInt

    # When was the last bid placed
    lastBidPlaced: Time

    # Who has placed the last bid
    lastBidder: Address

    # Winner of the auction (null if not resolved)
    winner: Address

    # When was the auction resolved (null if not resolved)
    resolved: Time

    # Is the auction creator owner of the token yet?
    isActive: Boolean!
}
```

## Banned NFTs (moderators)

```graphql
# BannedNft represents a banned NFT.
type BannedNft {
    # Address of the token contract
    contract: Address!

    # ID of the token (in given token contract)
    tokenId: BigInt!

    # The liked token
    token: Token
}

type BannedNftEdge {
    cursor: Cursor!
    node: BannedNft!
}

type BannedNftConnection {
    # Edges contains provided edges of the sequential list.
    edges: [BannedNftEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}
```

## Categories of collections

```graphql
# Category is set of token collections.
type Category {
    # Identifier of the category
    id: Int!

    # Name of the category
    name: String!

    # Icon of the category - SVG file content
    icon: String
}
```

## Collections

```graphql
# Collection represents tokens contract.
type Collection {
    # Address of the token contract
    contract: Address!

    # Name of the token contract
    name: String!

    # Description of the token contract
    description: String!

    # Categories IDs of the collection
    categories: [Int!]!

    # IPFS hash of the collection image
    image: String!

    # Owner of the collection
    owner: Address

    # Owner of the collection
    ownerUser: User

    # Recipient of royalty transfer fee
    feeRecipient: Address

    # Recipient of royalty transfer fee
    feeRecipientUser: User

    # Royalty - fee in percents (with decimals)
    royalty: String!

    # Social: E-mail
    email: String!

    # Social: Site URL
    site: String!

    # Social: Discord URL
    discord: String!

    # Social: Telegram URL
    telegram: String!

    # Social: Medium URL
    medium: String!

    # Social: Twitter URL
    twitter: String!

    # Social: Instagram URL
    instagram: String!

    # canMint checks if the given user address can create new tokens on the collection
    canMint(user:Address!, fee: BigInt): Boolean!
}

type CollectionEdge {
    cursor: Cursor!
    node: Collection!
}

type CollectionConnection {
    # Edges contains provided edges of the sequential list.
    edges: [CollectionEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}
```

## Core contracts

```graphql
# Primary used Artion contracts addresses.
type Contracts {
    marketplace: Address!
    auctionHall: Address!
}
```

## Events for subscriptions

```graphql
# Event represents something what should be reported to a subscribed client.
type Event {
    type: EventType!
    auction: Auction
    offer: Offer
}

# EventType represent a type of event reported to a subscribed client.
enum EventType {
    AUCTION_BID,
    AUCTION_BID_WITHDRAW,
    AUCTION_RESERVE_UPDATED,
    AUCTION_RESOLVED,
    AUCTION_CANCELLED,
    GOT_OFFER,
    TRANSFER,
}
```

## Following

```graphql
# Follow represents "following" of one user by another user.
type Follow {
    # The follower address
    follower: Address!

    # The follower
    followerUser: User!

    # The followed user address
    followed: Address!

    # The followed user
    followedUser: User!
}

type FollowEdge {
    cursor: Cursor!
    node: Follow!
}

type FollowConnection {
    # Edges contains provided edges of the sequential list.
    edges: [FollowEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}
```

## Fulltext search

```graphql
# TextSearchEdge is a single edge in a paginated text search output set.
type TextSearchEdge {
    # The collection, if the item is a collection.
    collection: Collection

    # The token, if the item is a token.
    token: Token

    # The user, if the item is an user.
    user: User
}
```

## Listings

```graphql
# Listing represents offer for anybody to buy given token from the owner.
type Listing {
    # The seller of the item
    owner: Address!

    # The seller of the item
    ownerUser: User!

    # Address of the token contract
    contract: Address!

    # ID of the token (in given token contract)
    tokenId: BigInt!

    # The listed token (detail)
    token: Token

    # The FantomMarketplace contract
    marketplace: Address!

    # The listed amount (only this exact amount can be bought)
    quantity: BigInt!

    # The token used to pay for the listed item (zeros for native token)
    payToken: Address!

    # The price for one unit of the token (total = unitPrice * quantity)
    unitPrice: BigInt!

    # When sale of the item starts
    startTime: Time!

    # When was the item sold or unlisted
    closed: Time

    # Is the listing creator owner of the token yet?
    isActive: Boolean!
}

type ListingEdge {
    cursor: Cursor!
    node: Listing!
}

type ListingConnection {
    # Edges contains provided edges of the sequential list.
    edges: [ListingEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}
```

## PageInfo

```graphql
# PageInfo contains information about a sequential access list page.
# Specified by https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo {
    # startCursor is the cursor of the first edge of the edges list. null for empty list.
    startCursor: Cursor

    # endCursor if the cursor of the last edge of the edges list. null for empty list.
    endCursor: Cursor

    # hasNextPage specifies if there is another edge after the last one.
    hasNextPage: Boolean!

    # hasPreviousPage specifies if there is another edge before the first one.
    hasPreviousPage: Boolean!
}
```

## Minting fee and gas estimation

```graphql
# Platform fee and estimated gas for mint operation
# Return type of estimateMintFeeGas operation.
type MintFeeGas {
    # The error message, if the minting has failed in the estimating simulation.
    error: String

    # The platform fee for the token minting in native FTM tokens.
    platformFee: BigInt

    # The gas estimated to be necessary for the minting.
    gas: BigInt
}
```

## Notification settings

```graphql
# NotificationSettings represents notification settings for a user.
type NotificationSettings {
    # Your Activity Notifications
    sNotification: Boolean!
    # You have purchased a bundle.
    sBundleBuy: Boolean!
    # Your bundle is sold.
    sBundleSell: Boolean!
    # You get an offer for your bundle.
    sBundleOffer: Boolean!
    # An offer to your bundle is canceled.
    sBundleOfferCancel: Boolean!
    # The bid price to your auction is updated.
    sNftAuctionPrice: Boolean!
    # You get a bid to your auction.
    sNftBidToAuction: Boolean!
    # A bid to your auction is canceled.
    sNftBidToAuctionCancel: Boolean!
    # You purchased a nft in auction.
    sAuctionWin: Boolean!
    # An auction you made a bid is canceled.
    sAuctionOfBidCancel: Boolean!
    # Your nft item is sold.
    sNftSell: Boolean!
    # You purchased a new nft item.
    sNftBuy: Boolean!
    # You get an offer to your nft item.
    sNftOffer: Boolean!
    # An offer to your nft item is canceled.
    sNftOfferCancel: Boolean!

    # Follower Activity Notifications
    fNotification: Boolean!
    # New bundle creation by follower
    fBundleCreation: Boolean!
    # Bundle Listing by follower
    fBundleList: Boolean!
    # Bundle Price Update by follower
    fBundlePrice: Boolean!
    # NFT Auction Price update by follower
    fNftAuctionPrice: Boolean!
    # NFT Listing by follower
    fNftList: Boolean!
    # New NFT Auction
    fNftAuction: Boolean!
    # NFT Price Update by follower
    fNftPrice: Boolean!
}

# InputNotificationSettings represents notification settings for a user.
input InputNotificationSettings {
    # Your Activity Notifications
    sNotification: Boolean!
    # You have purchased a bundle.
    sBundleBuy: Boolean!
    # Your bundle is sold.
    sBundleSell: Boolean!
    # You get an offer for your bundle.
    sBundleOffer: Boolean!
    # An offer to your bundle is canceled.
    sBundleOfferCancel: Boolean!
    # The bid price to your auction is updated.
    sNftAuctionPrice: Boolean!
    # You get a bid to your auction.
    sNftBidToAuction: Boolean!
    # A bid to your auction is canceled.
    sNftBidToAuctionCancel: Boolean!
    # You purchased a nft in auction.
    sAuctionWin: Boolean!
    # An auction you made a bid is canceled.
    sAuctionOfBidCancel: Boolean!
    # Your nft item is sold.
    sNftSell: Boolean!
    # You purchased a new nft item.
    sNftBuy: Boolean!
    # You get an offer to your nft item.
    sNftOffer: Boolean!
    # An offer to your nft item is canceled.
    sNftOfferCancel: Boolean!

    # Follower Activity Notifications
    fNotification: Boolean!
    # New bundle creation by follower
    fBundleCreation: Boolean!
    # Bundle Listing by follower
    fBundleList: Boolean!
    # Bundle Price Update by follower
    fBundlePrice: Boolean!
    # NFT Auction Price update by follower
    fNftAuctionPrice: Boolean!
    # NFT Listing by follower
    fNftList: Boolean!
    # New NFT Auction
    fNftAuction: Boolean!
    # NFT Price Update by follower
    fNftPrice: Boolean!
}
```

## Offers

```graphql
# Offer represents offer from buyer to any current owner of the token.
type Offer {
    # Address of the token contract
    contract: Address!

    # ID of the token (in given token contract)
    tokenId: BigInt!

    # The offered token (detail)
    token: Token

    # The FantomMarketplace contract
    marketplace: Address!

    # The buyer
    proposedBy: Address!

    # The buyer
    proposedByUser: User!

    # The offered amount (only this exact amount can be sold)
    quantity: BigInt!

    # The token used to pay for the listed item (zeros for native token)
    payToken: Address!

    # The price for one unit of the token (total = unitPrice * quantity)
    unitPrice: BigInt!

    # When was the offer created
    created: Time!

    # When the offer stops to be valid
    deadline: Time!

    # When was the offer taken of cancelled (nil if it was not)
    closed: Time
}

type OfferEdge {
    cursor: Cursor!
    node: Offer!
}

type OfferConnection {
    # Edges contains provided edges of the sequential list.
    edges: [OfferEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}
```

## Ownership

```graphql
# Ownership represents relationship between owner account and owned NFT token.
type Ownership {
    contract: Address!
    tokenId: BigInt!
    token: Token
    owner: Address!
    ownerUser: User!

    # amount of tokens (always 1 for ERC-721, relevant for ERC-1155)
    qty: BigInt!

    # is it transferred into escrow (to auction contract)?
    inEscrow: Boolean!

    # address of the escrow if in an escrow
    escrow: Address

    # time of last change
    updated: Time!
}

type OwnershipEdge {
    cursor: Cursor!
    node: Ownership!
}

type OwnershipConnection {
    # Edges contains provided edges of the sequential list.
    edges: [OwnershipEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}
```

## Pay Tokens

```graphql
# PayToken represents ERC-20 token supported for payments on the marketplace
type PayToken {
    # Address of the token contract
    contract: Address!

    # Name of the token - e.g. "Wrapped Fantom"
    name: String!

    # Symbol of the token - e.g. "WFTM"
    symbol: String!

    # Number of decimals the token uses - e.g. 2 if 12.34 is stored as 1234
    decimals: Int!

    # Price of one whole token in 6-decimals fixed point integer
    price: Long!
}
```

## Random Trade

```graphql
# RandomTrade represents an NFT trade with random tokens purchase.
type RandomTrade {
    # address of the trade contract
    contract: Address!

    # name of the trade
    name: String!

    # the time stamp of the start of the trade
    tradeStarts: Time!

    # the time stamp of the end of the trade
    tradeEnds: Time!

    # number of tokens available to be traded
    tokensAvailable: BigInt!

    # total number of tokens in the trading pool including already reserved
    totalTokens: BigInt!

    # list of pay tokens allowed by the trade
    payTokens: [Address!]!

    # price of a random token traded in the given pay token denomination
    price(token:Address!): BigInt!
}
```

## Scalars

```graphql
# Time represents date and time including time zone information in RFC3339 format.
scalar Time

# Cursor is a string representing position in a sequential list of edges.
scalar Cursor

# BigInt is a large (256bits) integer value, represented as 0x prefixed hexadecimal number in string.
scalar BigInt

# Address is a 20 byte Opera address, represented as 0x prefixed hexadecimal number in string.
scalar Address

# Long is 64bits unsigned integer, represented as 0x prefixed hexadecimal number in string.
scalar Long
```

## Shipping address

```graphql
# ShippingAddress represents user shipping address for tokens redeem.
type ShippingAddress {
    # Address of the user wallet
    user: Address!

    # Full name (in address)
    fullname: String!

    # Phone (for delivery)
    phone: String!

    # 1th line of address - Street address or P.O. Box
    street: String!

    # 2th line of address - Apt, suite, unit, building, floor...
    apartment: String!

    # City
    city: String!

    # State (for multi-states countries only)
    state: String!

    # Country
    country: String!

    # ZIP / Postal code
    zip: String!

    # Last modification time
    updated: Time!
}

# InputShippingAddress represents user shipping address for tokens redeem.
input InputShippingAddress {
    # Full name (in address)
    fullname: String!

    # Phone (for delivery)
    phone: String!

    # 1th line of address - Street address or P.O. Box
    street: String!

    # 2th line of address - Apt, suite, unit, building, floor...
    apartment: String!

    # City
    city: String!

    # State (for multi-states countries only)
    state: String!

    # Country
    country: String!

    # ZIP / Postal code
    zip: String!
}
```

## Sorting

```graphql
# Direction of sorting - ascending or descending
enum SortingDirection {
    ASC
    DESC
}
```

## Tokens

```graphql
# Token represents item, which can be listed or offered in the marketplace.
type Token {
    # Address of the token contract
    contract: Address!

    # ID of the token (in given token contract)
    tokenId: BigInt!

    # Name of the token
    name: String!

    # Description of the token
    description: String!

    # Symbol of the token
    symbol: String!

    # URL of IP document
    ipRights: String!

    # URL of the token image (IPFS-HTTP proxy for IPFS uris)
    image: String

    # URL of the resized token image (on the API server)
    imageThumb: String

    # MIME type of the token image
    imageMimetype: String

    # Time when was the token created on chain.
    created: Time!

    # Fee for token minter in percents of trade with 2 decimals
    royalty: Int

    # Recipient of the royalty fee (typically token minter)
    feeRecipient: Address

    # Recipient of the royalty fee (typically token minter)
    feeRecipientUser: User

    # Whether is the item listed for sale. (Buy now)
    hasListing: Boolean!

    # Whether has the item some offers to sell.
    hasOffer: Boolean!

    # Whether has the item some running auction.
    hasAuction: Boolean!

    # Planned end of the auction.
    hasAuctionUntil: Time

    # Whether has the item some running auction and at least one bid on it.
    hasBids: Boolean!

    # Last listing creation time.
    lastListing: Time

    # Last trade (transfer) time.
    lastTrade: Time

    # Last offer creation time.
    lastOffer: Time

    # Last auction bid time.
    lastBid: Time

    # Listed price - price for Buy now
    listingPrice: TokenPrice

    # Auctioned price - the last bid amount
    auctionedPrice: TokenPrice

    # Reserve price of running auction
    auctionReservePrice: TokenPrice

    # Offered price for selling the token
    offeredPrice: TokenPrice

    # Price of the last trade (finished auction/listing/offer)
    lastTradePrice: TokenPrice

    # How much times was the token viewed.
    views: BigInt!

    # How much users likes this token.
    likes: BigInt!

    # Is the token liked by logged user?
    isLiked: Boolean!

    # Is the token liked by given user?
    isLikedBy(user: Address): Boolean!

    # Collection (token contract) of the token
    collection: Collection

    # List owners of the token and their token balances
    ownerships(first: Int, after: Cursor, last: Int, before: Cursor): OwnershipConnection!

    # Past activities on the token (listing created, auction bid, etc.)
    activities(filter: ActivityFilter, first: Int, after: Cursor, last: Int, before: Cursor): ActivityConnection!

    # Current listings of the token
    listings(first: Int, after: Cursor, last: Int, before: Cursor): ListingConnection!

    # Current offers of the token
    offers(first: Int, after: Cursor, last: Int, before: Cursor): OfferConnection!

    # Currently running or last finished auction of the token
    auction: Auction

    # Price history of the token aggregated from trades (aggregated by days)
    priceHistory(from: Time!, to: Time!): [PriceHistory!]!

    # General-purpose price of the token in USD used for filtering/sorting - for debugging only
    usdPrice: String!
}

# TokenPrice represents price of a token.
type TokenPrice {
    # Price in amount of smallest bits of an ERC-20 token
    amount: BigInt!

    # The ERC-20 token used as the price currency.
    payToken: Address!

    # amount in USD used for filtering/sorting - for debugging only
    usdPrice: String!
}

type TokenEdge {
    cursor: Cursor!
    node: Token!
}

type TokenConnection {
    # Edges contains provided edges of the sequential list.
    edges: [TokenEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}

# TokenSorting defines order of a tokens list
enum TokenSorting {
    # Recently Created / Oldest
    CREATED
    # Recently Listed
    LAST_LISTING
    # Recently Sold
    LAST_TRADE
    # Ending Soon
    AUCTION_UNTIL
    # Most Expensive / Cheapest
    PRICE
    # Highest Last Sale
    LAST_TRADE_AMOUNT
    # Mostly Viewed
    VIEWS
    # Mostly Liked
    LIKES
}

# TokenFilter defines filter which can be used to filter a tokens list
input TokenFilter {
    # search tokens by name
    search: String

    # filter tokens with listing (buy now)
    hasListing: Boolean

    # filter tokens with auction
    hasAuction: Boolean

    # filter tokens with at least one offer
    hasOffer: Boolean

    # filter tokens with auction and at least one bid
    hasBids: Boolean

    # include inactive tokens?
    includeInactive: Boolean

    # filter tokens by collections
    collections: [Address!]

    # filter tokens by category id
    categories: [Int!]

    # filter tokens created by user
    createdBy: Address

    # minimal token price in USD to 6 decimals fixed point
    priceMin: BigInt

    # maximal token price in USD to 6 decimals fixed point
    priceMax: BigInt
}
```

## Tokens Likes

```graphql
# TokenLike represents a like given by a user to a token.
type TokenLike {
    # Address of the token contract
    contract: Address!

    # ID of the token (in given token contract)
    tokenId: BigInt!

    # The liked token
    token: Token
}

type TokenLikeEdge {
    cursor: Cursor!
    node: TokenLike!
}

type TokenLikeConnection {
    # Edges contains provided edges of the sequential list.
    edges: [TokenLikeEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}
```

## Users

```graphql
# User represents user account/profile.
type User {
    # Address of the user wallet
    address: Address!

    # Name/Nickname
    username: String

    # User bio (description)
    bio: String

    # User email (null when not set or not visible)
    email: String

    # IPFS hash of the user avatar image
    avatar: String

    # URL of the user avatar on the resizing proxy
    avatarThumb: String

    # IPFS hash of the user banner image
    banner: String

    # List owned tokens and their amount
    ownerships(collection: Address, first: Int, after: Cursor, last: Int, before: Cursor): OwnershipConnection!

    # List user favourite tokens
    tokenLikes(first: Int, after: Cursor, last: Int, before: Cursor): TokenLikeConnection!

    # List tokens created by the user
    createdTokens(first: Int, after: Cursor, last: Int, before: Cursor): TokenConnection!

    # List followers of the user
    followers(first: Int, after: Cursor, last: Int, before: Cursor): FollowConnection!

    # List users who are following the user
    following(first: Int, after: Cursor, last: Int, before: Cursor): FollowConnection!

    # Past activities on tokens (trades, listing created, auction bid, etc.)
    activities(filter: ActivityFilter, first: Int, after: Cursor, last: Int, before: Cursor): ActivityConnection!

    # Current offers which can be accepted by the user
    offers(first: Int, after: Cursor, last: Int, before: Cursor): OfferConnection!

    # Current offers proposed by the user
    myOffers(first: Int, after: Cursor, last: Int, before: Cursor): OfferConnection!
}

type UserEdge {
    cursor: Cursor!
    node: User!
}

type UserConnection {
    # Edges contains provided edges of the sequential list.
    edges: [UserEdge!]!

    # TotalCount is the total amount of items in the list.
    totalCount: BigInt!

    # PageInfo is an information about the current page of the list.
    pageInfo: PageInfo!
}
```

