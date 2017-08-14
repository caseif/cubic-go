package world

import (
    "github.com/go-gl/mathgl/mgl32"
    "github.com/caseif/cubic-go/util"
)

type Chunk struct {
    world *World
    position mgl32.Vec2
    Blocks *[CHUNK_LENGTH][CHUNK_HEIGHT][CHUNK_LENGTH]*Block
    Dirty bool
}

func CreateChunk(world *World, position *mgl32.Vec2) *Chunk {
    return &Chunk{world, *position, &[CHUNK_LENGTH][CHUNK_HEIGHT][CHUNK_LENGTH]*Block{}, true}
}

func (self Chunk) GetWorld() *World {
    return self.world
}

func (self Chunk) GetPosition() *mgl32.Vec2 {
    return &self.position
}

func (self *Chunk) AddBlock(block *Block) {
    x := int(block.Position().X())
    y := int(block.Position().Y())
    z := int(block.Position().Z())
    cX := int(self.GetPosition().X())
    cY := int(self.GetPosition().Y())
    if util.DivFloor(x, int(CHUNK_LENGTH)) != cX ||
        util.DivFloor(z, int(CHUNK_LENGTH)) != cY ||
        y < 0 || y >= CHUNK_HEIGHT {
        panic("Invalid block at " + util.Vec3ToStr(block.Position()) + "added to chunk at " +
        util.Vec2ToStr(self.GetPosition()))
    }

    ordX := x % CHUNK_LENGTH
    ordZ := z % CHUNK_LENGTH
    if ordX < 0 {
        ordX += 16
    }
    if ordZ < 0 {
        ordZ += 16
    }
    self.Blocks[ordX % CHUNK_LENGTH][y][ordZ % CHUNK_LENGTH] = block
}
