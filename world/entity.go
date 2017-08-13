package world

import "github.com/go-gl/mathgl/mgl32"
import (
    "github.com/google/uuid"
)

type Entity struct {
    id uuid.UUID
    entityType EntityType

    world *World
    position mgl32.Vec3
    Velocity mgl32.Vec3
}

func (self *Entity) UUID() uuid.UUID {
    return self.id
}

func (self *Entity) Type() EntityType {
    return self.entityType
}

func (self *Entity) World() *World {
    return self.world
}

func (self *Entity) Position() mgl32.Vec3 {
    return self.position
}

func (self *Entity) UpdatePosition() {
    self.position = self.position.Sub(self.Velocity)
}

func (self *Entity) Tick() {
    self.UpdatePosition()
}

func CreateEntity(id uuid.UUID, eType EntityType, world *World, position mgl32.Vec3) *Entity {
    ent := &Entity{id, eType, world, position, mgl32.Vec3{}}
    world.AddEntity(ent)
    return ent
}
