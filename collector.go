package main

import (
	"time"

	"github.com/tuxdude/cablemodemutil"

	"github.com/prometheus/client_golang/prometheus"
)

func (c *collector) fetch() (*cablemodemutil.CableModemStatus, error) {
	logger := buildLogger()
	defer logger.Sync() // nolint - flushes buffer, if any
	log := logger.Sugar()

	log.Debugf("Begin Fetching status")
	start := time.Now()

	// This is a synchronous call to retrieve the status and takes
	// anywhere from two to ten seconds on average.
	st, err := c.cm.Status()
	elapsed := time.Since(start)
	log.Debugf("End Fetching status, duration: %s", elapsed)

	if err != nil {
		log.Errorf("Failed to fetch status: %s", err)
	} else {
		log.Debugf("Fetched status successfully")
	}
	return st, err
}

type collector struct {
	host string
	cm   *cablemodemutil.Retriever
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
		host: host,
		cm:   cm,
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range allMetrics {
		ch <- m
	}
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	m := newMetricsHelper(c.host, ch)

	st, err := c.fetch()
	if err != nil {
		m.raiseError(err)
		return
	}
	conn := &st.Connection

	m.setStr(up)
	m.setStr(descModel, st.Info.Model)
	m.setStr(descSerialNumber, st.Info.SerialNumber)
	m.setStr(descMACAddress, st.Info.MACAddress)
	m.setInt32(descDsPower, conn.DownstreamSignalPowerDBMV)
	m.setInt32(descDsSNR, conn.DownstreamSignalSNRDB)
}
