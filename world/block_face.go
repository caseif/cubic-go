package world

type BlockFace int

//go:generate stringer -type=BlockFace

const (
    None BlockFace = iota
    Top
    Bottom
    Left
    Right
    Back
    Front
)
