package command

import "github.com/df-mc/dragonfly/server/cmd"

type Annonce struct {
	Message cmd.Varargs `cmd:"message"`
}
