package sharedi2c

import "github.com/kidoman/embd"
import _ "github.com/kidoman/embd/host/all"
import "fmt"

var portMap map[byte]chan I2CMsg

type SharedWriter struct {
	busId byte
	msgChan chan I2CMsg
}

func init() {
	portMap = make(map[byte]chan I2CMsg)
}

func NewSharedWriter(busId byte) *SharedWriter {
	msgChan, ok := portMap[busId]
	if (!ok) {
		fmt.Println("allocating new channel")
		msgChan = make(chan I2CMsg)
		portMap[busId] = msgChan

	}
	writer := new(SharedWriter)
	writer.msgChan = msgChan
	writer.busId = busId
	if (!ok) {
		go busWriter(writer)
	}
	return writer
}

func (w *SharedWriter) SendMsg(m I2CMsg) {
	w.msgChan <- m
}

func busWriter(config *SharedWriter) {
	// fmt.Println("starting bus writer ", config.busId)
	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()
	bus := embd.NewI2CBus(config.busId)
	for {
		msg, more := <-config.msgChan
		if more {
			// fmt.Println("got msg ")
			bus.WriteByte(msg.Addr, msg.Value)
		} else {
			fmt.Println("got all msgs, exiting")
			return
		}
	}
}

