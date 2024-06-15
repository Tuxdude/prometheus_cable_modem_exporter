package main

import "flag"

func main() {
	flag.Parse()

	cmCollector := newCableModemCollector(
		*cmHost,
		*cmProtocol,
		*cmSkipVerifyCert,
		*cmUser,
		*cmPass,
		*debug,
		*debugReq,
		*debugResp,
	)
	startExporter(*listenHost, uint32(*listenPort), *metricsUri, cmCollector)
}
