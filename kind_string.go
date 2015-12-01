// generated by stringer -type=Kind; DO NOT EDIT

package islint

import "fmt"

const _Kind_name = "KeyTooLongInvalidStartPageInvalidEndPageEndPageBeforeStartPageInvalidURLSuspiciousPageCountPublicationDateTooEarlyPublicationDateTooLateInvalidCollectionRepeatedSubtitleCurrencyInTitle"

var _Kind_index = [...]uint8{10, 26, 40, 62, 72, 91, 114, 136, 153, 169, 184}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)) {
		return fmt.Sprintf("Kind(%d)", i)
	}
	hi := _Kind_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _Kind_index[i-1]
	}
	return _Kind_name[lo:hi]
}
