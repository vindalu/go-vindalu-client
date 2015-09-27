package vindalu

import (
	log "github.com/euforia/simplelog"
	"github.com/nats-io/nats"

	"github.com/euforia/vindaloo/events"
)

type VindaluSubscriber struct {
	conn   *nats.Conn
	enConn *nats.EncodedConn
	// We do not want to allow writing to the channel
	//unwritableChan chan *events.Event
}

func NewVindaluSubscriber(servers []string, logger *log.Logger) (vs *VindaluSubscriber, err error) {
	vs = &VindaluSubscriber{}

	opts := nats.DefaultOptions
	opts.Servers = servers

	if vs.conn, err = opts.Connect(); err != nil {
		return
	}
	logger.Debug.Printf("nats client connected to: %v!\n", vs.conn.ConnectedUrl())

	vs.conn.Opts.ReconnectedCB = func(nc *nats.Conn) {
		logger.Debug.Printf("nats client reconnected to: %v!\n", nc.ConnectedUrl())
	}

	vs.conn.Opts.DisconnectedCB = func(_ *nats.Conn) {
		logger.Debug.Printf("nats client disconnected!\n")
	}

	vs.enConn, err = nats.NewEncodedConn(vs.conn, nats.JSON_ENCODER)

	return
}

func (vs *VindaluSubscriber) Subscribe(topic string) (ch chan *events.Event, err error) {
	// Goes no where as we do not want to allow writing (i.e publishing)
	if err = vs.enConn.BindSendChan(topic, make(chan *events.Event)); err != nil {
		return
	}

	ch = make(chan *events.Event)
	if _, err = vs.enConn.BindRecvChan(topic, ch); err != nil {
		return
	}
	return
}

func (vs *VindaluSubscriber) Close() {
	vs.enConn.Close()
	vs.conn.Close()
}
