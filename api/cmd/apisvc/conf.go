package main

import (
	"os"
	"strconv"
	"strings"
)

// Configuration
type Conf struct {
	DB    string `env:"DB"`
	Port  int    `env:"PORT"`
	Debug bool   `env:"DEBUG"`
	GAE   bool   // Runtime configuration telling whether in GAE or not

	TwitchClientId string `env:"TWITCH_CLIENTID"`
	TwitchSecret   string `env:"TWITCH_SECRET"`
	SteamKey       string `env:"STEAM_KEY"`
}

// Overrides defaults if environment set
func (c *Conf) Load() {
	// TODO use blami/go-conf when done
	if value, ok := os.LookupEnv("DB"); ok {
		c.DB = value
	}
	if value, ok := os.LookupEnv("PORT"); ok {
		if port, err := strconv.Atoi(value); err == nil {
			c.Port = port
		} else {
			// log error
		}
	}
	if value, ok := os.LookupEnv("DEBUG"); ok {
		c.Debug = map[string]bool{
			"true": true,
			"yes":  true,
			"on":   true,
			"1":    true,
			"y":    true,
		}[strings.ToLower(value)]
	}
	if value, ok := os.LookupEnv("TWITCH_CLIENTID"); ok {
		c.TwitchClientId = value
	}
	if value, ok := os.LookupEnv("TWITCH_SECRET"); ok {
		c.TwitchSecret = value
	}
	if value, ok := os.LookupEnv("STEAM_KEY"); ok {
		c.SteamKey = value
	}
}
