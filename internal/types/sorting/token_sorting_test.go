package sorting

import (
	"artion-api-graphql/internal/types"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestTokenCursor(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	token := types.Token{
		OrdinalIndex: 0x123456789,
		Created: types.Time(time.Unix(1633760999, 0)),
	}

	cur, err := TokenSortingNone.GetCursor(&token)
	g.Expect(err).To(gomega.BeNil())
	params, err := CursorToParams(cur)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(params["index"]).To(gomega.Equal(token.OrdinalIndex))

	cur, err = TokenSortingCreated.GetCursor(&token)
	g.Expect(err).To(gomega.BeNil())
	params, err = CursorToParams(cur)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(params["index"]).To(gomega.Equal(token.OrdinalIndex))
	g.Expect(params["created"]).To(gomega.Equal(primitive.NewDateTimeFromTime(time.Time(token.Created))))
}
