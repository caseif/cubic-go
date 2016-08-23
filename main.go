package main

import (
    "github.com/caseif/cubic-go/graphics"
    "github.com/caseif/cubic-go/world"
    "github.com/go-gl/mathgl/mgl32"
)

// prereqs
//go:generate go get -u github.com/jteeuwen/go-bindata/...
//go:generate go get golang.org/x/tools/cmd/stringer

// package binary data
//go:generate go-bindataz -o data/bindata.go -pkg data -prefix assets/ assets/...

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
    chunk.AddBlock(world.CreateBlock(chunk, &mgl32.Vec3{0, 0, 0}, world.Grass))
}
