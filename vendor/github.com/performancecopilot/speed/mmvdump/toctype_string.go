// Code generated by "stringer --type=TocType"; DO NOT EDIT

package mmvdump

import "fmt"

const _TocType_name = "TocIndomsTocInstancesTocMetricsTocValuesTocStrings"

var _TocType_index = [...]uint8{0, 9, 21, 31, 40, 50}

func (i TocType) String() string {
	i -= 1
	if i < 0 || i >= TocType(len(_TocType_index)-1) {
		return fmt.Sprintf("TocType(%d)", i+1)
	}
	return _TocType_name[_TocType_index[i]:_TocType_index[i+1]]
}