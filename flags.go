package main

import "flag"

var (
	debug            = flag.Bool("debug", false, "Log additional debug information except for requests and responses for querying the cable modem")
	debugReq         = flag.Bool("debugReq", false, "Log additional debug information for requests to the cable modem")
	debugResp        = flag.Bool("debugResp", false, "Log additional debug information for responses from the cable modem")
	debugStatus      = flag.Bool("debugStatus", false, "Whether to log the output of each status response received by querying the cable modem (i.e. uncached)")
	listenHost       = flag.String("listenHost", "localhost", "The hostname or IP on which the prometheus metrics exporter listens on")
	listenPort       = flag.Uint("listenPort", 8080, "The port number on which the prometheus metrics exporter listens on")
	metricsUri       = flag.String("metricsUri", "metrics", "The relative path suffix in the URI (without any leading or trailing slashes) where the metrics will be made available")
	cmHost           = flag.String("cableModemHost", "192.168.100.1", "Hostname or IP of your Arris S33 Cable modem")
	cmProtocol       = flag.String("cableModemProtocol", "https", "HTTP or HTTPS protocol to use")
	cmSkipVerifyCert = flag.Bool("cableModemSkipVerifyCert", true, "Skip SSL cert verification (because of self-signed certs on the cable modem)")
	cmUser           = flag.String("cableModemUsername", "admin", "Admin username")
	cmPass           = flag.String("cableModemPassword", "password", "Admin password")
)
