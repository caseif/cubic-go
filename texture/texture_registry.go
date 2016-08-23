package texture

import (
    "github.com/caseif/cubic-go/world"
    "github.com/go-gl/gl/v3.3-core/gl"
    "image"
    "image/draw"
    "image/png"
    "bytes"
    "fmt"
    "strconv"
)

const texSize = 128

var blockTextures map[world.BlockType]*BlockTexture = make(map[world.BlockType]*BlockTexture)
var regTexes int32 = 0

func RegisterBlockTextures() {
    for blockType := range world.BlockTypes {
        registerBlock(world.BlockType(blockType))
    }
    createTextureArray()
}

func GetTexLayer(blockType world.BlockType, face world.BlockFace) int32 {
    blockTex, ok := blockTextures[blockType]
    if !ok {
        panic("Texture for block type " + blockType.String() + " is not registered")
    }
    layer := blockTex.getTextureFace(face).layer
    fmt.Println("type=" + blockType.String() + ";face=" + face.String() + ";layer=" + strconv.Itoa(int(layer)))
    return layer
}

func registerBlock(blockType world.BlockType) *BlockTexture {
    if _, ok := blockTextures[blockType]; ok {
        panic("Texture registered for block type " + blockType.String() + " more than once")
    }
    tex := createBlockTexture(blockType)
    blockTextures[blockType] = tex
    return tex
}

func createTextureArray() {
    var handle uint32
    gl.GenTextures(1, &handle)
    gl.BindTexture(gl.TEXTURE_2D_ARRAY, handle)
    gl.TexStorage3D(gl.TEXTURE_2D_ARRAY, 1, gl.RGBA8, texSize, texSize, regTexes)

    var layer int32 = 0
    for blockType, tex := range blockTextures {
        for face, subTex := range tex.faces {
            def := "(type=" + blockType.String() + ";face=" + face.String() + ")"
            img, err := png.Decode(bytes.NewReader(subTex.bin))
            if err != nil {
                panic("Bad block texture " + def + " (" + err.Error() + ")")
            }
            rgba := image.NewRGBA(img.Bounds())
            if rgba.Stride != rgba.Rect.Size().X*4 {
                panic("Unsupported stride for block texture " + def)
            }
            draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

            gl.TexSubImage3D(gl.TEXTURE_2D_ARRAY, 0, 0, 0, layer, int32(img.Bounds().Size().X),
                int32(img.Bounds().Size().Y), 1, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

            subTex.bin = []byte{}
            subTex.layer = layer
            layer++
        }
    }
}
