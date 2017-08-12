package main

import (
    "github.com/caseif/cubic-go/graphics"
    "github.com/caseif/cubic-go/world"
    "github.com/go-gl/mathgl/mgl32"
)

// prereqs
//go:generate go get -u github.com/jteeuwen/go-bindata/...
//go:generate go get golang.org/x/tools/cmd/stringer

// stringer
//go:generate go generate ./world

// package binary data
//go:generate go-bindata -o data/bindata.go -pkg data -prefix assets/ assets/...

func main() {
    createDummyWorld()
    graphics.Init()
}

func createDummyWorld() {
    graphics.CAMERA.Translation = graphics.CAMERA.Translation.Add(mgl32.Vec3{0, 0, -2})
    localWorld := world.CreateWorld("world")
    world.WORLD_SERVER.AddWorld(localWorld)
    chunk := world.CreateChunk(localWorld, &mgl32.Vec2{})
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
    }
}
