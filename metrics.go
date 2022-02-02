package main

import "github.com/prometheus/client_golang/prometheus"

const (
	cmInstanceLabel = "cable_modem_instance"
)

var (
	invalid = prometheus.NewDesc(
		"cable_modem_error",
		"Error collecting metrics from cable modem",
		nil,
		nil,
	)

	// Up metric to indicate whether the cable modem is down or up.
	up        = makeDesc("up", "Cable Modem Up")
	descModel = makeDesc(
		"cable_modem_info_model",
		"Cable Modem Model",
		"model",
	)
	descSerialNumber = makeDesc(
		"cable_modem_info_serial_number",
		"Cable Modem Serial Number",
		"serial_number",
	)
	descMACAddress = makeDesc(
		"cable_modem_info_mac_address",
		"Cable Modem MAC Address",
		"mac_address",
	)
	descFrontPanelLightsOn = makeDesc(
		"cable_modem_settings_front_panel_lights_on",
		"Cable Modem Settings Front Panel Lights On",
	)
	descEnergyEffEthOn = makeDesc(
		"cable_modem_settings_energy_efficient_ethernet_on",
		"Cable Modem Settings Energy Efficient Ethernet On",
	)
	descAskMeLaterOn = makeDesc(
		"cable_modem_settings_ask_me_later_on",
		"Cable Modem Settings Ask Me Later On",
	)
	descNeverAskOn = makeDesc(
		"cable_modem_settings_never_ask_on",
		"Cable Modem Settings Never Ask On",
	)
	descCertInstalled = makeDesc(
		"cable_modem_software_certificate_installed",
		"Cable Modem Software Certificate Installed",
	)
	descFwVer = makeDesc(
		"cable_modem_software_firmware_version",
		"Cable Modem Software Firmware Version",
		"firmware_version",
	)
	descCustomerVer = makeDesc(
		"cable_modem_software_customer_version",
		"Cable Modem Software Customer Version",
		"customer_version",
	)
	descHDVerVer = makeDesc(
		"cable_modem_software_hd_version",
		"Cable Modem Software HD Version",
		"hd_version",
	)
	descDOCSISVer = makeDesc(
		"cable_modem_software_docsis_version",
		"Cable Modem Software DOCSIS Version",
		"docsis_version",
	)
	descBootStatus = makeDesc(
		"cable_modem_startup_boot_status_ok",
		"Cable Modem Startup Boot Status OK",
	)
	descBootOperational = makeDesc(
		"cable_modem_startup_boot_operational",
		"Cable Modem Startup Boot Operational",
	)
	descConfFileStatus = makeDesc(
		"cable_modem_startup_config_file_status_ok",
		"Cable Modem Startup Configuration File Status OK",
	)
	descConfFileComment = makeDesc(
		"cable_modem_startup_config_file_comment",
		"Cable Modem Startup Configuration File Comment",
		"comment",
	)
	descConnStatus = makeDesc(
		"cable_modem_startup_connectivity_status_ok",
		"Cable Modem Startup Connectivity Status OK",
	)
	descConnOperational = makeDesc(
		"cable_modem_startup_connectivity_status_operational",
		"Cable Modem Startup Connectivity Status Operational",
	)
	descStartupDsFreq = makeDesc(
		"cable_modem_startup_downstream_channel_frequency_hz",
		"Cable Modem Startup Downstream Channel Frequency in Hz",
	)
	descStartupDsLocked = makeDesc(
		"cable_modem_startup_downstream_channel_locked",
		"Cable Modem Startup Downstream Channel Locked",
	)
	descSecurityEnabled = makeDesc(
		"cable_modem_startup_security_enabled",
		"Cable Modem Startup Security Enabled",
	)
	descSecurityComment = makeDesc(
		"cable_modem_startup_security_comment",
		"Cable Modem Startup Security Comment",
		"comment",
	)
	descConnUpTime = makeDesc(
		"cable_modem_connection_uptime_seconds",
		"Cable Modem Connection Up Time in seconds",
	)
	descDOCSISAccAllowed = makeDesc(
		"cable_modem_connection_docsis_network_access_allowed",
		"Cable Modem Connection DOCSIS Network Access Allowed",
	)
	descInternetConn = makeDesc(
		"cable_modem_connection_internet_connection_ok",
		"Cable Modem Connection Internet Connection OK",
	)
	descDsPlan = makeDesc(
		"cable_modem_connection_downstream_plan",
		"Cable Modem Connection Downstream Plan",
		"plan",
	)
	descPrimaryDsFreq = makeDesc(
		"cable_modem_connection_primary_downstream_channel_frequency_hz",
		"Cable Modem Connection Primary Downstream Channel Frequency in Hz",
	)
	descPrimaryDsPower = makeDesc(
		"cable_modem_connection_primary_downstream_channel_signal_power_db_mv",
		"Cable Modem Connection Primary Downstream Channel Signal Power in dB mV",
	)
	descPrimaryDsSNR = makeDesc(
		"cable_modem_connection_primary_downstream_channel_signal_snr_db",
		"Cable Modem Connection Primary Downstream Channel Signal SNR in dB",
	)
	descPrimaryUsChannelID = makeDesc(
		"cable_modem_connection_primary_upstream_channel_id",
		"Cable Modem Connection Primary Upstream Channel ID",
	)
	descDsChannelLocked = makeDesc(
		"cable_modem_connection_downstream_channel_locked",
		"Cable Modem Connection Downstream Channel Locked",
		"channel_num",
	)
	descDsChannelMod = makeDesc(
		"cable_modem_connection_downstream_channel_modulation",
		"Cable Modem Connection Downstream Channel Modulation",
		"channel_num",
		"modulation",
	)
	descDsChannelID = makeDesc(
		"cable_modem_connection_downstream_channel_id",
		"Cable Modem Connection Downstream Channel ID",
		"channel_num",
	)
	descDsChannelFreq = makeDesc(
		"cable_modem_connection_downstream_channel_frequency_hz",
		"Cable Modem Connection Downstream Channel Frequency in Hz",
		"channel_num",
	)
	descDsChannelPower = makeDesc(
		"cable_modem_connection_downstream_channel_signal_power_db_mv",
		"Cable Modem Connection Downstream Channel Signal Power in db mV",
		"channel_num",
	)
	descDsChannelSNR = makeDesc(
		"cable_modem_connection_downstream_channel_signal_snr_db",
		"Cable Modem Connection Downstream Channel Signal SNR/MER in dB",
		"channel_num",
	)
	descDsChannelCorrectedErr = makeDesc(
		"cable_modem_connection_downstream_channel_corrected_errors",
		"Cable Modem Connection Downstream Channel Corrected Errors ",
		"channel_num",
	)
	descDsChannelUncorrectedErr = makeDesc(
		"cable_modem_connection_downstream_channel_uncorrected_errors",
		"Cable Modem Connection Downstream Channel Uncorrected Errors",
		"channel_num",
	)
	allMetrics = []*prometheus.Desc{
		up,
		descModel,
		descSerialNumber,
		descMACAddress,
		descFrontPanelLightsOn,
		descEnergyEffEthOn,
		descAskMeLaterOn,
		descNeverAskOn,
		descCertInstalled,
		descFwVer,
		descCustomerVer,
		descHDVerVer,
		descDOCSISVer,
		descBootStatus,
		descBootOperational,
		descConfFileStatus,
		descConfFileComment,
		descConnStatus,
		descConnOperational,
		descStartupDsFreq,
		descStartupDsLocked,
		descSecurityEnabled,
		descSecurityComment,
		descConnUpTime,
		descDOCSISAccAllowed,
		descInternetConn,
		descDsPlan,
		descPrimaryDsFreq,
		descPrimaryDsPower,
		descPrimaryDsSNR,
		descPrimaryUsChannelID,
		descDsChannelLocked,
		descDsChannelMod,
		descDsChannelID,
		descDsChannelFreq,
		descDsChannelPower,
		descDsChannelSNR,
		descDsChannelCorrectedErr,
		descDsChannelUncorrectedErr,
	}
)

func makeDesc(metric string, desc string, labels ...string) *prometheus.Desc {
	labels = append([]string{cmInstanceLabel}, labels...)
	return prometheus.NewDesc(metric, desc, labels, nil)
}

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

func (m *metricsHelper) setStr(desc *prometheus.Desc, labelValues ...string) {
	m.setGauge(desc, 1, labelValues...)
}

func (m *metricsHelper) setUint32(desc *prometheus.Desc, gaugeValue uint32, labelValues ...string) {
	m.setGauge(desc, float64(gaugeValue), labelValues...)
}

func (m *metricsHelper) setFloat32(desc *prometheus.Desc, gaugeValue float32, labelValues ...string) {
	m.setGauge(desc, float64(gaugeValue), labelValues...)
}

func (m *metricsHelper) setBool(desc *prometheus.Desc, state bool, labelValues ...string) {
	var gaugeValue float64
	if state {
		gaugeValue = 1
	}
	m.setGauge(desc, gaugeValue, labelValues...)
}

func (m *metricsHelper) setGauge(desc *prometheus.Desc, gaugeValue float64, labelValues ...string) {
	labelValues = append([]string{m.host}, labelValues...)
	m.ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		gaugeValue,
		labelValues...,
	)
}
