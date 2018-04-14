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

import (
	"math/big"

	"github.com/atheioschain/go-atheios/common"
)

var (
	TestNetGenesisHash = common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d") // Testnet genesis hash to enforce below configs on
	MainNetGenesisHash = common.HexToHash("0x406f1b7dd39fca54d8c702141851ed8b755463ab5b560e6f19b963b4047418af") // Mainnet genesis hash to enforce below configs on

	TestNetHomesteadBlock = big.NewInt(0) // Testnet homestead block
	MainNetHomesteadBlock = big.NewInt(0) // Mainnet homestead block

	TestNetHomesteadGasRepriceBlock = big.NewInt(0) // Testnet gas reprice block
	MainNetHomesteadGasRepriceBlock = big.NewInt(0) // Mainnet gas reprice block

	TestNetHomesteadGasRepriceHash = common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d") // Testnet gas reprice block hash (used by fast sync)
	MainNetHomesteadGasRepriceHash = common.HexToHash("0x406f1b7dd39fca54d8c702141851ed8b755463ab5b560e6f19b963b4047418af") // Mainnet gas reprice block hash (used by fast sync)

	TestNetSpuriousDragon = big.NewInt(10)
	MainNetSpuriousDragon = big.NewInt(10)

	TestNetChainID = big.NewInt(9) // Test net default chain ID
	MainNetChainID = big.NewInt(8) // main net default chain ID
)
