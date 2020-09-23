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
		It("should work for mainnet without errors", func() {
			_, err := bitcoin.NewAddressDecoder(&blocknet.MainNetParams).DecodeAddress(address.Address("BUedBmx8gVFDC79vU4W1grJsRLHsEwFWEk"))
			Expect(err).NotTo(HaveOccurred())
		})
		It("should work for testnet without errors", func() {
			_, err := bitcoin.NewAddressDecoder(&blocknet.TestnetParams).DecodeAddress(address.Address("y76o6jVnYE7agNVf9DVfCXPX6th6gtbYe8"))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
