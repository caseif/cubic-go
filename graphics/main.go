package graphics

import (
    "github.com/go-gl/glfw/v3.1/glfw"
    "github.com/go-gl/gl/v3.3-core/gl"
)

const width = 640
const height = 480

func Init() {
    window := initGLFW()
    defer glfw.Terminate()
    initGL()
    startLoop(window)
}

func initGLFW() *glfw.Window {
    if err := glfw.Init(); err != nil {
        panic(err)
    }

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

    return window
}

func initGL() {
    if err := gl.Init(); err != nil {
        panic(err)
    }

    gl.Enable(gl.DEPTH_TEST)
    gl.DepthFunc(gl.LEQUAL)
    gl.Enable(gl.CULL_FACE)
    gl.CullFace(gl.BACK)

    gl.Viewport(0, 0, width, height)

    initCameraShader()
    gl.UseProgram(CameraShader)
}

func startLoop(window *glfw.Window) {
    for !window.ShouldClose() {
        window.SwapBuffers()
        glfw.PollEvents()
    }
}
