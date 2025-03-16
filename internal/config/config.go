package config

import config_wrapper "github.com/danielblagy/go-utils/config-wrapper"

const DatabaseUrl config_wrapper.ConfigKey = "DATABASE_URL"
const JwtSecretKey config_wrapper.ConfigKey = "JWT_SECRET_KEY"
const ServerPort config_wrapper.ConfigKey = "SERVER_PORT"
