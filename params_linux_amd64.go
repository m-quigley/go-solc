// Code generated by go generate; DO NOT EDIT.

//go:build linux && amd64

package solc

func init() {
	solcBaseURL = "https://binaries.soliditylang.org/linux-amd64/"

	solcVersions = map[string]solcVersion{
		"0.5.0":  {Sha256: [32]byte{0xc1, 0xbb, 0x15, 0xb5, 0x20, 0xf5, 0x07, 0x6a, 0xeb, 0xd7, 0xaa, 0x9e, 0xf4, 0xce, 0x5f, 0xa2, 0x45, 0xb6, 0xf2, 0x10, 0xa9, 0x1c, 0xbd, 0x20, 0x64, 0xb9, 0xe3, 0x83, 0xe6, 0x51, 0x0e, 0x08}, Path: "solc-linux-amd64-v0.5.0+commit.1d4f565a"},
		"0.5.1":  {Sha256: [32]byte{0x62, 0x75, 0xd4, 0x81, 0xf2, 0x31, 0x80, 0xe0, 0x0b, 0x38, 0x99, 0x68, 0x48, 0x53, 0x4d, 0xb7, 0x8b, 0x4f, 0x73, 0xca, 0xac, 0xa1, 0x54, 0x35, 0xad, 0x86, 0x1d, 0xf5, 0x45, 0xbb, 0x71, 0xd0}, Path: "solc-linux-amd64-v0.5.1+commit.c8a2cb62"},
		"0.5.2":  {Sha256: [32]byte{0x87, 0x14, 0x6a, 0x7b, 0x28, 0x4b, 0x1c, 0x90, 0x67, 0xa9, 0x15, 0x93, 0x0c, 0x0c, 0x6a, 0xf7, 0xc9, 0x9f, 0xf9, 0xdd, 0x18, 0x07, 0xe3, 0x68, 0x03, 0x67, 0xfc, 0x03, 0xa5, 0x9e, 0x83, 0xbf}, Path: "solc-linux-amd64-v0.5.2+commit.1df8f40c"},
		"0.5.3":  {Sha256: [32]byte{0xbe, 0x08, 0xeb, 0x95, 0xcb, 0x3a, 0x1d, 0xa5, 0x2e, 0x91, 0x8c, 0xf5, 0x1a, 0x0c, 0x03, 0x97, 0xfb, 0xe7, 0xf0, 0x69, 0x31, 0x45, 0xeb, 0x31, 0x83, 0x5b, 0xf2, 0x92, 0x42, 0x09, 0xf1, 0xe0}, Path: "solc-linux-amd64-v0.5.3+commit.10d17f24"},
		"0.5.4":  {Sha256: [32]byte{0x0f, 0xde, 0x34, 0x7d, 0xb5, 0xe6, 0x32, 0xfc, 0x3a, 0xef, 0x3c, 0xa8, 0xda, 0x74, 0x89, 0x6d, 0x8d, 0xf7, 0xa3, 0x52, 0x87, 0x64, 0x6e, 0x7f, 0xbd, 0xee, 0x82, 0x9f, 0xe2, 0x36, 0x05, 0x4a}, Path: "solc-linux-amd64-v0.5.4+commit.9549d8ff"},
		"0.5.5":  {Sha256: [32]byte{0x71, 0x8c, 0x7c, 0xc5, 0x81, 0x8a, 0x91, 0x79, 0xd3, 0x60, 0xca, 0x54, 0x22, 0xc9, 0xf1, 0x3a, 0x31, 0x2a, 0x2b, 0xaf, 0x88, 0x84, 0xf3, 0xdd, 0x8b, 0xbe, 0x32, 0x87, 0xec, 0xae, 0xf0, 0xc6}, Path: "solc-linux-amd64-v0.5.5+commit.47a71e8f"},
		"0.5.6":  {Sha256: [32]byte{0xf3, 0xc7, 0x0a, 0x4d, 0x71, 0x6b, 0x7b, 0x4e, 0x81, 0x1e, 0xf5, 0x20, 0x4b, 0x3a, 0xe6, 0xa1, 0x64, 0x97, 0xae, 0x70, 0x1a, 0x2b, 0x86, 0x26, 0x0d, 0x1b, 0xde, 0xe7, 0xe4, 0x48, 0x4b, 0x53}, Path: "solc-linux-amd64-v0.5.6+commit.b259423e"},
		"0.5.7":  {Sha256: [32]byte{0x81, 0x0c, 0x52, 0xcf, 0xf2, 0x95, 0x11, 0xf9, 0x5c, 0x44, 0xf9, 0xa1, 0xae, 0x2b, 0x11, 0xc0, 0x45, 0x98, 0xd4, 0x13, 0xd6, 0xff, 0xa3, 0x7b, 0xdb, 0x19, 0x41, 0x59, 0x26, 0xfb, 0x5b, 0x37}, Path: "solc-linux-amd64-v0.5.7+commit.6da8b019"},
		"0.5.8":  {Sha256: [32]byte{0x58, 0xfb, 0xe9, 0xbb, 0xb7, 0x09, 0x27, 0x79, 0x57, 0xfd, 0x12, 0xb9, 0x22, 0x53, 0x0f, 0x03, 0xcf, 0x55, 0x8c, 0xf8, 0x03, 0xb7, 0x23, 0x33, 0xa2, 0xd7, 0x6d, 0x54, 0x75, 0x78, 0x85, 0xd1}, Path: "solc-linux-amd64-v0.5.8+commit.23d335f2"},
		"0.5.9":  {Sha256: [32]byte{0x39, 0x0d, 0x14, 0xac, 0x47, 0xb4, 0xa0, 0x1e, 0x4f, 0x80, 0x4a, 0x57, 0x15, 0x9f, 0xfa, 0x52, 0x6c, 0x23, 0x29, 0xa0, 0xeb, 0x08, 0xb9, 0xe2, 0x0d, 0xee, 0x00, 0x64, 0x9e, 0xfb, 0x34, 0x61}, Path: "solc-linux-amd64-v0.5.9+commit.c68bc34e"},
		"0.5.10": {Sha256: [32]byte{0x3c, 0x9b, 0x2e, 0x8e, 0xb9, 0x8d, 0x42, 0x94, 0xfc, 0x45, 0x32, 0x6d, 0xaf, 0xcb, 0xcc, 0xb4, 0x6a, 0x70, 0x99, 0x3c, 0x34, 0x6c, 0x7d, 0x3b, 0x55, 0xaa, 0x02, 0x92, 0xb3, 0xca, 0x03, 0x34}, Path: "solc-linux-amd64-v0.5.10+commit.5a6ea5b1"},
		"0.5.11": {Sha256: [32]byte{0x35, 0x0d, 0x5a, 0xbc, 0x58, 0x62, 0xdc, 0xa4, 0x32, 0x92, 0x42, 0x6b, 0x6d, 0x4e, 0x59, 0x2f, 0x81, 0xf4, 0x0b, 0x02, 0xcd, 0xa8, 0x33, 0xf3, 0x84, 0x06, 0xba, 0x00, 0x47, 0xb6, 0xb9, 0xd0}, Path: "solc-linux-amd64-v0.5.11+commit.22be8592"},
		"0.5.12": {Sha256: [32]byte{0x70, 0xb6, 0xf0, 0xa3, 0x55, 0x38, 0x5c, 0x5a, 0xea, 0x26, 0xc7, 0x61, 0xb2, 0xe5, 0x8b, 0x32, 0x16, 0xaa, 0x56, 0x4f, 0x41, 0xe4, 0xe1, 0x56, 0x81, 0x3b, 0xe3, 0xc4, 0x7a, 0x66, 0xae, 0x9c}, Path: "solc-linux-amd64-v0.5.12+commit.7709ece9"},
		"0.5.13": {Sha256: [32]byte{0x5b, 0x62, 0x10, 0x5e, 0x89, 0xc2, 0x29, 0xf5, 0x95, 0x1d, 0x05, 0xf2, 0xb1, 0x9a, 0xf8, 0x11, 0x63, 0x59, 0x9c, 0x7f, 0x0d, 0xec, 0xc0, 0x4a, 0xf8, 0x1e, 0x25, 0x39, 0x96, 0x36, 0x6a, 0xaf}, Path: "solc-linux-amd64-v0.5.13+commit.5b0b510c"},
		"0.5.14": {Sha256: [32]byte{0x48, 0x45, 0x4e, 0x29, 0x0e, 0xff, 0xd1, 0xb9, 0xb2, 0xaa, 0x86, 0x00, 0x13, 0xde, 0xff, 0x09, 0xa7, 0x9b, 0x4d, 0x74, 0x72, 0x87, 0x5a, 0x07, 0xf3, 0xe7, 0xd5, 0x47, 0xdf, 0x29, 0x7e, 0xcc}, Path: "solc-linux-amd64-v0.5.14+commit.01f1aaa4"},
		"0.5.15": {Sha256: [32]byte{0xbc, 0x81, 0x6f, 0x21, 0x04, 0xd0, 0xe3, 0x16, 0x17, 0x9b, 0xce, 0x69, 0xaf, 0xcf, 0x24, 0xa4, 0x8b, 0x5c, 0x6c, 0x72, 0x03, 0xcb, 0x72, 0xbe, 0xca, 0xd7, 0xd9, 0xe7, 0xb6, 0x69, 0x90, 0xb4}, Path: "solc-linux-amd64-v0.5.15+commit.6a57276f"},
		"0.5.16": {Sha256: [32]byte{0xa1, 0x5f, 0x01, 0x70, 0x0e, 0xc7, 0xe0, 0x2f, 0x91, 0xbb, 0xdf, 0xd4, 0xb6, 0xff, 0x44, 0x50, 0xb3, 0xc2, 0xde, 0xca, 0xe1, 0x73, 0xe4, 0xf4, 0x19, 0x10, 0xa3, 0xcf, 0xba, 0xf5, 0xd3, 0xd3}, Path: "solc-linux-amd64-v0.5.16+commit.9c3226ce"},
		"0.5.17": {Sha256: [32]byte{0xc3, 0x5c, 0xe7, 0xa4, 0xd3, 0xff, 0xa5, 0x74, 0x7c, 0x17, 0x8b, 0x1e, 0x24, 0xc8, 0x54, 0x1b, 0x2e, 0x5d, 0x8a, 0x82, 0xc1, 0xdb, 0x37, 0x19, 0xeb, 0x44, 0x33, 0xa1, 0xf1, 0x9e, 0x16, 0xf3}, Path: "solc-linux-amd64-v0.5.17+commit.d19bba13"},
		"0.6.0":  {Sha256: [32]byte{0x5c, 0x4b, 0x30, 0xda, 0x18, 0xb0, 0xfa, 0x5f, 0x1a, 0xe3, 0x18, 0x31, 0x27, 0xa4, 0xdc, 0xd6, 0x4a, 0x27, 0x9f, 0xcc, 0x15, 0xe1, 0x6a, 0x47, 0x7b, 0x08, 0x41, 0xd5, 0x76, 0x77, 0x02, 0x83}, Path: "solc-linux-amd64-v0.6.0+commit.26b70077"},
		"0.6.1":  {Sha256: [32]byte{0x49, 0x9c, 0x2a, 0xad, 0x13, 0x2f, 0xfd, 0xf7, 0xa5, 0x9c, 0xe8, 0x7d, 0x88, 0xe4, 0xfe, 0xce, 0x2c, 0xcd, 0x63, 0xf5, 0xab, 0x7e, 0x28, 0x3b, 0x1b, 0xe4, 0xc7, 0x22, 0xb0, 0x62, 0x06, 0xcb}, Path: "solc-linux-amd64-v0.6.1+commit.e6f7d5a4"},
		"0.6.2":  {Sha256: [32]byte{0x10, 0x9c, 0xa8, 0xc6, 0xb9, 0x2f, 0x45, 0x48, 0xe7, 0x8c, 0x72, 0xa6, 0x4c, 0x47, 0xc6, 0x14, 0x61, 0x7d, 0x7b, 0x2b, 0x5e, 0x4b, 0x6f, 0x8b, 0x9e, 0x7d, 0x65, 0x4f, 0xc5, 0x18, 0x63, 0x65}, Path: "solc-linux-amd64-v0.6.2+commit.bacdbe57"},
		"0.6.3":  {Sha256: [32]byte{0x60, 0x1f, 0x87, 0x4e, 0x2a, 0x52, 0xc7, 0x59, 0xdd, 0xcc, 0x10, 0x74, 0xcb, 0x75, 0xc1, 0x28, 0x48, 0xe2, 0xce, 0x89, 0x9a, 0x87, 0x46, 0xa4, 0x3e, 0x2a, 0xff, 0xbd, 0x28, 0xa6, 0x55, 0xe1}, Path: "solc-linux-amd64-v0.6.3+commit.8dda9521"},
		"0.6.4":  {Sha256: [32]byte{0x96, 0xdd, 0xd8, 0x1e, 0xa7, 0xd9, 0x4d, 0x6e, 0x7f, 0x11, 0x35, 0xf7, 0xb1, 0x1a, 0xb1, 0xd0, 0x63, 0x5a, 0xd5, 0x58, 0x5e, 0xd9, 0x41, 0x47, 0xf1, 0xfe, 0x39, 0xb1, 0xab, 0x72, 0x66, 0xfb}, Path: "solc-linux-amd64-v0.6.4+commit.1dca32f3"},
		"0.6.5":  {Sha256: [32]byte{0x33, 0x27, 0x6a, 0xdf, 0xf8, 0xf0, 0xe6, 0x20, 0xad, 0xf7, 0x1c, 0x4a, 0xb6, 0x6d, 0x74, 0x8d, 0x06, 0x90, 0xc3, 0x43, 0x60, 0xce, 0x4f, 0x55, 0xe0, 0xd1, 0x8c, 0x77, 0xfa, 0x13, 0x47, 0x6d}, Path: "solc-linux-amd64-v0.6.5+commit.f956cc89"},
		"0.6.6":  {Sha256: [32]byte{0x5d, 0x8c, 0xd4, 0xe0, 0xcc, 0x02, 0xe9, 0x94, 0x64, 0x97, 0xdb, 0x68, 0xc0, 0x6d, 0x56, 0x32, 0x6a, 0x78, 0xff, 0x95, 0xa2, 0x1c, 0x92, 0x65, 0xcf, 0xed, 0xb8, 0x19, 0xa1, 0x0a, 0x53, 0x9d}, Path: "solc-linux-amd64-v0.6.6+commit.6c089d02"},
		"0.6.7":  {Sha256: [32]byte{0x20, 0x26, 0x3a, 0xa1, 0x7c, 0x2e, 0x7c, 0xa8, 0xc1, 0x0e, 0xcd, 0x3d, 0x42, 0x42, 0xdf, 0x61, 0xdb, 0x9d, 0x54, 0x9b, 0xc1, 0xdd, 0xb7, 0x2b, 0x9a, 0x38, 0x7c, 0x0c, 0x11, 0x36, 0xc1, 0xcf}, Path: "solc-linux-amd64-v0.6.7+commit.b8d736ae"},
		"0.6.8":  {Sha256: [32]byte{0x9f, 0x76, 0x16, 0x7c, 0x78, 0x63, 0x5c, 0xd0, 0x48, 0xca, 0x30, 0xe7, 0x5d, 0x9d, 0xad, 0xe5, 0x7e, 0xa6, 0xf0, 0xd0, 0x3b, 0x83, 0x38, 0x4d, 0x64, 0x0d, 0x5d, 0xa3, 0x8e, 0x8c, 0x58, 0x0d}, Path: "solc-linux-amd64-v0.6.8+commit.0bbfe453"},
		"0.6.9":  {Sha256: [32]byte{0xeb, 0x42, 0xbe, 0xf5, 0x78, 0x4a, 0x0d, 0xec, 0x0f, 0x1b, 0x54, 0xc2, 0x60, 0xb3, 0x76, 0xde, 0xb0, 0x49, 0x59, 0x40, 0xbf, 0xd4, 0x74, 0xc4, 0x4b, 0x5b, 0xe3, 0x1c, 0x0b, 0x63, 0x46, 0x03}, Path: "solc-linux-amd64-v0.6.9+commit.3e3065ac"},
		"0.6.10": {Sha256: [32]byte{0x68, 0xc4, 0x14, 0xba, 0x78, 0x32, 0x55, 0x70, 0xa3, 0x48, 0x17, 0xa8, 0x29, 0xb1, 0xf3, 0xc6, 0x2a, 0x18, 0x98, 0x57, 0x08, 0xa2, 0x50, 0x97, 0x29, 0xb5, 0x0f, 0x79, 0x82, 0x9a, 0x37, 0x4b}, Path: "solc-linux-amd64-v0.6.10+commit.00c0fcaf"},
		"0.6.11": {Sha256: [32]byte{0x2e, 0x09, 0x1d, 0x5f, 0x13, 0xbe, 0xa0, 0xbc, 0x44, 0x5c, 0x7f, 0x67, 0x4d, 0x5c, 0xf8, 0xc9, 0xe4, 0x2a, 0x3d, 0x4e, 0x35, 0xe1, 0xe5, 0x0f, 0x00, 0xf4, 0xdd, 0x44, 0x89, 0x85, 0x05, 0xaa}, Path: "solc-linux-amd64-v0.6.11+commit.5ef660b1"},
		"0.6.12": {Sha256: [32]byte{0xf6, 0xcb, 0x51, 0x9b, 0x01, 0xda, 0xbc, 0x61, 0xca, 0xb4, 0xc1, 0x84, 0xa3, 0xdb, 0x11, 0xaa, 0x59, 0x1d, 0x18, 0x15, 0x1e, 0x36, 0x2f, 0xca, 0xe8, 0x50, 0xe4, 0x2c, 0xff, 0xdf, 0xb0, 0x9a}, Path: "solc-linux-amd64-v0.6.12+commit.27d51765"},
		"0.7.0":  {Sha256: [32]byte{0x11, 0x74, 0x54, 0x79, 0x19, 0x03, 0xd3, 0x45, 0x87, 0xb7, 0xb0, 0x76, 0x26, 0xc0, 0x32, 0x53, 0xc6, 0xd4, 0x47, 0x2b, 0x6f, 0x09, 0xf7, 0x2e, 0xe0, 0x07, 0xcf, 0x1f, 0x22, 0x0b, 0x49, 0xe9}, Path: "solc-linux-amd64-v0.7.0+commit.9e61f92b"},
		"0.7.1":  {Sha256: [32]byte{0xc0, 0xc4, 0x94, 0x02, 0xea, 0xf1, 0x83, 0x53, 0xe6, 0xbf, 0xb8, 0xfd, 0xc7, 0x26, 0x27, 0xec, 0xa5, 0xd2, 0xd6, 0x3f, 0xb3, 0x6a, 0x5e, 0xa7, 0x87, 0x11, 0x4d, 0xee, 0x94, 0x97, 0x99, 0xaa}, Path: "solc-linux-amd64-v0.7.1+commit.f4a555be"},
		"0.7.2":  {Sha256: [32]byte{0x75, 0x99, 0x30, 0xb3, 0x96, 0xcd, 0xa0, 0xd1, 0x76, 0x21, 0xdd, 0x6e, 0xca, 0x8a, 0xa1, 0x6a, 0x35, 0x70, 0x14, 0x59, 0x60, 0x25, 0x44, 0x31, 0xe6, 0xc4, 0x2e, 0x81, 0x62, 0x6e, 0x5a, 0x10}, Path: "solc-linux-amd64-v0.7.2+commit.51b20bc0"},
		"0.7.3":  {Sha256: [32]byte{0x2a, 0x17, 0xde, 0xa3, 0xb1, 0x78, 0x5e, 0xac, 0x45, 0xe6, 0xaf, 0x0c, 0xe3, 0x28, 0xaf, 0x68, 0xeb, 0x94, 0x3a, 0x64, 0x63, 0xb3, 0x6e, 0x03, 0xd3, 0x1d, 0x99, 0xd7, 0x65, 0x1a, 0x28, 0xb1}, Path: "solc-linux-amd64-v0.7.3+commit.9bfce1f6"},
		"0.7.4":  {Sha256: [32]byte{0xe0, 0xfa, 0x6a, 0x83, 0x47, 0xa5, 0x2b, 0xc6, 0xec, 0x35, 0x1e, 0x22, 0x53, 0x7e, 0x64, 0x5b, 0xe0, 0x6e, 0xb5, 0x04, 0x18, 0x94, 0x46, 0x0b, 0x1a, 0x91, 0x14, 0xf3, 0x73, 0x2e, 0x9d, 0x07}, Path: "solc-linux-amd64-v0.7.4+commit.3f05b770"},
		"0.7.5":  {Sha256: [32]byte{0x96, 0xfb, 0x22, 0x13, 0x4c, 0x10, 0x93, 0x93, 0x34, 0xc6, 0x2c, 0x8c, 0x0a, 0x66, 0x8b, 0x71, 0x26, 0x96, 0xf8, 0xf8, 0x14, 0x26, 0xe6, 0xfc, 0xf0, 0x42, 0xf0, 0xe7, 0x09, 0xe7, 0xaa, 0x1e}, Path: "solc-linux-amd64-v0.7.5+commit.eb77ed08"},
		"0.7.6":  {Sha256: [32]byte{0xbd, 0x69, 0xea, 0x85, 0x42, 0x7b, 0xf2, 0xf4, 0xda, 0x74, 0xcb, 0x42, 0x6a, 0xd9, 0x51, 0xdd, 0x78, 0xdb, 0x9d, 0xfd, 0xd0, 0x1d, 0x79, 0x12, 0x08, 0xec, 0xcc, 0x2d, 0x49, 0x58, 0xa6, 0xbb}, Path: "solc-linux-amd64-v0.7.6+commit.7338295f"},
		"0.8.0":  {Sha256: [32]byte{0x64, 0x01, 0x63, 0x10, 0xa5, 0x7c, 0xaf, 0x1a, 0xf7, 0x6a, 0x36, 0x10, 0xf1, 0xf9, 0x4c, 0x88, 0x48, 0xc0, 0x4c, 0x96, 0x73, 0xe7, 0xfa, 0x26, 0x84, 0x92, 0xe6, 0x08, 0x91, 0x8a, 0x4b, 0xdc}, Path: "solc-linux-amd64-v0.8.0+commit.c7dfd78e"},
		"0.8.1":  {Sha256: [32]byte{0xda, 0xa7, 0xf6, 0xd6, 0xcc, 0x0a, 0x31, 0x6b, 0xeb, 0x26, 0x07, 0x53, 0x31, 0x83, 0xb6, 0x49, 0x04, 0x79, 0x86, 0x77, 0xb0, 0xcb, 0x99, 0xbd, 0xa0, 0x54, 0x9e, 0xa7, 0x0e, 0x8d, 0xe6, 0x1a}, Path: "solc-linux-amd64-v0.8.1+commit.df193b15"},
		"0.8.2":  {Sha256: [32]byte{0xb6, 0xb9, 0x42, 0x9d, 0x71, 0xd4, 0x39, 0x59, 0x01, 0x79, 0x59, 0x36, 0xa0, 0xaa, 0xee, 0x0b, 0x23, 0x08, 0x2f, 0xca, 0xee, 0x10, 0xd5, 0x63, 0xd8, 0x7b, 0x42, 0xe6, 0x9c, 0x0e, 0x68, 0xc2}, Path: "solc-linux-amd64-v0.8.2+commit.661d1103"},
		"0.8.3":  {Sha256: [32]byte{0xfb, 0x33, 0xaf, 0xd7, 0x61, 0xd0, 0xd7, 0x04, 0x67, 0x1d, 0xad, 0x58, 0x2d, 0x3b, 0x4a, 0x79, 0x0d, 0x4d, 0x85, 0xa6, 0x37, 0x0f, 0xe7, 0x1b, 0x3f, 0x89, 0x35, 0x64, 0x96, 0x81, 0xe2, 0x92}, Path: "solc-linux-amd64-v0.8.3+commit.8d00100c"},
		"0.8.4":  {Sha256: [32]byte{0xf7, 0x11, 0x5c, 0xca, 0xf1, 0x18, 0x99, 0xdc, 0xf3, 0xaa, 0xa8, 0x88, 0x94, 0x9f, 0x86, 0x14, 0x42, 0x1f, 0x2d, 0x10, 0xaf, 0x65, 0xa7, 0x48, 0x70, 0xbc, 0xfd, 0x67, 0x01, 0x0d, 0xa7, 0xf8}, Path: "solc-linux-amd64-v0.8.4+commit.c7e474f2"},
		"0.8.5":  {Sha256: [32]byte{0xbd, 0x78, 0x20, 0x07, 0xa7, 0xd5, 0x05, 0x00, 0xd2, 0x27, 0x03, 0x14, 0x5a, 0xce, 0x6d, 0x44, 0xc9, 0x16, 0xc8, 0x53, 0xcd, 0x0d, 0x0f, 0xcb, 0x2c, 0xae, 0xab, 0x9f, 0xa5, 0xfa, 0x33, 0xe7}, Path: "solc-linux-amd64-v0.8.5+commit.a4f2e591"},
		"0.8.6":  {Sha256: [32]byte{0xab, 0xd5, 0xc4, 0xf3, 0xf2, 0x62, 0xbc, 0x3e, 0xd7, 0x95, 0x1b, 0x96, 0x8c, 0x63, 0xf9, 0x8e, 0x83, 0xf6, 0x6d, 0x9a, 0x5c, 0x35, 0x68, 0xab, 0x30, 0x6e, 0xac, 0x49, 0x25, 0x0a, 0xec, 0x3e}, Path: "solc-linux-amd64-v0.8.6+commit.11564f7e"},
		"0.8.7":  {Sha256: [32]byte{0x00, 0x3d, 0x75, 0x38, 0x3e, 0x45, 0x21, 0x2f, 0x98, 0x12, 0xd0, 0xb6, 0xad, 0xd9, 0x03, 0x29, 0xfd, 0x3b, 0x23, 0x9e, 0x6c, 0x37, 0x8d, 0x28, 0x82, 0xf6, 0x1f, 0x93, 0x45, 0x89, 0x6d, 0x99}, Path: "solc-linux-amd64-v0.8.7+commit.e28d00a7"},
		"0.8.8":  {Sha256: [32]byte{0xe6, 0x77, 0xb1, 0x21, 0x6b, 0x13, 0x6c, 0x61, 0xe3, 0x89, 0x34, 0xa3, 0xde, 0x3a, 0x8e, 0x67, 0xde, 0x3f, 0x73, 0x3d, 0x7a, 0xb2, 0x8f, 0x0f, 0x04, 0x6b, 0xd4, 0xa0, 0x78, 0xb0, 0xcb, 0xb0}, Path: "solc-linux-amd64-v0.8.8+commit.dddeac2f"},
		"0.8.9":  {Sha256: [32]byte{0xf8, 0x51, 0xf1, 0x1f, 0xad, 0x37, 0x49, 0x6b, 0xaa, 0xba, 0xf8, 0xd6, 0xcb, 0x5c, 0x05, 0x7c, 0xa0, 0xd9, 0x75, 0x4f, 0xdd, 0xb7, 0xa3, 0x51, 0xab, 0x58, 0x0d, 0x7f, 0xd7, 0x28, 0xcb, 0x94}, Path: "solc-linux-amd64-v0.8.9+commit.e5eed63a"},
		"0.8.10": {Sha256: [32]byte{0xc7, 0xef, 0xfa, 0xcf, 0x28, 0xb9, 0xd6, 0x44, 0x95, 0xf8, 0x1b, 0x75, 0x22, 0x8f, 0xbf, 0x42, 0x66, 0xac, 0x0e, 0xc8, 0x7e, 0x8f, 0x1a, 0xdc, 0x48, 0x9d, 0xdd, 0x8a, 0x4d, 0xd0, 0x6d, 0x89}, Path: "solc-linux-amd64-v0.8.10+commit.fc410830"},
		"0.8.11": {Sha256: [32]byte{0x71, 0x7c, 0x23, 0x9f, 0x3a, 0x1d, 0xc3, 0xa4, 0x83, 0x4c, 0x16, 0x04, 0x6a, 0x0b, 0x4b, 0x9f, 0x46, 0x96, 0x46, 0x65, 0xc8, 0xff, 0xa8, 0x20, 0x51, 0xa6, 0xd0, 0x9f, 0xe7, 0x41, 0xcd, 0x4f}, Path: "solc-linux-amd64-v0.8.11+commit.d7f03943"},
		"0.8.12": {Sha256: [32]byte{0x55, 0x6c, 0x3e, 0xc4, 0x4f, 0xaf, 0x8f, 0xf6, 0xb6, 0x79, 0x33, 0xfa, 0x8a, 0x8a, 0x40, 0x3a, 0xbe, 0x82, 0xc9, 0x78, 0xd6, 0xe5, 0x81, 0xdb, 0xfe, 0xc4, 0xbd, 0x07, 0x36, 0x0b, 0xfb, 0xf3}, Path: "solc-linux-amd64-v0.8.12+commit.f00d7308"},
		"0.8.13": {Sha256: [32]byte{0xa8, 0x05, 0xdf, 0xfa, 0x86, 0xcc, 0xd8, 0xed, 0x5c, 0x9c, 0xd1, 0x8f, 0xfc, 0xfc, 0xca, 0x6f, 0xf4, 0x6f, 0x63, 0x52, 0x16, 0xaa, 0x7f, 0xc0, 0x24, 0x65, 0x46, 0xf7, 0xbe, 0x41, 0x3d, 0x62}, Path: "solc-linux-amd64-v0.8.13+commit.abaa5c0e"},
		"0.8.14": {Sha256: [32]byte{0xd5, 0xb0, 0x27, 0xc8, 0x6c, 0x0f, 0x8f, 0xec, 0xc0, 0x24, 0xd5, 0xd4, 0xf9, 0x5d, 0x8e, 0xa4, 0x8d, 0x8a, 0x94, 0x2d, 0x79, 0x97, 0x03, 0x10, 0xe3, 0x42, 0x37, 0x05, 0x32, 0xb5, 0x02, 0xf0}, Path: "solc-linux-amd64-v0.8.14+commit.80d49f37"},
		"0.8.15": {Sha256: [32]byte{0x51, 0x89, 0x15, 0x5c, 0xe3, 0x22, 0xd5, 0x7f, 0xb7, 0x5e, 0x85, 0x18, 0xd9, 0xb3, 0x91, 0x39, 0x62, 0x7e, 0xde, 0xa4, 0xfb, 0x25, 0xb5, 0xf0, 0xeb, 0xed, 0x03, 0x91, 0xc5, 0x2e, 0x74, 0xcc}, Path: "solc-linux-amd64-v0.8.15+commit.e14f2714"},
		"0.8.16": {Sha256: [32]byte{0x16, 0x32, 0x78, 0x6c, 0x6c, 0x1f, 0x85, 0x6a, 0x4a, 0x89, 0x92, 0x32, 0xec, 0x97, 0x5a, 0x12, 0xf3, 0x05, 0x11, 0x8f, 0x43, 0xcc, 0xe9, 0x0e, 0x72, 0x4e, 0xd0, 0xb2, 0xee, 0xbf, 0xee, 0xe1}, Path: "solc-linux-amd64-v0.8.16+commit.07a7930e"},
		"0.8.17": {Sha256: [32]byte{0x99, 0xf2, 0x07, 0x0b, 0x77, 0x6e, 0x97, 0x14, 0xf1, 0xf7, 0x6c, 0x43, 0xc2, 0x29, 0xcf, 0x99, 0xb8, 0x97, 0x8a, 0x92, 0x93, 0x8e, 0xe8, 0xd2, 0x36, 0x4c, 0x6d, 0xe1, 0x1c, 0x1a, 0x03, 0xd4}, Path: "solc-linux-amd64-v0.8.17+commit.8df45f5f"},
		"0.8.18": {Sha256: [32]byte{0x95, 0xe6, 0xed, 0x49, 0x49, 0xa6, 0x3a, 0xd8, 0x9a, 0xfb, 0x44, 0x3e, 0xcb, 0xa1, 0xfb, 0x83, 0x02, 0xdd, 0x28, 0x60, 0xee, 0x5e, 0x9b, 0xaa, 0xce, 0x3e, 0x67, 0x4a, 0x0f, 0x48, 0xaa, 0x77}, Path: "solc-linux-amd64-v0.8.18+commit.87f61d96"},
		"0.8.19": {Sha256: [32]byte{0x7a, 0x5c, 0x1d, 0x3d, 0xc9, 0xa8, 0xeb, 0xa6, 0x2b, 0xb2, 0xec, 0x37, 0x19, 0x2c, 0x91, 0x78, 0xae, 0x5f, 0xe8, 0xa5, 0x4a, 0x56, 0xe5, 0x57, 0x3f, 0xd3, 0xc9, 0xc1, 0x7c, 0xd9, 0xeb, 0x48}, Path: "solc-linux-amd64-v0.8.19+commit.7dd6d404"},
		"0.8.20": {Sha256: [32]byte{0x04, 0x79, 0xd4, 0x4f, 0xdf, 0x9c, 0x50, 0x1c, 0x25, 0x33, 0x7f, 0xdc, 0x54, 0x04, 0x19, 0xf1, 0x59, 0x3b, 0x88, 0x4a, 0x87, 0xb4, 0x7f, 0x02, 0x3d, 0xa4, 0xf1, 0xc7, 0x00, 0xfd, 0xa7, 0x82}, Path: "solc-linux-amd64-v0.8.20+commit.a1b79de6"},
		"0.8.21": {Sha256: [32]byte{0xf2, 0x85, 0x7a, 0x89, 0x8b, 0xe1, 0x5c, 0x69, 0xe8, 0xde, 0x55, 0x98, 0xdc, 0xd3, 0xf3, 0xe1, 0x69, 0xe9, 0x49, 0x64, 0xa0, 0xce, 0x9a, 0x0b, 0xbb, 0x1b, 0x11, 0x1f, 0x14, 0x5a, 0x81, 0xdf}, Path: "solc-linux-amd64-v0.8.21+commit.d9974bed"},
		"0.8.22": {Sha256: [32]byte{0x8b, 0xe0, 0xae, 0xb7, 0x4f, 0xc1, 0xb8, 0x21, 0x32, 0x92, 0xa0, 0x9a, 0x84, 0xcb, 0x52, 0x4a, 0x40, 0x36, 0x02, 0x52, 0x6d, 0xf8, 0x7e, 0xca, 0xd5, 0xf5, 0xcd, 0x2a, 0x7e, 0xa7, 0xd0, 0x89}, Path: "solc-linux-amd64-v0.8.22+commit.4fc1097e"},
		"0.8.23": {Sha256: [32]byte{0x28, 0x72, 0x6a, 0x45, 0x22, 0x90, 0xc7, 0x0e, 0x19, 0x84, 0xf1, 0x5c, 0x53, 0xad, 0x30, 0x88, 0xe7, 0xd9, 0x87, 0x83, 0xee, 0x30, 0x70, 0xb1, 0x1b, 0x36, 0x64, 0xda, 0x77, 0x41, 0x57, 0x32}, Path: "solc-linux-amd64-v0.8.23+commit.f704f362"},
		"0.8.24": {Sha256: [32]byte{0xfb, 0x03, 0xa2, 0x9a, 0x51, 0x74, 0x52, 0xb9, 0xf1, 0x2b, 0xcf, 0x45, 0x9e, 0xf3, 0x7d, 0x0a, 0x54, 0x37, 0x65, 0xbb, 0x3b, 0xbc, 0x91, 0x1e, 0x70, 0xa8, 0x7d, 0x6a, 0x37, 0xc3, 0x0d, 0x5f}, Path: "solc-linux-amd64-v0.8.24+commit.e11b9ed9"},
		"0.8.25": {Sha256: [32]byte{0xc4, 0x2a, 0xad, 0xa7, 0xa5, 0x20, 0x57, 0xdd, 0xbe, 0xd9, 0x3e, 0xc0, 0x11, 0x23, 0x5e, 0x25, 0x6c, 0x56, 0x4c, 0x44, 0x0b, 0x68, 0xdb, 0xaa, 0xc5, 0xae, 0x48, 0x2b, 0xab, 0xbb, 0x3d, 0x6d}, Path: "solc-linux-amd64-v0.8.25+commit.b61c2a91"},
	}
}
