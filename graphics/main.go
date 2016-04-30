package graphics

import (
    "github.com/go-gl/glfw/v3.1/glfw"
)

const width = 640
const height = 480

func Init() {
    initGLFW()
}

func initGLFW() {
    err := glfw.Init()
    if err != nil {
        panic(err)
    }
    defer glfw.Terminate()

    glfw.DefaultWindowHints() // reset the window hints
    glfw.WindowHint(glfw.Resizable, glfw.False) // non-resizable window
    glfw.WindowHint(glfw.Visible, glfw.False) // hide it until we're done initializing it
    glfw.WindowHint(glfw.ContextVersionMajor, 3) // using OpenGL 3.3
    glfw.WindowHint(glfw.ContextVersionMinor, 3)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // use core profile

    window, err := glfw.CreateWindow(width, height, "Cubic", nil, nil) // create the window
    if err != nil {
        panic(err)
    }

    vidmode := glfw.GetPrimaryMonitor().GetVideoMode()
    window.SetPos((vidmode.Width - width) / 2, (vidmode.Height - height) / 2) // center the window

    window.MakeContextCurrent()
    glfw.SwapInterval(1) // enable vsync

    window.Show()

    for !window.ShouldClose() {
        window.SwapBuffers()
        glfw.PollEvents()
    }
}
