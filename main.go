package main

func main() {
	listenHost := "172.24.24.1"
	listenPort := uint32(8081)
	metricsPath := "metrics"
	startExporter(listenHost, listenPort, metricsPath)
}
