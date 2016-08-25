package world

var GlobalWorldServer WorldServer = WorldServer{}

type WorldServer struct {
    worlds map[string]*World
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
