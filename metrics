package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	mWorkerUp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxy_worker_up",
			Help: "Status of worker.",
		},
		[]string{"proxy", "worker", "user"},
	)
	mPoolUp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxy_pool_up",
			Help: "Status of pool.",
		},
		[]string{"proxy", "hash", "pool"},
	)
	mPoolDivider = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxy_pool_divider",
			Help: "Divider of algorithm of pool.",
		},
		[]string{"proxy", "hash", "pool"},
	)
	mSended = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "proxy_worker_sended",
			Help: "Number of shares of worker sent to the pool.",
		},
		[]string{"proxy", "worker", "user", "hash", "pool"},
	)
	mOneSended = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "proxy_worker_one_sended",
			Help: "Number of 1-difficulty shares of worker sent to the pool.",
		},
		[]string{"proxy", "worker", "user", "hash", "pool"},
	)
	mAccepted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "proxy_worker_accepted",
			Help: "Number of shares of worker accepted by the pool.",
		},
		[]string{"proxy", "worker", "user", "hash", "pool"},
	)
	mOneAccepted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "proxy_worker_one_accepted",
			Help: "Number of 1-difficulty shares of worker accepted by the pool.",
		},
		[]string{"proxy", "worker", "user", "hash", "pool"},
	)
	mSpeed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxy_worker_speed",
			Help: "Speed of worker in hashes/s.",
		},
		[]string{"proxy", "worker", "user", "hash", "pool"},
	)
	mDifficulty = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxy_worker_difficulty",
			Help: "Value of difficulty of the worker.",
		},
		[]string{"proxy", "worker", "user", "hash", "pool"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(mWorkerUp)
	prometheus.MustRegister(mPoolUp)
	prometheus.MustRegister(mPoolDivider)
	prometheus.MustRegister(mSended)
	prometheus.MustRegister(mOneSended)
	prometheus.MustRegister(mAccepted)
	prometheus.MustRegister(mOneAccepted)
	prometheus.MustRegister(mSpeed)
	prometheus.MustRegister(mDifficulty)
}
