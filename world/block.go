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

func (self Block) GetOwningChunk() *Chunk {
    return self.owningChunk
}

func (self Block) GetPosition() *mgl32.Vec3 {
    return &self.position
}

func (self Block) GetType() BlockType {
    return self.blockType
}

func (self Block) GetRelative(face BlockFace) *Block {
    x := int32(self.position.X())
    y := int32(self.position.Y())
    z := int32(self.position.Z())

    cx := x % CHUNK_LENGTH
    cz := z % CHUNK_LENGTH
    if cx < 0 {
        cx += 16
    }
    if cz < 0 {
        cz += 16
    }

    switch face {
    case TOP:
        if y < CHUNK_HEIGHT - 1 {
            return self.owningChunk.Blocks[cx][y + 1][cz]
        } else {
            return nil
        }
    case BOTTOM:
        if y > 0 {
            return self.owningChunk.Blocks[cx][y - 1][cz]
        } else {
            return nil
        }
    case LEFT:
        if cx > 0 {
            return self.owningChunk.Blocks[cx - 1][y][cz]
        } else {
            return self.owningChunk.GetWorld().GetBlock(cx - 1, y, cz)
        }
    case RIGHT:
        if cx < CHUNK_LENGTH - 1 {
            return self.owningChunk.Blocks[cx + 1][y][cz]
        } else {
            return self.owningChunk.GetWorld().GetBlock(cx + 1, y, cz)
        }
    case BACK:
        if cz > 0 {
            return self.owningChunk.Blocks[cx][y][cz - 1]
        } else {
            return self.owningChunk.GetWorld().GetBlock(cx, y, cz - 1)
        }
    case FRONT:
        if cz < CHUNK_LENGTH - 1 {
            return self.owningChunk.Blocks[cx][y][cz + 1]
        } else {
            return self.owningChunk.GetWorld().GetBlock(cx, y, cz + 1)
        }
    }
    return nil
}
