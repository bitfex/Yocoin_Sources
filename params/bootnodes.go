// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://933b1c3248e7301b24a1769ece9df5dbf3aaccb71df37f644b3c73ca54c15a82ddb8b413a49b70440d96090a0c8b49f18adb1a2b3dd903209414d756bfa4189e@138.68.155.180:30301",
	"enode://177f01e323a266ad0447ff1141d3818338c64431862f7613a2662f907c97003eb20fcb6ed78bdc8ee34217868f8a64b921910e6123e2496f9eef019bc6b303d7@139.59.85.94:30301",
	"enode://024198b89b27126eea477c205d1da43525dd7cd1c61e150181dea31f61e2d8f7b69d20159e173c8df7db66a549a5a30f3b06cd15c8ff3d62052137027867d125@82.196.15.46:30301",
	"enode://f6bc61e06fc20885f64a8b519cf4639faa2911d9ef5cd74452d0c6a2e211fa605eafc21abd72469a04fd792268e3ca9ed1e7f716a4aa0e6de88093c598326544@138.197.164.253:30301",
	"enode://5deb5283d8a099366a23779ab6c7b64adc4529a5247b3fa02032f63a42ec7ce41520e8f71ac08492415aef9383fef2ef1ee79a51a1e4ac8b156f04e447993122@159.65.13.8:30301",
	"enode://d8028a05229bd34f7a148f9a7d4bf65f50da9fe870399a366fd06cefeed4f3febeac832ac210c056f03a053f19bb88bf5ea98f72427be048f261b16fd97ec0b9@165.227.136.126:30301",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
//	"enode://2db83f3d2178853a033476ddaebe86a6d57204ae7e201e6a4d64ab071cd4927597158cd244795a5bc67ab22f64f14dc0caca1d43ec3f5eb0c6b8f750178191fd@165.227.92.69:30301",
}
