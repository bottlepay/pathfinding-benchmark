package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const lndManageJHost = "lnd-managej"

type lndManageJConnection struct {
	lndConnection

	cacheReset          bool
	lndSynced           bool
	localChannelsSynced bool
}

func getLndManageJConnection(alias string) (nodeInterface, error) {
	lnd, err := getLndConnection(alias)
	if err != nil {
		return nil, err
	}

	return &lndManageJConnection{
		lndConnection: *lnd,
	}, nil
}

func (l *lndManageJConnection) resetCache() error {
	resp, err := http.Get(
		fmt.Sprintf(
			"http://%v:8081/api/payments/reset-graph-cache",
			lndManageJHost,
		),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("cache reset error")
	}

	log.Debugw("Cache reset")

	return nil
}

func (l *lndManageJConnection) localChannelCount() (int, error) {
	resp, err := http.Get(
		fmt.Sprintf(
			"http://%v:8081/api/status/all-channels",
			lndManageJHost,
		),
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, errors.New("all-channels error")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var channels struct {
		Channels []interface{}
	}

	if err = json.Unmarshal(body, &channels); err != nil {
		return 0, err
	}

	return len(channels.Channels), nil
}

func (l *lndManageJConnection) IsSynced(totalEdges, localChannels int) (bool, error) {
	if !l.lndSynced {
		lndSynced, err := l.lndConnection.IsSynced(totalEdges, localChannels)
		if err != nil {
			return false, err
		}
		if !lndSynced {
			return false, nil
		}

		l.lndSynced = true
	}

	// Reset cache if not yet done.
	if !l.cacheReset {
		err := l.resetCache()
		if err != nil {
			return false, err
		}

		l.cacheReset = true
	}

	if !l.localChannelsSynced {
		channels, err := l.localChannelCount()
		if err != nil {
			return false, err
		}

		log.Debugw("Syncing local channels",
			"channels", channels, "totalChannels", localChannels)

		if channels != localChannels {
			return false, nil
		}

		l.localChannelsSynced = true
	}

	return true, nil
}

type TimestampAndMessage struct {
	Timestamp string
	Message   string
}

func (l *lndManageJConnection) SendPayment(invoice string, aliasMap map[string]string) error {
	requestBodyReader := strings.NewReader("{\"feeRateWeight\": 1}")
	url := fmt.Sprintf("http://%v:8081/api/payments/pay-payment-request/%v", lndManageJHost, invoice)
	resp, err := http.Post(url, "application/json", requestBodyReader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("payment error")
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		var timestampAndMessage TimestampAndMessage
		line := scanner.Text()
		_ = json.Unmarshal([]byte(line), &timestampAndMessage)
		timestamp, _ := time.Parse(time.RFC3339Nano, timestampAndMessage.Timestamp)
		formattedTimestamp := timestamp.Format("2006-01-02T15:04:05.000Z07")
		fmt.Printf("%v: %v\n", formattedTimestamp, timestampAndMessage.Message)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
