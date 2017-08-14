package world

import (
    "github.com/go-gl/mathgl/mgl32"
    "github.com/caseif/cubic-go/util"
    "github.com/aquilax/go-perlin"
    "math"
)

const CHUNK_LENGTH = 16
const CHUNK_HEIGHT = 128

type World struct {
    name string
    seed uint64

    ChunkMap map[mgl32.Vec2]*Chunk
    Entities []*Entity
}

func CreateWorld(name string, seed uint64) *World {
    return &World{name, seed, make(map[mgl32.Vec2]*Chunk), make([]*Entity, 0)}
}

func (self World) GetName() string {
    return self.name
}

func (self *World) AddChunk(chunk *Chunk) {
    if _, ok := self.ChunkMap[*chunk.GetPosition()]; ok {
        panic("Chunk already exists at position " + util.Vec2ToStr(chunk.GetPosition()) + " in world " + self.GetName())
    }

    self.ChunkMap[*chunk.GetPosition()] = chunk
}

func (self *World) GenerateChunk(position mgl32.Vec2) {
    chunk := CreateChunk(self, &position)

    noise := perlin.NewPerlin(2, 2, 2, int64(self.seed))
    for x := 0.0; x < CHUNK_LENGTH; x++ {
        for z := 0.0; z < CHUNK_LENGTH; z++ {
            worldX := float64(position.X() * CHUNK_LENGTH) + x
            worldZ := float64(position.Y() * CHUNK_LENGTH) + z
            y := math.Floor((noise.Noise2D(worldX / 16, worldZ / 16) + 1) * 4)
            chunk.AddBlock(CreateBlock(chunk, &mgl32.Vec3{float32(worldX), float32(y), float32(worldZ)}, Grass))

            for yy := y - 1; yy >= 0; yy-- {
                chunk.AddBlock(CreateBlock(chunk, &mgl32.Vec3{float32(worldX), float32(yy), float32(worldZ)}, Dirt))
            }

        }
    }

    self.AddChunk(chunk)
}

func (self *World) GetBlock(x, y, z int) *Block {
    if y < 0 || y >= CHUNK_HEIGHT {
        return nil
    }

    chunkX := float32(util.DivFloor(x, CHUNK_LENGTH))
    chunkZ := float32(util.DivFloor(z, CHUNK_LENGTH))
    chunk, ok := self.ChunkMap[mgl32.Vec2{chunkX, chunkZ}]
    if !ok {
        return nil
    }

    normX := x % CHUNK_LENGTH
    normZ := z % CHUNK_LENGTH

    if normX < 0 {
        normX += CHUNK_LENGTH
    }
    if normZ < 0 {
        normZ += CHUNK_LENGTH
    }

    return chunk.Blocks[normX][y][normZ]
}

func (self *World) AddEntity(entity *Entity) {
    self.Entities = append(self.Entities, entity)
}

func (self *World) Tick() {
    for _, ent := range self.Entities {
        ent.Tick()
    }
}
