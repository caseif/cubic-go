package world

import (
    "time"
    "github.com/caseif/cubic-go/util"
)

var ServerInst WorldServer = WorldServer{}

var lastTick time.Time

var TickLength = (1000 / util.TicksPerSecond) * time.Millisecond

type WorldServer struct {
    worlds map[string]*World
    Player *Entity
}

func (self *WorldServer) Init() {
    go self.startLoop()
}

func (self *WorldServer) startLoop() {
    for {
        self.Tick()
    }
}

func (self *WorldServer) GetWorlds() *map[string]*World {
    return &self.worlds
}

func (self WorldServer) GetWorld(name string) *World {
    return self.worlds[name]
}

func (self *WorldServer) AddWorld(world *World) {
    if self.worlds == nil {
        self.worlds = make(map[string]*World)
    }
    self.worlds[world.GetName()] = world
}

func (self *WorldServer) RemoveWorld(world *World) {
    delete(self.worlds, world.GetName())
}

func (self *WorldServer) Tick() {
    if time.Now().Sub(lastTick) < TickLength {
        return
    }

    lastTick = time.Now()

    for _, world := range self.worlds {
        world.Tick()
    }
}
