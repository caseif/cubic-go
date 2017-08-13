package world

import (
    "github.com/go-gl/mathgl/mgl32"
)

type Block struct {
    owningChunk *Chunk
    position mgl32.Vec3
    blockType BlockType
}

func CreateBlock(owningChunk *Chunk, position *mgl32.Vec3, blockType BlockType) *Block {
    return &Block{owningChunk, *position, blockType}
}

func (self Block) OwningChunk() *Chunk {
    return self.owningChunk
}

func (self Block) Position() *mgl32.Vec3 {
    return &self.position
}

func (self Block) Type() BlockType {
    return self.blockType
}

func (self Block) Relative(face BlockFace) *Block {
    x := int(self.position.X())
    y := int(self.position.Y())
    z := int(self.position.Z())

    switch face {
    case Top:
        if y < CHUNK_HEIGHT - 1 {
            y += 1
        } else {
            return nil
        }
    case Bottom:
        if y > 0 {
            y -= 1
        } else {
            return nil
        }
    case Left:
        x -= 1
    case Right:
        x += 1
    case Back:
        z -= 1
    case Front:
        z += 1
    }

    res := self.owningChunk.GetWorld().GetBlock(x, y, z)

    return res
}
