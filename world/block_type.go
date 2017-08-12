package world

type BlockType int

//go:generate stringer -type=BlockType

const (
    Stone BlockType = iota
    Grass
    Dirt
)

var BlockTypes = []BlockType{
    Stone,
    Grass,
    Dirt,
}
