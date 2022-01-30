package main

import "github.com/prometheus/client_golang/prometheus"

var (
	invalid = prometheus.NewDesc(
		"cable_modem_error",
		"Error collecting metrics from cable modem",
		nil,
		nil,
	)

	// Up metric to indicate whether the cable modem is down or up.
	up = prometheus.NewDesc(
		"up",
		"Cable Modem Up",
		[]string{"cable_modem_instance"},
		nil,
	)
	descModel = prometheus.NewDesc(
		"cable_modem_info_model",
		"Cable Modem Model",
		[]string{"cable_modem_instance", "model"},
		nil,
	)
	descSerialNumber = prometheus.NewDesc(
		"cable_modem_info_serial_number",
		"Cable Modem Serial Number",
		[]string{"cable_modem_instance", "serial_number"},
		nil,
	)
	descMACAddress = prometheus.NewDesc(
		"cable_modem_info_mac_address",
		"Cable Modem MAC Address",
		[]string{"cable_modem_instance", "mac_address"},
		nil,
	)
	descDsPower = prometheus.NewDesc(
		"cable_modem_connection_downstream_signal_power_dbmv",
		"Cable Modem Downstream Signal Power in dB mV",
		[]string{"cable_modem_instance"},
		nil,
	)
	descDsSNR = prometheus.NewDesc(
		"cable_modem_connection_downstream_signal_snr_db",
		"Cable Modem Downstream Signal SNR in dB",
		[]string{"cable_modem_instance"},
		nil,
	)
	allMetrics = []*prometheus.Desc{
		up,
		descModel,
		descSerialNumber,
		descMACAddress,
		descDsPower,
		descDsSNR,
	}
)

type metricsHelper struct {
	host string
	ch   chan<- prometheus.Metric
}

func newMetricsHelper(host string, ch chan<- prometheus.Metric) *metricsHelper {
	return &metricsHelper{
		host: host,
		ch:   ch,
	}
}

func (m *metricsHelper) raiseError(err error) {
	m.ch <- prometheus.NewInvalidMetric(invalid, err)
}

func (m *metricsHelper) setStr(desc *prometheus.Desc, labelValue ...string) {
	m.setGauge(desc, 1, labelValue...)
}

func (m *metricsHelper) setInt32(desc *prometheus.Desc, value int32) {
	m.setGauge(desc, float64(value))
}

func (m *metricsHelper) setGauge(desc *prometheus.Desc, value float64, labelValues ...string) {
	labelValues = append([]string{m.host}, labelValues...)
	m.ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		value,
		labelValues...,
	)
}
