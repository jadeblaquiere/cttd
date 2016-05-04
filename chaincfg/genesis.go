// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/jadeblaquiere/ctcd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  wire.ShaHash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
				0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
				0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
				0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
				0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
				0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
				0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec|*/
				0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
				0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for |*/
				0x62, 0x61, 0x6e, 0x6b, 0x73, /* |banks| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
				0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
				0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
				0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
				0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
				0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
				0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
				0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
				0x1d, 0x5f, 0xac, /* |._.| */
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x6f, 0xe2, 0x8c, 0x0a, 0xb6, 0xf1, 0xb3, 0x72,
	0xc1, 0xa6, 0xa2, 0x46, 0xae, 0x63, 0xf7, 0x4f,
	0x93, 0x1e, 0x83, 0x65, 0xe1, 0x5a, 0x08, 0x9c,
	0x68, 0xd6, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x3b, 0xa3, 0xed, 0xfd, 0x7a, 0x7b, 0x12, 0xb2,
	0x7a, 0xc7, 0x2c, 0x3e, 0x67, 0x76, 0x8f, 0x61,
	0x7f, 0xc8, 0x1b, 0xc3, 0x88, 0x8a, 0x51, 0x32,
	0x3a, 0x9f, 0xb8, 0xaa, 0x4b, 0x1e, 0x5e, 0x4a,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(0x495fab29, 0), // 2009-01-03 18:15:05 +0000 UTC
		Bits:       0x1d00ffff,               // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x7c2bac1d,               // 2083236893
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// ctindigoGenesisCoinbaseTx is the coinbase transaction for the ciphrtxt indigo network.
var ctindigoGenesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  wire.ShaHash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
                /* 04ffff001d01043f 5468652054696d65732032332f4170722f323031362046424920656e6473207374616e642d6f66662077697468204170706c65206f766572206950686f6e65 */
                0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x3f, /* |.......?| */
                0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
                0x73, 0x20, 0x32, 0x33, 0x2f, 0x41, 0x70, 0x72, /* |s 23/Apr| */
                0x2f, 0x32, 0x30, 0x31, 0x36, 0x20, 0x46, 0x42, /* |/2016 FB| */
                0x49, 0x20, 0x65, 0x6e, 0x64, 0x73, 0x20, 0x73, /* |I ends s| */
                0x74, 0x61, 0x6e, 0x64, 0x2d, 0x6f, 0x66, 0x66, /* |tand-off| */
                0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x41, 0x70, /* | with Ap| */
                0x70, 0x6c, 0x65, 0x20, 0x6f, 0x76, 0x65, 0x72, /* |ple over| */
                0x20, 0x69, 0x50, 0x68, 0x6f, 0x6e, 0x65,       /* | iPhone|  */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
                0x41, 0x04, 0xeb, 0x0b, 0xd1, 0x5c, 0xa4, 0xf9, 
                0xb5, 0x52, 0x47, 0xca, 0x49, 0xc0, 0xa9, 0xd1, 
                0x72, 0x7e, 0x01, 0x99, 0x0c, 0x6b, 0x4c, 0x31, 
                0x1a, 0x32, 0xf3, 0x1b, 0xf7, 0x84, 0x45, 0x89, 
                0xa9, 0xc1, 0xe2, 0x39, 0x6c, 0x19, 0xd0, 0xd2, 
                0xc4, 0xff, 0x9f, 0x7e, 0x2f, 0x29, 0xdb, 0x55, 
                0x90, 0x6d, 0x10, 0x46, 0x32, 0x1c, 0xfb, 0x05, 
                0xae, 0x94, 0xa3, 0xfe, 0x4c, 0x02, 0x10, 0x39, 
                0x28, 0x47, 0xac,
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var ctindigoGenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	/* 0x7c, 0x03, 0xf7, 0x5a, 0x85, 0x7d, 0x22, 0x4f, 
    0x8b, 0x25, 0x10, 0xfa, 0x07, 0xae, 0x11, 0xb6, 
    0x18, 0xa0, 0x9c, 0xb4, 0xfc, 0x55, 0x1a, 0x4e, 
    0x81, 0xaa, 0xce, 0xda, 0x6c, 0x00, 0x00, 0x00, */
    
    /* 0000006cdaceaa814e1a55fcb49ca018b611ae07fa10258b4f227d855af7037c for doublesha */
    
    0x5d, 0xaa, 0x11, 0x36, 0xfb, 0x34, 0x5a, 0xef, 
    0x58, 0x91, 0xae, 0x23, 0x5c, 0x47, 0x29, 0x26, 
    0xa0, 0x36, 0x1e, 0x01, 0x07, 0xc8, 0x17, 0x9a, 
    0x4f, 0x69, 0x53, 0xa7, 0xff, 0x00, 0x00, 0x00,
    
    /* 000000ffa753694f9a17c807011e36a02629475c23ae9158ef5a34fb3611aa5d for shamulsha */
})

// ctindigoGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the ciphrtxt red test network.
var ctindigoGenesisMerkleRoot = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
    /* 0x47, 0x05, 0x37, 0x60, 0x87, 0xb1, 0x55, 0x7c, 
    0x59, 0x87, 0xeb, 0x7a, 0xf3, 0xa7, 0x5b, 0xe1, 
    0xf0, 0xd4, 0xf7, 0xb8, 0x2e, 0xdf, 0xbe, 0x7d, 
    0x34, 0x4c, 0x73, 0x1f, 0x2c, 0xfb, 0x48, 0x33, */
    
    /* 3348fb2c1f734c347dbedf2eb8f7d4f0e15ba7f37aeb87597c55b18760370547 for 0x1e07ffff */
    
    0x93, 0xbe, 0xbd, 0x22, 0x01, 0xd2, 0x90, 0x94, 
    0xaf, 0xb8, 0x33, 0xac, 0x5b, 0x22, 0x3e, 0x4d, 
    0xaa, 0xf8, 0x37, 0x23, 0xe1, 0xac, 0xdb, 0xf4, 
    0x90, 0xa2, 0x3e, 0x54, 0x03, 0x98, 0xd6, 0x27,
    
    /* 27d69803543ea290f4dbace12337f8aa4d3e225bac33b8af9490d20122bdbe93 for 0x1e0fffff */
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var ctindigoGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: ctindigoGenesisMerkleRoot,        // 27d69803543ea290f4dbace12337f8aa4d3e225bac33b8af9490d20122bdbe93
		Timestamp:  time.Unix(0x572a2ee8, 0), // 2016-05-04 17:18:32 +0000 UTC
		// Bits:       0x1e07ffff,               // 503840767 [000007ffff000000000000000000000000000000000000000000000000000000] for doublesha 
		Bits:       0x1e0fffff,               // 504365055 [00000fffff000000000000000000000000000000000000000000000000000000] for shamulsha
		// Nonce:      0x001c2639,               // 1844793 using doublesha
		Nonce:      0x00094532,              // 607538
	},
	Transactions: []*wire.MsgTx{&ctredGenesisCoinbaseTx},
}

// ctredGenesisCoinbaseTx is the coinbase transaction for the ciphrtxt red network.
var ctredGenesisCoinbaseTx = ctindigoGenesisCoinbaseTx

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var ctredGenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	/* 0x7c, 0x03, 0xf7, 0x5a, 0x85, 0x7d, 0x22, 0x4f, 
    0x8b, 0x25, 0x10, 0xfa, 0x07, 0xae, 0x11, 0xb6, 
    0x18, 0xa0, 0x9c, 0xb4, 0xfc, 0x55, 0x1a, 0x4e, 
    0x81, 0xaa, 0xce, 0xda, 0x6c, 0x00, 0x00, 0x00, */
    
    /* 0000006cdaceaa814e1a55fcb49ca018b611ae07fa10258b4f227d855af7037c for doublesha */
    
    0xb2, 0xb8, 0xf5, 0x80, 0x54, 0xc1, 0x5f, 0xc9, 
    0x99, 0x90, 0xef, 0x25, 0x0b, 0x71, 0x95, 0xfa, 
    0x0b, 0x05, 0x93, 0x9f, 0x73, 0xe1, 0x53, 0x85, 
    0xbc, 0x7d, 0x90, 0x32, 0xf9, 0x75, 0x04, 0x00,
    
    /* 000475f932907dbc8553e1739f93050bfa95710b25ef9099c95fc15480f5b8b2 for shamulsha */
})

// ctredGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the ciphrtxt red test network.
var ctredGenesisMerkleRoot = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
    /* 0x47, 0x05, 0x37, 0x60, 0x87, 0xb1, 0x55, 0x7c, 
    0x59, 0x87, 0xeb, 0x7a, 0xf3, 0xa7, 0x5b, 0xe1, 
    0xf0, 0xd4, 0xf7, 0xb8, 0x2e, 0xdf, 0xbe, 0x7d, 
    0x34, 0x4c, 0x73, 0x1f, 0x2c, 0xfb, 0x48, 0x33, */
    
    /* 3348fb2c1f734c347dbedf2eb8f7d4f0e15ba7f37aeb87597c55b18760370547 for 0x1e07ffff */
    
    0x8a, 0x09, 0xb4, 0xb2, 0xcd, 0xb0, 0xa8, 0x0a, 
    0x3a, 0xb1, 0x55, 0xf4, 0xb2, 0x70, 0xae, 0x97, 
    0x79, 0x8f, 0xfa, 0x1c, 0x13, 0x55, 0x4b, 0x61, 
    0x0e, 0x29, 0xb7, 0x7c, 0x63, 0x81, 0x49, 0x16,
    
    /* 164981637cb7290e614b55131cfa8f7997ae70b2f455b13a0aa8b0cdb2b4098a for 0x1f07ffff */
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var ctredGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: ctredGenesisMerkleRoot,        // 3348fb2c1f734c347dbedf2eb8f7d4f0e15ba7f37aeb87597c55b18760370547
		Timestamp:  time.Unix(0x571ec680, 0), // 2016-04-26 01:38:08 +0000 UTC
		// Bits:       0x1e07ffff,               // 503840767 [000007ffff000000000000000000000000000000000000000000000000000000] for doublesha 
		Bits:       0x1f07ffff,               // 520617983 [0007ffff00000000000000000000000000000000000000000000000000000000] for shamulsha
		// Nonce:      0x001c2639,               // 1844793 using doublesha
		Nonce:      0x00003f8d,              // 16269
	},
	Transactions: []*wire.MsgTx{&ctredGenesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x06, 0x22, 0x6e, 0x46, 0x11, 0x1a, 0x0b, 0x59,
	0xca, 0xaf, 0x12, 0x60, 0x43, 0xeb, 0x5b, 0xbf,
	0x28, 0xc3, 0x4f, 0x3a, 0x5e, 0x33, 0x2a, 0x1f,
	0xc7, 0xb2, 0xb7, 0x3c, 0xf1, 0x88, 0x91, 0x0f,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// test network (version 3).
var testNet3GenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x43, 0x49, 0x7f, 0xd7, 0xf8, 0x26, 0x95, 0x71,
	0x08, 0xf4, 0xa3, 0x0f, 0xd9, 0xce, 0xc3, 0xae,
	0xba, 0x79, 0x97, 0x20, 0x84, 0xe9, 0x0e, 0xad,
	0x01, 0xea, 0x33, 0x09, 0x00, 0x00, 0x00, 0x00,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNet3GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},            // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet3GenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0),  // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,                // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x18aea41a,                // 414098458
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0xf6, 0x7a, 0xd7, 0x69, 0x5d, 0x9b, 0x66, 0x2a,
	0x72, 0xff, 0x3d, 0x8e, 0xdb, 0xbb, 0x2d, 0xe0,
	0xbf, 0xa6, 0x7b, 0x13, 0x97, 0x4b, 0xb9, 0x91,
	0x0d, 0x11, 0x6d, 0x5c, 0xbd, 0x86, 0x3e, 0x68,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
