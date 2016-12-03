package machinetest

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"koding/klient/machine"
)

// DynamicClientOpts creates test-friendly options for dynamic client.
func DynamicClientOpts(s *Server, b *NilBuilder) machine.DynamicClientOpts {
	return machine.DynamicClientOpts{
		AddrFunc:        s.AddrFunc(),
		Builder:         b,
		DynAddrInterval: 10 * time.Millisecond,
		PingInterval:    50 * time.Millisecond,
	}
}

// TurnOnAddr returns address that "starts" test machine.
func TurnOnAddr() machine.Addr {
	return machine.Addr{
		Network:   "on",
		Value:     "0",
		UpdatedAt: time.Now(),
	}
}

// TurnOffAddr returns address that "stops" test machine.
func TurnOffAddr() machine.Addr {
	return machine.Addr{
		Network:   "off",
		Value:     "0",
		UpdatedAt: time.Now(),
	}
}

// Server simulates the real server.
type Server struct {
	mu   sync.Mutex
	addr machine.Addr
}

// AddrFunc returns machine.DynamicAddrFunc that can be used as addresses source.
func (s *Server) AddrFunc() machine.DynamicAddrFunc {
	return func(network string) (machine.Addr, error) {
		s.mu.Lock()
		defer s.mu.Unlock()

		if network == s.addr.Network {
			return s.addr, nil
		}

		return machine.Addr{}, machine.ErrAddrNotFound
	}
}

// TurnOn simulates on-line server.
func (s *Server) TurnOn() {
	s.mu.Lock()
	defer s.mu.Unlock()

	value := s.addr.Value
	if value == "" {
		value = "0" // initial address.
	}

	s.addr = machine.Addr{
		Network: "on",
		Value:   value,
	}
}

// TurnOff simulates off-line server.
func (s *Server) TurnOff() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.addr.Network = "off"
}

// ChangeAddr generates a new address for server.
func (s *Server) ChangeAddr() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.addr.Value += "0"
}

// NilBuilder uses Server logic to build nil clients.
type NilBuilder struct {
	buildsN int64
	ch      chan struct{}
}

// NewNilBuilder creates a new NilBuilder object.
func NewNilBuilder() *NilBuilder {
	return &NilBuilder{
		ch: make(chan struct{}),
	}
}

// Ping returns machine statuses based on Server response.
func (*NilBuilder) Ping(dynAddr machine.DynamicAddrFunc) (machine.Status, machine.Addr, error) {
	if dynAddr == nil {
		panic("nil dynamic function generator")
	}

	states := map[string]machine.State{
		"":    machine.StateUnknown,
		"on":  machine.StateOnline,
		"off": machine.StateOffline,
	}
	for network, state := range states {
		if addr, err := dynAddr(network); err == nil {
			return machine.Status{State: state}, addr, err
		}
	}

	return machine.Status{}, machine.Addr{}, machine.ErrAddrNotFound
}

// Build creates a nil client and increases builds counter.
func (n *NilBuilder) Build(_ context.Context, _ machine.Addr) machine.Client {
	go func() {
		atomic.AddInt64(&n.buildsN, 1)
		n.ch <- struct{}{}
	}()

	return nil
}

// WaitForBuild waits for invocation of Build method. It times out after
// specified duration.
func (n *NilBuilder) WaitForBuild(timeout time.Duration) error {
	select {
	case <-n.ch:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("timed out after %s", timeout)
	}
}

// BuildsCount returns how many times Build method was invoked.
func (n *NilBuilder) BuildsCount() int {
	return int(atomic.LoadInt64(&n.buildsN))
}
