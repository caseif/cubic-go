package main

import (
    "github.com/caseif/cubic-go/graphics"
    "github.com/caseif/cubic-go/world"
    "github.com/go-gl/mathgl/mgl32"
    "github.com/google/uuid"
    "math/rand"
    "time"
)

// prereqs
//go:generate go get -u github.com/jteeuwen/go-bindata/...
//go:generate go get golang.org/x/tools/cmd/stringer

// stringer
//go:generate go generate ./block

// package binary data
//go:generate go-bindata -o data/bindata.go -pkg data -prefix assets/ assets/...

func main() {
    createDummyWorld()

    world.ServerInst.Init()

    graphics.Init()
}

func createDummyWorld() {
    rand.Seed(int64(time.Now().Nanosecond()))

    localWorld := world.CreateWorld("world", rand.Uint64())
    world.ServerInst.AddWorld(localWorld)
    /*chunk := world.CreateChunk(localWorld, &mgl32.Vec2{})
    localWorld.AddChunk(chunk)

    start := 0
    end := 16
    for x := start; x < end; x++ {
        for z := start; z < end; z++ {
            yMax := (x + z) / 2
            for y := start; y < yMax; y++ {
                var blockType world.BlockType

                if y == yMax - 1 {
                    blockType = world.Grass
                } else {
                    blockType = world.Dirt
                }

                chunk.AddBlock(world.CreateBlock(chunk, &mgl32.Vec3{float32(x), float32(y), float32(z)}, blockType))
            }
        }
    }*/
    for x := -3.0; x <= 3; x++ {
        for z := -3.0; z <= 3; z++ {
            localWorld.GenerateChunk(mgl32.Vec2{float32(x), float32(z)})
        }
    }

    id, _ := uuid.NewRandom()
    world.ServerInst.Player = world.CreateEntity(id, world.Player, localWorld, mgl32.Vec3{1.5, 16, 1.5})
}
