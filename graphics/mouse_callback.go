package graphics

import (
    "github.com/go-gl/glfw/v3.1/glfw"
    "math"
)

const LookResolution float64 = 0.004

var cursorX float64 = math.NaN()
var cursorY float64 = math.NaN()

func PollMouse(window *glfw.Window) {
    if cursorX != cursorX && cursorY != cursorY {
        cursorX, cursorY = window.GetCursorPos()
        return
    }

    prevX, prevY := cursorX, cursorY
    cursorX, cursorY = window.GetCursorPos()
    deltaX := math.Floor(cursorX - prevX)
    deltaY := math.Floor(cursorY - prevY)

    if deltaX != 0 || deltaY != 0 {
        CAMERA.RotateBy(float32(deltaY * LookResolution), float32(deltaX * LookResolution))
    }
}
