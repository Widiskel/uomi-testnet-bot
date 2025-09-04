package ui

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/pterm/pterm"
	"github.com/widiskel/uomi-testnet-bot/internal/domain/model"
)

var (
	multi    *pterm.MultiPrinter
	spinners = make(map[int]*pterm.SpinnerPrinter)
	mu       sync.Mutex
)

func StartUISystem() {
	m, _ := pterm.DefaultMultiPrinter.Start()
	multi = m
}

func StopUISystem() {
	if multi != nil {
		multi.Stop()
	}
}

func UpdateStatus(session model.Session, status string, remainingDelay time.Duration) {
	mu.Lock()
	defer mu.Unlock()

	balanceStr := formatBalances(session.WalletBalance)
	delayStr := FormatDelay(remainingDelay)

	content := fmt.Sprintf(`
=============== Account %d ================
Address  : %s
Balances : %s

Status   : %s
Delay    : %s
========================================`,
		session.AccIdx+1,
		session.Address,
		balanceStr,
		status,
		delayStr)

	if spinner, ok := spinners[session.AccIdx]; ok {
		spinner.UpdateText(content)
	} else {
		spinner, _ := pterm.DefaultSpinner.
			WithWriter(multi.NewWriter()).
			WithRemoveWhenDone(false).
			Start(content)
		spinners[session.AccIdx] = spinner
	}
}

func SetSpinnerSuccess(session model.Session, finalMessage string) {
	mu.Lock()
	defer mu.Unlock()
	if spinner, ok := spinners[session.AccIdx]; ok {
		UpdateStatus(session, finalMessage, 0)
		spinner.Success()
	}
}

func SetSpinnerError(session model.Session, finalMessage string) {
	mu.Lock()
	defer mu.Unlock()
	if spinner, ok := spinners[session.AccIdx]; ok {
		UpdateStatus(session, finalMessage, 0)
		spinner.Fail()
	}
}

func FormatDelay(d time.Duration) string {
	d = d.Round(time.Second)
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d H %02d M %02d S", h, m, s)
}

func formatBalances(wallet model.WalletBalance) string {
	if len(wallet.Balances) == 0 {
		return ""
	}

	var builder strings.Builder
	for _, tb := range wallet.Balances {
		builder.WriteString(fmt.Sprintf("\n- %s : %s %s", tb.Token.Symbol, tb.BalanceStr, tb.Token.Symbol))
	}

	return builder.String()
}
