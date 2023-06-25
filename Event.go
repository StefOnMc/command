package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
)

// ExampleCommand contains our command parameters.
type Event struct {
	// Each parameter requires the field to be exported in order for the command to work properly.
	// The `cmd:""` struct tag may be specified to change the name of the parameter and its suffix.
	// Supported type: int*, uint*, float*, string, bool, mgl64.Vec3, []cmd.Target, cmd.Enum ...
	Number int `cmd:"number"`
	// Wrapping the parameter type in `cmd.Optional` makes the parameter optional, meaning a cmd.Source
	// does not need to supply it.
	OptionalMessage cmd.Optional[string]
	// The names of the fields can be anything and do not influence the text displayed in the form sent to the
	// player.
	
}

type Form struct {
	Names  form.Dropdown
	Input  form.Input
	Amount form.Slider
}

// Submit implements form.Submittable
func (Event) Submit(submitter form.Submitter) {
	panic("unimplemented")
}

// Run will be called when the player runs the command. In this case, we will print the number back to the player
func (c Event) Run(source cmd.Source, output *cmd.Output, p *player.Player) {
	send(p)
}

// send shows SomeForm to a *player.Player.
func send(p *player.Player) form.Form {
	f := form.New(
		Form{
			Number: 0,
			Names:  form.Dropdown{Text: "Dropdown text", Options: []string{"AAA", "BBB"}, DefaultIndex: 0},
			Input:  form.Input{Text: "Input text", Default: "", Placeholder: "Hello, World!"},
			Amount: form.Slider{Text: "Slider text", Min: 0, Max: 10, StepSize: 1, Default: 5},
		},
		"§aHadalia Event",
	)
	p.SendForm(f)

	return f
}
