package cron

import (
	"github.com/robfig/cron/v3"
)

type Job interface {
	cron.Job
	Spec() string // cron表达式或内置规则
}

type Manager struct {
	cron *cron.Cron
}

func NewManager() *Manager {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(cron.Recover(cron.DefaultLogger), cron.DelayIfStillRunning(cron.DefaultLogger)),
	)
	return &Manager{c}
}


func (m *Manager) Register(job ...Job) {
	for _, j := range job {
		m.cron.AddJob(j.Spec(), j)
	}
}

func (m *Manager) RegisterFunc(frequency string, cmd func()) {
	m.cron.AddFunc(frequency, cmd)
}

func (m *Manager) Start() {
	m.cron.Start()
}

func (m *Manager) Stop() {
	m.cron.Stop()
}

func (m *Manager) Entries() []cron.Entry {
	return m.cron.Entries()
}