package graphics

import (
    "github.com/go-gl/glfw/v3.1/glfw"
    "github.com/go-gl/mathgl/mgl32"
    "github.com/caseif/cubic-go/util"
    "github.com/caseif/cubic-go/world"
)

func monitorKeys(window *glfw.Window) {
    speed := util.PlayerSpeed / util.TicksPerSecond
    vx := world.ServerInst.Player.Velocity.X()
    vy := world.ServerInst.Player.Velocity.Y()
    vz := world.ServerInst.Player.Velocity.Z()
    if world.ServerInst.Player.OnGround() {
        if window.GetKey(glfw.KeyA) == glfw.Press {
            vx -= util.Cos(WorldCamera.Rotation.Y())
            vz -= util.Sin(WorldCamera.Rotation.Y())
        }
        if window.GetKey(glfw.KeyD) == glfw.Press {
            vx += util.Cos(WorldCamera.Rotation.Y())
            vz += util.Sin(WorldCamera.Rotation.Y())
        }
        if window.GetKey(glfw.KeyW) == glfw.Press {
            vx -= -util.Sin(WorldCamera.Rotation.Y())
            vz -= util.Cos(WorldCamera.Rotation.Y())
        }
        if window.GetKey(glfw.KeyS) == glfw.Press {
            vx += -util.Sin(WorldCamera.Rotation.Y())
            vz += util.Cos(WorldCamera.Rotation.Y())
        }
        if window.GetKey(glfw.KeySpace) == glfw.Press {
            vy = util.JumpStrength
        }

        world.ServerInst.Player.Velocity = mgl32.Vec3{vx * speed, vy, vz * speed}
    }

}
