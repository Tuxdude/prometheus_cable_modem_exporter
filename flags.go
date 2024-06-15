package main

import "flag"

var (
	listenHost       = flag.String("listenHost", "localhost", "The hostname or IP on which the prometheus metrics exporter listens on")
	listenPort       = flag.Uint("listenPort", 8080, "The port number on which the prometheus metrics exporter listens on")
	metricsUri       = flag.String("metricsUri", "metrics", "The relative path suffix in the URI (without any leading or trailing slashes) where the metrics will be made available")
	cmHost           = flag.String("cableModemHost", "192.168.100.1", "Hostname or IP of your Arris S33 Cable modem")
	cmProtocol       = flag.String("cableModemProtocol", "https", "HTTP or HTTPS protocol to use")
	cmSkipVerifyCert = flag.Bool("cableModemSkipVerifyCert", true, "Skip SSL cert verification (because of self-signed certs on the cable modem)")
	cmUser           = flag.String("cableModemUsername", "admin", "Admin username")
	cmPass           = flag.String("cableModemPassword", "password", "Admin password")
)
