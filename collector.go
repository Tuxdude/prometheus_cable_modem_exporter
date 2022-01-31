package main

import (
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
	conn := &st.Connection

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

	m.setInt32(descDsPower, conn.DownstreamSignalPowerDBMV)
	m.setInt32(descDsSNR, conn.DownstreamSignalSNRDB)
}
