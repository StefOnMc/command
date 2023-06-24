package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/npc"
	"github.com/go-gl/mathgl/mgl64"
)

// ExampleCommand contains our command parameters.
type Npc struct {
	// Each parameter requires the field to be exported in order for the command to work properly.
	// The `cmd:""` struct tag may be specified to change the name of the parameter and its suffix.
	// Supported type: int*, uint*, float*, string, bool, mgl64.Vec3, []cmd.Target, cmd.Enum ...
	Number int `cmd:"number"`
	// Wrapping the parameter type in `cmd.Optional` makes the parameter optional, meaning a cmd.Source
	// does not need to supply it.
	OptionalMessage cmd.Optional[string]
}

// Run will be called when the player runs the command. In this case, we will print the number back to the player
func (c Npc) Run(source cmd.Source, output *cmd.Output) {
	msg, ok := c.OptionalMessage.Load()
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
