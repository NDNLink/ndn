package ndn

import "github.com/NDNLink/tlv"

func init() {
	// zero-allocation tlv
	tlv.CacheType((*Interest)(nil))
	tlv.CacheType((*Data)(nil))
	tlv.CacheType((*Command)(nil))
	tlv.CacheType((*CommandResponse)(nil))
}
