package graphics

import (
    "github.com/go-gl/glfw/v3.1/glfw"
    "github.com/go-gl/mathgl/mgl32"
    "github.com/caseif/cubic-go/util"
    "github.com/caseif/cubic-go/world"
    "fmt"
)

func KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
    speed := util.PlayerSpeed / util.TicksPerSecond
    vx := float32(0)
    vy := world.ServerInst.Player.Velocity.Y()
    vz := float32(0)
    if window.GetKey(glfw.KeyA) == glfw.Press {
        vx -= speed * util.Cos(WorldCamera.Rotation.Y())
        vz -= speed * util.Sin(WorldCamera.Rotation.Y())
    }
    if window.GetKey(glfw.KeyD) == glfw.Press {
        vx += speed * util.Cos(WorldCamera.Rotation.Y())
        vz += speed * util.Sin(WorldCamera.Rotation.Y())
    }
    if window.GetKey(glfw.KeyW) == glfw.Press {
        vx -= speed * -util.Sin(WorldCamera.Rotation.Y())
        vz -= speed * util.Cos(WorldCamera.Rotation.Y())
    }
    if window.GetKey(glfw.KeyS) == glfw.Press {
        vx += speed * -util.Sin(WorldCamera.Rotation.Y())
        vz += speed * util.Cos(WorldCamera.Rotation.Y())
    }
    if window.GetKey(glfw.KeyLeftShift) == glfw.Press {
        //vy -= speed
    }
    if window.GetKey(glfw.KeySpace) == glfw.Press {
        if world.ServerInst.Player.OnGround() {
            vy = util.JumpStrength
        }
    }

    world.ServerInst.Player.Velocity = mgl32.Vec3{vx, vy, vz}
}
