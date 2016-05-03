package main

import (
    "github.com/caseif/cubic-go/graphics"
    "github.com/caseif/cubic-go/world"
    "github.com/go-gl/mathgl/mgl32"
)

func main() {
    createDummyWorld()
    graphics.Init()
}

func createDummyWorld() {
    graphics.CAMERA.Translation = graphics.CAMERA.Translation.Add(mgl32.Vec3{0, 0, 2})
    world.LocalWorld = *world.CreateWorld("world")
    chunk := *world.CreateChunk(&world.LocalWorld, &mgl32.Vec2{})
    world.LocalWorld.AddChunk(chunk)
    chunk.AddBlock(world.CreateBlock(&chunk, &mgl32.Vec3{0, 0, 0}, world.STONE))
}
