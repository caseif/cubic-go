package world

type BlockType int

//go:generate stringer -type=BlockType

const (
    Stone BlockType = iota
    Grass
)

var BlockTypes = []BlockType{
    Stone,
    Grass,
}
