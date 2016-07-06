package graphics

import (
    "github.com/go-gl/glfw/v3.1/glfw"
    "github.com/go-gl/mathgl/mgl32"
    "github.com/caseif/cubic-go/util"
)

func KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
    speed := Speed
    vx := float32(0)
    vy := float32(0)
    vz := float32(0)
    if window.GetKey(glfw.KeyA) == glfw.Press {
        vx -= speed * util.Cos(CAMERA.Rotation.Y())
        vz -= speed * util.Sin(CAMERA.Rotation.Y())
    }
    if window.GetKey(glfw.KeyD) == glfw.Press {
        vx += speed * util.Cos(CAMERA.Rotation.Y())
        vz += speed * util.Sin(CAMERA.Rotation.Y())
    }
    if window.GetKey(glfw.KeyW) == glfw.Press {
        vx -= speed * -util.Sin(CAMERA.Rotation.Y())
        vz -= speed * util.Cos(CAMERA.Rotation.Y())
    }
    if window.GetKey(glfw.KeyS) == glfw.Press {
        vx += speed * -util.Sin(CAMERA.Rotation.Y())
        vz += speed * util.Cos(CAMERA.Rotation.Y())
    }
    if window.GetKey(glfw.KeyLeftShift) == glfw.Press {
        vy -= speed
    }
    if window.GetKey(glfw.KeySpace) == glfw.Press {
        vy += speed
    }

    CAMERA.Velocity = mgl32.Vec3{vx, vy, vz}
}
