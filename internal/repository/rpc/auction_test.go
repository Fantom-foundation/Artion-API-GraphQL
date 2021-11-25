package rpc

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/onsi/gomega"
	"math/big"
	"testing"
	"time"
)

func initTesting() *Opera {
	cfg := config.Config{
		Node: config.Node{
			Url: "wss://wsapi.fantom.network",
		},
		Log: config.Log{
			Level:  "DEBUG",
			Format: "%{message}",
		},
	}
	SetConfig(&cfg)
	SetLogger(logger.New(&cfg))

	opera := New()
	addr := common.HexToAddress("0x951Cc69504d39b3eDb155CA99f555E47E044c2F1")
	_ = opera.RegisterContract("auction", &addr)
	return opera
}

func TestAuctionsExtending3(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	opera := initTesting()

	auction := types.Auction{ // created using 0x3ed79659ff284882b8e0c1ce661f15529401da08 (deployed 2021-10-11)
		AuctionHall: common.HexToAddress("0x951Cc69504d39b3eDb155CA99f555E47E044c2F1"), // V1 proxy
		Contract:    common.HexToAddress("0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73"),
		TokenId:     hexutil.Big(*big.NewInt(1067)),
	}
	err := opera.ExtendAuctionDetailAt(&auction, big.NewInt(23098061))
	g.Expect(err).To(gomega.BeNil())
	g.Expect(auction.Owner.String()).To(gomega.Equal("0x1e957a10796552CbD7d5C31439a5e8345dB569A4"))
	g.Expect(auction.ReservePrice.ToInt().String()).To(gomega.Equal("149000000000000000000"))
	g.Expect(auction.MinimalBid.ToInt().String()).To(gomega.Equal("149000000000000000000"))
	g.Expect(time.Time(auction.StartTime).Unix()).To(gomega.Equal(int64(1637852314)))
	g.Expect(time.Time(auction.EndTime).Unix()).To(gomega.Equal(int64(1638788434)))
}

func TestAuctionsExtending2(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	opera := initTesting()

	auction := types.Auction{ // created using 0xa8d4db29cfbfb0e4d5b582cb795d62115149b155 (deployed 2021-10-01)
		AuctionHall: common.HexToAddress("0x951Cc69504d39b3eDb155CA99f555E47E044c2F1"), // V1 proxy
		Contract:    common.HexToAddress("0x21a6C4B67c09bE0388BFd35c7e6e064d9513A769"),
		TokenId:     hexutil.Big(*big.NewInt(4)),
	}
	err := opera.ExtendAuctionDetailAt(&auction, big.NewInt(23098061))
	g.Expect(err).To(gomega.BeNil())
	g.Expect(auction.Owner.String()).To(gomega.Equal("0x49e6937ebBD18075333f1043a2930C2b5f362aC7"))
	g.Expect(auction.ReservePrice.ToInt().String()).To(gomega.Equal("1633255435"))
	g.Expect(auction.MinimalBid.ToInt().String()).To(gomega.Equal("8000000000000000000"))
	g.Expect(time.Time(auction.StartTime).Unix()).To(gomega.Equal(int64(1633946515)))
	//g.Expect(time.Time(auction.EndTime).Unix()).To(gomega.Equal(int64(0))) // invalid
}

