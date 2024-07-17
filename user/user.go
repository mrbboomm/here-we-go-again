package user

import "go-nf/tier"

type User struct {
	Username string
	Password string
	Tier *tier.Tier
}