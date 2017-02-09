package main

type Config struct {
	BotToken string `toml:"bot_token"`
	DB string `toml:"database_path"`
}