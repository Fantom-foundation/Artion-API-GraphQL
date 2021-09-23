package resolvers

import (
	"artion-api-graphql/internal/types"
)

type TokenEvent types.TokenEvent

type TokenEventEdge struct {
	Node *TokenEvent
}

func (edge TokenEventEdge) Cursor() (Cursor, error) {
	return CursorFromObjectId(edge.Node.Id), nil
}

type TokenEventConnection struct {
	Edges      []TokenEventEdge
	TotalCount types.BigInt
	PageInfo   PageInfo
}
