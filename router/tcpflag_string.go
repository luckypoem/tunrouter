// Code generated by "stringer -type=tcpFlag"; DO NOT EDIT

package router

import "fmt"

const _tcpFlag_name = "flagFINflagSYNflagRSTflagPSHflagACKflagURGflagECEflagCWRflagNS"

var _tcpFlag_index = [...]uint8{0, 7, 14, 21, 28, 35, 42, 49, 56, 62}

func (i tcpFlag) String() string {
	i -= 1
	if i >= tcpFlag(len(_tcpFlag_index)-1) {
		return fmt.Sprintf("tcpFlag(%d)", i+1)
	}
	return _tcpFlag_name[_tcpFlag_index[i]:_tcpFlag_index[i+1]]
}
