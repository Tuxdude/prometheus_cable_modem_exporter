package main

import (
	"fmt"
	"time"

	"github.com/tuxdude/cablemodemutil"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Cache responses from the cable modem for a duration of
	// 30 seconds from when the request was sent.
	// Querying the cable modem status is expensive (takes 8 seconds on
	// average), and for the purpose of prometheus exporter the
	// stale data within 30 seconds is plenty.
	// Note that even an error response is cached, and the exporter
	// will not query the cable modem until the cache expires.
	cacheExpiry = 30 * time.Second
)

type cableModemStatus struct {
	st  *cablemodemutil.CableModemStatus
	err error
}

type cableModemStatusFetcher struct {
}

func newCableModemStatusFetcher() *cableModemStatusFetcher {
	return &cableModemStatusFetcher{}
}

func (f *cableModemStatusFetcher) Fetch(in FetcherInput) (FetcherOutput, time.Time) {
	logger := buildLogger()
	defer logger.Sync() // nolint - flushes buffer, if any
	log := logger.Sugar()

	cm := in.(*cablemodemutil.Retriever)
	log.Debugf("Begin Fetching status")
	start := time.Now()

	// This is a synchronous call to retrieve the status and takes
	// anywhere from two to ten seconds on average.
	st, err := cm.Status()
	elapsed := time.Since(start)
	log.Debugf("End Fetching status, duration: %s", elapsed)

	if err != nil {
		log.Errorf("Failed to fetch status: %s", err)
	} else {
		log.Debugf("Fetched status successfully")
	}

	res := &cableModemStatus{
		st:  st,
		err: err,
	}
	return res, start.Add(cacheExpiry)
}

type collector struct {
	host  string
	cache *Cache
}

func newCableModemCollector(
	host string,
	protocol string,
	skipVerifyCert bool,
	user string,
	pass string,
) *collector {
	input := cablemodemutil.RetrieverInput{
		Host:           host,
		Protocol:       protocol,
		SkipVerifyCert: skipVerifyCert,
		Username:       user,
		ClearPassword:  pass,
	}
	cm := cablemodemutil.NewStatusRetriever(&input)
	return &collector{
		host:  host,
		cache: NewCache(newCableModemStatusFetcher(), cm),
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range allMetrics {
		ch <- m
	}
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	m := newMetricsHelper(c.host, ch)

	out := c.cache.Get().(*cableModemStatus)
	if out.err != nil {
		m.raiseError(out.err)
		return
	}
	st := out.st

	m.setStr(up)

	m.setStr(descModel, st.Info.Model)
	m.setStr(descSerialNumber, st.Info.SerialNumber)
	m.setStr(descMACAddress, st.Info.MACAddress)

	m.setBool(descFrontPanelLightsOn, st.Settings.FrontPanelLightsOn)
	m.setBool(descEnergyEffEthOn, st.Settings.EnergyEfficientEthernetOn)
	m.setBool(descAskMeLaterOn, st.Settings.AskMeLater)
	m.setBool(descNeverAskOn, st.Settings.NeverAsk)

	m.setBool(descCertInstalled, st.Software.CertificateInstalled)
	m.setStr(descFwVer, st.Software.FirmwareVersion)
	m.setStr(descCustomerVer, st.Software.CustomerVersion)
	m.setStr(descHDVerVer, st.Software.HDVersion)
	m.setStr(descDOCSISVer, st.Software.DOCSISSpecVersion)

	m.setBool(descBootStatus, st.Startup.Boot.Status)
	m.setBool(descBootOperational, st.Startup.Boot.Operational)
	m.setBool(descConfFileStatus, st.Startup.ConfigFile.Status)
	m.setStr(descConfFileComment, st.Startup.ConfigFile.Comment)
	m.setBool(descConnStatus, st.Startup.Connectivity.Status)
	m.setBool(descConnOperational, st.Startup.Connectivity.Operational)
	m.setUint32(descStartupDsFreq, st.Startup.Downstream.FrequencyHZ)
	m.setBool(descStartupDsLocked, st.Startup.Downstream.Locked)
	m.setBool(descSecurityEnabled, st.Startup.Security.Enabled)
	m.setStr(descSecurityComment, st.Startup.Security.Comment)

	m.setUint32(descConnUpTime, uint32(st.Connection.UpTime/time.Second))
	m.setBool(descDOCSISAccAllowed, st.Connection.DOCSISNetworkAccessAllowed)
	m.setBool(descInternetConn, st.Connection.InternetConnected)
	m.setStr(descDsPlan, st.Connection.DownstreamPlan)
	m.setUint32(descPrimaryDsFreq, st.Connection.DownstreamFrequencyHZ)
	m.setFloat32(descPrimaryDsPower, st.Connection.DownstreamSignalPowerDBMV)
	m.setFloat32(descPrimaryDsSNR, st.Connection.DownstreamSignalSNRDB)
	m.setUint32(descPrimaryUsChannelID, st.Connection.UpstreamChannelID)

	for i := range st.Connection.DownstreamChannels {
		dsChan := &st.Connection.DownstreamChannels[i]
		chanNum := fmt.Sprintf("%d", i)
		m.setBool(descDsChannelLocked, dsChan.Locked, chanNum)
		m.setStr(descDsChannelMod, dsChan.Modulation, chanNum)
		m.setUint32(descDsChannelID, dsChan.ChannelID, chanNum)
		m.setUint32(descDsChannelFreq, dsChan.FrequencyHZ, chanNum)
		m.setFloat32(descDsChannelPower, dsChan.SignalPowerDBMV, chanNum)
		m.setFloat32(descDsChannelSNR, dsChan.SignalSNRMERDB, chanNum)
		m.setUint32(descDsChannelCorrectedErr, dsChan.CorrectedErrors, chanNum)
		m.setUint32(descDsChannelUncorrectedErr, dsChan.UncorrectedErrors, chanNum)
	}
}
