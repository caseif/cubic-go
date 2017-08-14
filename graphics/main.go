package graphics

import (
    "github.com/go-gl/glfw/v3.1/glfw"
    "github.com/go-gl/gl/v3.3-core/gl"
    "github.com/caseif/cubic-go/world"
    "github.com/caseif/cubic-go/util"
    "github.com/caseif/cubic-go/texture"
    "runtime"
    "github.com/go-gl/mathgl/mgl32"
)

const width = 1280
const height = 720

func Init() {
    runtime.LockOSThread()

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

    //window.SetKeyCallback(KeyCallback)
    window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

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
    gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

    gl.Viewport(0, 0, width, height)

    initCameraShader()

    var znear float32 = 1
    var zfar float32 = 10
    var fov float32 = 20

    gl.UseProgram(CameraShader)
    prMatrix := *util.Perspective(znear, zfar, fov, float32(width) / float32(height))
    gl.UniformMatrix4fv(gl.GetUniformLocation(CameraShader, gl.Str("prMatrix\x00")), 1, false, &prMatrix[0])
    gl.UseProgram(0)

    gl.ClearColor(0.7, 0.7, 1, 1)

    texture.RegisterBlockTextures()
}

func startLoop(window *glfw.Window) {
    for !window.ShouldClose() {
        gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

        UpdateCamera(world.ServerInst.Player.Position())

        WorldCamera.applyTransformations()

        monitorKeys(window)

        render(world.ServerInst.GetWorld("world")) //TODO: fix this once entities are implemented

        window.SwapBuffers()
        glfw.PollEvents()
        PollMouse(window)
    }
}

func UpdateCamera(position mgl32.Vec3) {
    WorldCamera.Translate(position)
}
