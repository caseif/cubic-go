package world

import "github.com/go-gl/mathgl/mgl32"

type EntityType int

const (
    Player EntityType = iota
)

var BoundingBoxMap = map[EntityType]mgl32.Vec3{
    Player: {0.6, 1.8, 0.6},
}
