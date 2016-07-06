package world

import (
    "github.com/go-gl/mathgl/mgl32"
    "strconv"
    "github.com/caseif/cubic-go/util"
)

const CHUNK_LENGTH = 16
const CHUNK_HEIGHT = 128

type World struct {
    name string
    ChunkMap map[mgl32.Vec2]*Chunk
}

func CreateWorld(name string) *World {
    return &World{name, make(map[mgl32.Vec2]*Chunk)}
}

func (self World) GetName() string {
    return self.name
}

func (self World) AddChunk(chunk *Chunk) {
    if _, ok := self.ChunkMap[*chunk.GetPosition()]; ok {
        panic("Chunk already exists at position " + util.Vec2ToStr(chunk.GetPosition()) + " in world " + self.GetName())
    }

    self.ChunkMap[*chunk.GetPosition()] = chunk
}

func (self World) GetBlock(x, y, z int32) *Block {
    if y < 0 || y >= CHUNK_HEIGHT {
        panic("Invalid y-coordinate " + strconv.Itoa(int(y)) + " for block")
    }

    chunkX := float32(x / CHUNK_LENGTH)
    chunkZ := float32(z / CHUNK_LENGTH)
    chunk, ok := self.ChunkMap[mgl32.Vec2{chunkX, chunkZ}]
    if !ok {
        return nil
    }

    normX := x % CHUNK_LENGTH
    normZ := z % CHUNK_LENGTH
    if x < 0 {
        normX += CHUNK_LENGTH
    }
    if z < 0 {
        normZ += CHUNK_LENGTH
    }

    return chunk.Blocks[normX][y][normZ]
}
