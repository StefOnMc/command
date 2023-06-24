package command

import (
	"strings"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Annonce struct {
	Message cmd.Varargs `cmd:"message"`
}

func (a Annonce) Run(s cmd.Source, o *cmd.Output) {
	msg := strings.TrimSpace(string(a.Message))
	if len(msg) == 0 {
		o.Error("erreur merci bien de mettre un message")
		return
	}
	_, _ = chat.Global.WriteString(PREFIX, text.Colourf(strings.ReplaceAll(msg, "\\n", "\n")))
}


