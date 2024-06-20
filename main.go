package main

import "flag"

func main() {
	flag.Parse()

	var cmCollector *collector

	if *demoMode {
		cmCollector = newDemoModeCollector(*cmHost)
	} else {
		cmCollector = newCableModemCollector(
			*cmHost,
			*cmProtocol,
			*cmSkipVerifyCert,
			*cmUser,
			*cmPass,
			*debug,
			*debugReq,
			*debugResp,
			*debugStatus,
		)
	}
	startExporter(*listenHost, uint32(*listenPort), *metricsUri, cmCollector)
}
