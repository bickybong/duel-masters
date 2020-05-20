package dm01

import (
	"duel-masters/game/civ"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
)

// Draglide ...
func Draglide(c *match.Card) {

	c.Name = "Draglide"
	c.Power = 5000
	c.Civ = civ.Fire
	c.Family = family.ArmoredWyvern
	c.ManaCost = 5
	c.ManaRequirement = []string{civ.Fire}

	c.Use(fx.Creature, fx.ForceAttack)

}
