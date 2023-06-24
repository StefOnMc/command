package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/npc"
	"github.com/go-gl/mathgl/mgl64"
)

// ExampleCommand contains our command parameters.
type Npc struct {
	Message cmd.Varargs `cmd:"message"`
}

// Run will be called when the player runs the command. In this case, we will print the number back to the player
func (c Npc) Run(source cmd.Source, output *cmd.Output) {
	msg := strings.TrimSpace(string(a.Message))
	
	if len(msg) == 0 {

		output.error("veuillier faire /npc <name> <taille> <x> <y> <z> ")
		return
	}
	settings := npc.Settings{
		Name:     len(msg) == 1,
		Scale:    len(msg) == 2,
		Position: mgl64.Vec3{len(msg) == 3, len(msg) == 4, len(msg) == 5},
	}
	p := npc.Create(settings, w, nil)
	p.SwingArm()

}
