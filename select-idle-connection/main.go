package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Conn struct {
	ID string
	// It may contains some kind of connections like database, websocket and etc.
}

type Pool struct {
	idle chan *Conn
}

type Dispatcher struct {
	// Dispatcher knows the list of connections.
	pools []*Pool
}

func (s *Dispatcher) Select() (*Conn, error) {
	// Convert the list of cases
	cases := make([]reflect.SelectCase, len(s.pools))
	for i, pool := range s.pools {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(pool.idle),
		}
	}
	log.Print("dispatcher waiting any idle connection...")
	_, recv, ok := reflect.Select(cases)
	if !ok {
		return nil, errors.New("failed to select a case from connection pool")
	}
	conn, ok := recv.Interface().(*Conn)
	if !ok {
		return nil, errors.New("the type of received connection is invalid")
	}

	return conn, nil
}

func main() {
	// Prepare empty pools
	pools := make([]*Pool, 0)
	for i := 0; i < 4; i++ {
		p := new(Pool)
		p.idle = make(chan *Conn)
		pools = append(pools, p)
	}

	d := Dispatcher{pools: pools}

	// Notify a connection becomes idle
	go func() {
		for i, pool := range pools {
			c := &Conn{ID: fmt.Sprintf("%d", i)}
			pool.idle <- c
		}
	}()

	// Wait and select an idle connection from pools
	for i := 0; i < 4; i++ {
		selected, err := d.Select()
		if err != nil {
			fmt.Printf("err: %#v", err)
		}
		fmt.Printf("selected Connection ID: %s\n", selected.ID)
	}
}
