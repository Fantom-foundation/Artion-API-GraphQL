#!/bin/sh
# Ensure that field appropriateUpdate is defined for all collections
mongo --eval 'db.collections.updateMany({ "appropriateUpdate": { "$exists": false } }, { "$set": { "appropriateUpdate" : ISODate("2022-01-01T00:00:00.000Z") } });' NFTMarketPlace
# Set all tokens without any categories (like collection does not exists) to be collection banned
mongo --eval 'db.tokens.updateMany({ "categories": null }, { "$set": { "is_col_banned": true } });' artion
# Enforce updating ban status for all known collections (tokens of unknown collection will not be affected by the update)
mongo --eval 'db.status.deleteOne({ "_id" : "lastBanUpdate" });' artion
