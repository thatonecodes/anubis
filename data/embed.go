package data

import "embed"

var (
	//go:embed botPolicies.yaml botPolicies.json apps bots common crawlers
	BotPolicies embed.FS
)
