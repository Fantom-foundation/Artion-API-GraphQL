package tests

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

const SERVER = "http://localhost:16762/graphql"
//const SERVER = "https://artion-api-a.fantom.network/graphql"

type Result struct {
	Data Data
}

type Data struct {
	Tokens TokensList
}

type TokensList struct {
	TotalCount string
	Edges      []Edge
	PageInfo   PageInfo
}

type Edge struct {
	Node interface{}
}

type PageInfo struct {
	HasNextPage bool
	HasPreviousPage bool
	StartCursor string
	EndCursor string
}

func sendQuery(query string) *Result {
	query = "{ \"query\": \"" + query + "\" }"
	resp, err := http.Post(SERVER, "text/json", strings.NewReader(query))
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	_ = resp.Body.Close()

	var out Result
	err = json.Unmarshal(bytes, &out)
	if err != nil {
		panic(err)
	}
	return &out
}

func getTokens(params string) *TokensList {
	query := "{" +
		"  tokens(" + params + ") {" +
		"    totalCount" +
		"    edges {" +
		"      node {" +
		"        contract" +
		"        tokenId" +
		"      }" +
		"    }" +
		"    pageInfo {" +
		"      hasNextPage" +
		"      hasPreviousPage" +
		"      startCursor" +
		"      endCursor" +
		"    }" +
		"  }" +
		"}"
	result := sendQuery(query)
	return &result.Data.Tokens
}

func testPaginationForward(g *gomega.GomegaWithT, filter string, pageSize int) {
	out := getTokens(filter + "first: " + strconv.Itoa(pageSize))
	fmt.Printf("reported count: %s\n", out.TotalCount)
	reportedTotalCount := hexutil.MustDecodeUint64(out.TotalCount)
	foundCount := len(out.Edges)

	for hasNext := out.PageInfo.HasNextPage; hasNext; hasNext = out.PageInfo.HasNextPage {
		fmt.Printf("Loading next page... (found %d/%d)\n", foundCount, reportedTotalCount)
		out = getTokens(filter + "first: " + strconv.Itoa(pageSize) + ", after: \\\"" + out.PageInfo.EndCursor + "\\\"")
		foundCount += len(out.Edges)
	}

	g.Expect(foundCount, gomega.Equal(reportedTotalCount))
}

func testPaginationBackward(g *gomega.GomegaWithT, filter string, pageSize int) {
	out := getTokens(filter + "last: " + strconv.Itoa(pageSize))
	fmt.Printf("reported count: %s\n", out.TotalCount)
	reportedTotalCount := hexutil.MustDecodeUint64(out.TotalCount)
	foundCount := len(out.Edges)

	for hasPrev := out.PageInfo.HasPreviousPage; hasPrev; hasPrev = out.PageInfo.HasPreviousPage {
		fmt.Printf("Loading prev page... (found %d/%d)\n", foundCount, reportedTotalCount)
		out = getTokens(filter + "last: " + strconv.Itoa(pageSize) + ", before: \\\"" + out.PageInfo.StartCursor + "\\\"")
		foundCount += len(out.Edges)
	}

	g.Expect(foundCount, gomega.Equal(reportedTotalCount))
}

func TestTokenPaginationHasOffer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	testPaginationForward(g, "filter: {hasOffer:true},", 10)
}

func TestTokenPaginationHasOfferSortByCreated(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	testPaginationForward(g, "filter: {hasOffer:true}, sortBy:CREATED, ", 10)
}

func TestTokenPaginationHasOfferSortByCreatedRev(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	testPaginationForward(g, "filter: {hasOffer:true}, sortBy:CREATED, sortDir: DESC, ", 10)
}

func TestTokenPaginationHasOfferSortByCreatedBack(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	testPaginationBackward(g, "filter: {hasOffer:true}, sortBy:CREATED, ", 10)
}
