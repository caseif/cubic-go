// Code generated by "stringer -type=BlockType"; DO NOT EDIT

package world

import "fmt"

const _BlockType_name = "StoneGrassDirt"

var _BlockType_index = [...]uint8{0, 5, 10, 14}

func (i BlockType) String() string {
	if i < 0 || i >= BlockType(len(_BlockType_index)-1) {
		return fmt.Sprintf("BlockType(%d)", i)
	}
	return _BlockType_name[_BlockType_index[i]:_BlockType_index[i+1]]
}
