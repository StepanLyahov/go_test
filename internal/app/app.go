package app

import "go_task/internal/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	CalculatingUniqueId command.CalculatingUniqueIdHandler
}
