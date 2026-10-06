package adb

import "testing"

func TestParseMemInfo(t *testing.T) {
	tests := []struct {
		name          string
		output        string
		wantTotal     uint64
		wantAvailable uint64
	}{
		{
			name: "uses MemAvailable when present",
			output: `MemTotal:        7834512 kB
MemFree:          214340 kB
MemAvailable:    3120876 kB
Buffers:            5128 kB
Cached:          2861724 kB`,
			wantTotal:     7834512,
			wantAvailable: 3120876,
		},
		{
			// Kernels older than 3.14 have no MemAvailable line.
			name: "falls back to free plus buffers plus cached",
			output: `MemTotal:        1938244 kB
MemFree:          102400 kB
Buffers:           20480 kB
Cached:           409600 kB`,
			wantTotal:     1938244,
			wantAvailable: 102400 + 20480 + 409600,
		},
		{
			name:   "empty output",
			output: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total, available := parseMemInfo(tt.output)
			if total != tt.wantTotal || available != tt.wantAvailable {
				t.Errorf("parseMemInfo() = (%d, %d), want (%d, %d)",
					total, available, tt.wantTotal, tt.wantAvailable)
			}
		})
	}
}

func TestParseNetDev(t *testing.T) {
	tests := []struct {
		name   string
		output string
		wantRx uint64
		wantTx uint64
	}{
		{
			name: "sums every interface except loopback",
			output: `Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
    lo:  524288    1024    0    0    0     0          0         0   524288    1024    0    0    0     0       0          0
 wlan0: 9000000    7000    0    0    0     0          0         0  1500000    4000    0    0    0     0       0          0
rmnet0: 1000000     800    0    0    0     0          0         0   500000     300    0    0    0     0       0          0`,
			wantRx: 9000000 + 1000000,
			wantTx: 1500000 + 500000,
		},
		{
			name:   "reads interfaces with no space after the colon",
			output: ` wlan0:123 1 0 0 0 0 0 0 456 1 0 0 0 0 0 0`,
			wantRx: 123,
			wantTx: 456,
		},
		{
			name:   "skips lines with too few fields",
			output: ` wlan0: 123 1 0`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rx, tx := parseNetDev(tt.output)
			if rx != tt.wantRx || tx != tt.wantTx {
				t.Errorf("parseNetDev() = (%d, %d), want (%d, %d)", rx, tx, tt.wantRx, tt.wantTx)
			}
		})
	}
}
