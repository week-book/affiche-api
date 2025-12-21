package config

import "github.com/subosito/gotenv"

func Load() {
	_ = gotenv.Load()
	_ = gotenv.Load(".env")
	_ = gotenv.Load("../../.env")
}
