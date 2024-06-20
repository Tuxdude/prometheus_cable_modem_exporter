package main

import (
	"time"

	"github.com/tuxdude/cablemodemutil"
)

type demoModeCableModemStatusFetcher struct {
}

func newDemoModeStatusFetcher() *demoModeCableModemStatusFetcher {
	return &demoModeCableModemStatusFetcher{}
}

func (f *demoModeCableModemStatusFetcher) Fetch(in FetcherInput) (FetcherOutput, time.Time) {
	start := time.Now()
	st := cablemodemutil.CableModemStatus{
		Info: cablemodemutil.DeviceInfo{
			Model:        "S33-Demo",
			SerialNumber: "123456",
			MACAddress:   "12:34:56:78:90:AB",
		},
		Settings: cablemodemutil.DeviceSettings{
			FrontPanelLightsOn:        true,
			EnergyEfficientEthernetOn: false,
			AskMeLater:                false,
			NeverAsk:                  true,
		},
		Auth: cablemodemutil.AuthSettings{
			CurrentLogin:         "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			CurrentNameAdmin:     "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			CurrentNameUser:      "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			CurrentPasswordAdmin: "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			CurrentPasswordUser:  "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		},
		Software: cablemodemutil.SoftwareStatus{
			FirmwareVersion:      "123_456_789_012_345.67890",
			CertificateInstalled: true,
			CustomerVersion:      "Foo_Bar",
			HDVersion:            "V1.0",
			DOCSISSpecVersion:    "DOCSIS 3.1",
		},
		Startup: cablemodemutil.StartupStatus{
			Boot: cablemodemutil.BootStatus{
				Status:      true,
				Operational: true,
			},
			ConfigFile: cablemodemutil.ConfigFileStatus{
				Status:  true,
				Comment: "",
			},
			Connectivity: cablemodemutil.ConnectivityStatus{
				Status:      true,
				Operational: true,
			},
			Downstream: cablemodemutil.DownstreamStatus{
				FrequencyHZ: 507000000,
				Locked:      true,
			},
			Security: cablemodemutil.SecurityStatus{
				Enabled: true,
				Comment: "BPI+",
			},
		},
		Connection: cablemodemutil.ConnectionStatus{
			SystemTime:                 time.Now(),
			UpTime:                     time.Duration(532 * time.Minute),
			DOCSISNetworkAccessAllowed: true,
			InternetConnected:          true,
			Downstream: cablemodemutil.DownstreamConnectionStatus{
				Plan:            "NorthAmerica",
				FrequencyHZ:     507000000,
				SignalPowerDBMV: -3,
				SignalSNRDB:     39,
				Channels: []cablemodemutil.DownstreamChannelInfo{
					{
						Locked:            true,
						Modulation:        "QAM256",
						ChannelID:         20,
						FrequencyHZ:       507000000,
						SignalPowerDBMV:   -4,
						SignalSNRMERDB:    39,
						CorrectedErrors:   332,
						UncorrectedErrors: 1229,
					},
					{
						Locked:            true,
						Modulation:        "QAM256",
						ChannelID:         17,
						FrequencyHZ:       483000000,
						SignalPowerDBMV:   -5,
						SignalSNRMERDB:    39,
						CorrectedErrors:   446,
						UncorrectedErrors: 1200,
					},
					{
						Locked:            true,
						Modulation:        "QAM256",
						ChannelID:         18,
						FrequencyHZ:       489000000,
						SignalPowerDBMV:   -3,
						SignalSNRMERDB:    40,
						CorrectedErrors:   379,
						UncorrectedErrors: 1272,
					},
					{
						Locked:            true,
						Modulation:        "QAM256",
						ChannelID:         19,
						FrequencyHZ:       495000000,
						SignalPowerDBMV:   -3,
						SignalSNRMERDB:    40,
						CorrectedErrors:   358,
						UncorrectedErrors: 1190,
					},
					{
						Locked:            true,
						Modulation:        "OFDM PLC",
						ChannelID:         48,
						FrequencyHZ:       850000000,
						SignalPowerDBMV:   -8,
						SignalSNRMERDB:    37,
						CorrectedErrors:   4242522,
						UncorrectedErrors: 60,
					},
				},
			},
			Upstream: cablemodemutil.UpstreamConnectionStatus{
				ChannelID: 2,
				Channels: []cablemodemutil.UpstreamChannelInfo{
					{
						Locked:          true,
						Modulation:      "SC-QAM",
						ChannelID:       1,
						WidthHZ:         3200000,
						FrequencyHZ:     10400000,
						SignalPowerDBMV: 40,
					},
					{
						Locked:          true,
						Modulation:      "SC-QAM",
						ChannelID:       2,
						WidthHZ:         6400000,
						FrequencyHZ:     16400000,
						SignalPowerDBMV: 41.3,
					},
				},
			},
		},
		Logs: []cablemodemutil.LogEntry{
			{
				Timestamp: time.Now().Add(time.Duration(-600 * time.Minute)),
				Log:       "RNG-RSP CCAP Commanded Power Exceeds Value Corresponding to the Top of the DRW;CM-MAC=<12:34:56:78:90:ab>;CMTS-MAC=<cd:ef:12:34:56:78>;CM-QOS=1.1;CM-VER=3.1;",
			},
			{
				Timestamp: time.Now().Add(time.Duration(-60 * time.Minute)),
				Log:       "Dynamic Range Window violation",
			},
			{
				Timestamp: time.Now().Add(time.Duration(-10 * time.Minute)),
				Log:       "Successful LAN WebGUI login from 192.168.1.1 on 22/01/26 at 1:46 PM.",
			},
		},
	}
	res := &cableModemStatus{
		st:  &st,
		err: nil,
	}

	return res, start.Add(cacheExpiry)
}

func newDemoModeCollector(host string) *collector {
	return &collector{
		host:  host,
		cache: NewCache(newDemoModeStatusFetcher(), nil),
	}
}
