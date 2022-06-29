package nars

/*
* Instances of this represent a reasoner connected to a Memory, and set of Input and Output channels.
* All state is contained within Memory.  A Nar is responsible for managing I/O channels and executing
* It executes series sof cycles in two possible modes:
* step mode - controlled by an outside system, such as during debugging or testing
* goroutine mode - runs in a passable closed-loop at a specific maximum framerate.
 */

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nars_go/control"
	"nars_go/io"
	"nars_go/language"
	"nars_go/plugin/perception"
	"nars_go/storage"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var ErrInterrupt = errors.New("received interrupt")

type Nars struct {
	sensoryChannels map[*language.Term]perception.SensoryChannel
	interrupt       chan os.Signal
	stop            chan error
	memory          *storage.Memory
	inputLine       chan string
}

func New() *Nars {
	return &Nars{
		interrupt:       make(chan os.Signal, 1),
		stop:            make(chan error),
		memory:          storage.NewMemory(),
		sensoryChannels: make(map[*language.Term]perception.SensoryChannel, 0),
		inputLine:       make(chan string),
	}
}

func (n *Nars) AddChannel(termStr string, channel perception.SensoryChannel) {
	term, err := io.NewNarsese().ParseTerm(termStr)
	if err != nil {
		return
	}
	n.sensoryChannels[term] = channel
}

func (n *Nars) SetInputLine(line string) {
	n.inputLine <- line
}

func (n *Nars) Start() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	// 1. 监听退出信号
	signal.Notify(n.interrupt, syscall.SIGINT, syscall.SIGKILL) // 监听处理 ctrl+c syscall.SIGINT
	go func() {
		<-n.interrupt
		cancel()
		log.Printf("task is cancel please wait for stop...")
		wg.Wait() // 等待所有任务完成再退出
		n.stop <- ErrInterrupt
	}()

	go func() {
		defer wg.Done()
		n.run(ctx)
	}()

	select {
	case <-n.stop:
		break
	}
}

func (n *Nars) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case line := <-n.inputLine:
			fmt.Println(line)
		default:
			control.Inference(n.memory)
			time.Sleep(time.Second)
		}
	}
}
