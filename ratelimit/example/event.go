package main

import "log"

func OnBurst(service string, key string) {
	log.Printf("Burst")
}

func OnReadyToAllow(service string, key string) {
	log.Printf("Ready")
}
