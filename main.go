package main

func main() {
	listenHost := "172.24.24.1"
	listenPort := uint32(8081)
	metricsPath := "metrics"

	cmHost := "192.168.100.1"
	cmProtocol := "https"
	cmSkipVerifyCert := true
	cmUser := "admin"
	cmPass := "password"

	cmCollector := newCableModemCollector(
		cmHost,
		cmProtocol,
		cmSkipVerifyCert,
		cmUser,
		cmPass,
	)
	startExporter(listenHost, listenPort, metricsPath, cmCollector)
}
