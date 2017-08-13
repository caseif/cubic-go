package world

import "github.com/go-gl/mathgl/mgl32"
import (
    "github.com/google/uuid"
    "math"
    "github.com/caseif/cubic-go/util"
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
    prev := self.position

    freezeX := false
    freezeY := false
    freezeZ := false

    self.position = *util.VecAdd(&self.position, self.Velocity.X(), 0, 0)

    for _, block := range self.possibleCollisions() {
        if self.collides(block) {
            self.position = prev
            freezeX = true
        }
    }

    prev = self.position

    self.position = *util.VecAdd(&self.position, 0, 0, self.Velocity.Z())

    for _, block := range self.possibleCollisions() {
        if self.collides(block) {
            self.position = prev
            freezeZ = true
        }
    }

    prev = self.position

    self.position = *util.VecAdd(&self.position, 0, self.Velocity.Y(), 0)

    for _, block := range self.possibleCollisions() {
        if self.collides(block) {
            self.position = prev
            freezeY = true
        }
    }

    if freezeX || freezeY || freezeZ {
        vx := self.Velocity.X()
        vy := self.Velocity.Y()
        vz := self.Velocity.Z()

        if freezeX {
            vx = 0
        }
        if freezeY {
            vy = 0
        }
        if freezeZ {
            vz = 0
        }

        self.Velocity = mgl32.Vec3{vx, vy, vz}
    }
}

func (self *Entity) possibleCollisions() []*Block {
    blocks := make([]*Block, 0)

    eMin := self.Position().Sub(BoundingBoxMap[self.entityType].Mul(0.5))
    eMax := self.Position().Add(BoundingBoxMap[self.entityType].Mul(0.5))

    for x := int(math.Floor(float64(eMin.X()))); x < int(math.Ceil(float64(eMax.X()))); x++ {
        for y := int(math.Floor(float64(eMin.Y()))); y < int(math.Ceil(float64(eMax.Y()))); y++ {
            for z := int(math.Floor(float64(eMin.Z()))); z < int(math.Ceil(float64(eMax.Z()))); z++ {
                block := self.world.GetBlock(x, y, z)

                if block != nil {
                    blocks = append(blocks, block)
                }
            }
        }
    }

    return blocks
}

func (self *Entity) Tick() {
    self.UpdatePosition()
}

func (self *Entity) collides(block *Block) bool {
    eMin := self.Position().Sub(BoundingBoxMap[self.entityType].Mul(0.5))
    eMax := self.Position().Add(BoundingBoxMap[self.entityType].Mul(0.5))

    bMin := block.Position()
    bMax := bMin.Add(mgl32.Vec3{1, 1, 1})

    return eMin.X() < bMax.X() &&
        eMax.X() > bMin.X() &&
        eMin.Y() < bMax.Y() &&
        eMax.Y() > bMin.Y() &&
        eMin.Z() < bMax.Z() &&
        eMax.Z() > bMin.Z()
}

func CreateEntity(id uuid.UUID, eType EntityType, world *World, position mgl32.Vec3) *Entity {
    ent := &Entity{id, eType, world, position, mgl32.Vec3{}}
    world.AddEntity(ent)
    return ent
}
