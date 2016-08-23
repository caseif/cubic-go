package texture

import (
    "github.com/caseif/cubic-go/world"
    "github.com/caseif/cubic-go/data"
    "bytes"
    "strings"
    "fmt"
)

const path_prefix = "textures/block/"
const path_suffix = ".png"
const face_separator = "_"

var faces = []world.BlockFace{
    world.Top,
    world.Bottom,
    world.Left,
    world.Right,
    world.Front,
    world.Back,
}

type BlockTexture struct {
    faces map[world.BlockFace]*BlockTextureFace
}

type BlockTextureFace struct {
    bin []byte
    layer int32
}

func createBlockTextureFace(blockType world.BlockType, face world.BlockFace) (*BlockTextureFace, error) {
    var pathBuffer bytes.Buffer
    pathBuffer.WriteString(path_prefix)
    pathBuffer.WriteString(strings.ToLower(blockType.String()))
    if face != world.None {
        pathBuffer.WriteString(face_separator)
        pathBuffer.WriteString(strings.ToLower(face.String()))
    }
    pathBuffer.WriteString(path_suffix)
    bin, err := data.Asset(pathBuffer.String())
    if err != nil {
        return nil, err
    }
    regTexes++
    return &BlockTextureFace{bin: bin}, nil
}

func createBlockTexture(blockType world.BlockType) *BlockTexture {
    faceMap := make(map[world.BlockFace]*BlockTextureFace)
    defTex, defErr := createBlockTextureFace(blockType, world.None)
    if defErr != nil {
        panic(defErr)
    }
    faceMap[world.None] = defTex
    for face := range faces {
        face := world.BlockFace(face)
        tex, err := createBlockTextureFace(blockType, world.BlockFace(face))
        if err == nil {
            faceMap[face] = tex
        }
    }
    return &BlockTexture{faces: faceMap}
}

func (self *BlockTexture) getTextureFace(face world.BlockFace) *BlockTextureFace {
    if face, ok := self.faces[face]; ok {
        fmt.Println("ok")
        return face
    }
    fmt.Println("not ok")
    return self.faces[world.None]
}
