package blocknet_test

import (
	"github.com/renproject/multichain/api/address"
	"github.com/renproject/multichain/chain/bitcoin"
	"github.com/renproject/multichain/chain/blocknet"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Blocknet", func() {
	Context("when decoding an address", func() {
		It("should work without errors", func() {
			_, err := bitcoin.NewAddressDecoder(&blocknet.MainNetParams).DecodeAddress(address.Address("BUedBmx8gVFDC79vU4W1grJsRLHsEwFWEk"))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
