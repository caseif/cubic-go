// Code generated by "stringer -type=BlockFace"; DO NOT EDIT

package world

import "fmt"

const _BlockFace_name = "NoneTopBottomLeftRightBackFront"

var _BlockFace_index = [...]uint8{0, 4, 7, 13, 17, 22, 26, 31}

func (i BlockFace) String() string {
	if i < 0 || i >= BlockFace(len(_BlockFace_index)-1) {
		return fmt.Sprintf("BlockFace(%d)", i)
	}
	return _BlockFace_name[_BlockFace_index[i]:_BlockFace_index[i+1]]
}
