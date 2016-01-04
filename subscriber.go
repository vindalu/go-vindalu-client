package vindalu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/euforia/simplelog"
	"github.com/nats-io/nats"

	"github.com/vindalu/vindalu/core"
)

type VindaluSubscriber struct {
	conn   *nats.Conn
	enConn *nats.EncodedConn
}

func NewVindaluSubscriber(server string, logger *log.Logger) (vs *VindaluSubscriber, err error) {
	vs = &VindaluSubscriber{}

	opts := nats.DefaultOptions
	opts.Servers, err = GetNatsServers(server)
	if err != nil {
		return
	}

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

func GetNatsServers(server string) (natsServers []string, err error) {
	var (
		clusterStatus *core.VindaluClusterStatus
		resp          *http.Response
		body          []byte
	)
	resp, err = http.Get(server + "/status")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &clusterStatus)
	if err != nil {
		return
	}
	for _, member := range clusterStatus.ClusterMemberAddrs() {
		natsServers = append(natsServers, fmt.Sprintf("nats://%s:%d", member, 4223))
	}
	return
}

func (vs *VindaluSubscriber) Subscribe(topic string) (ch chan *core.Event, err error) {

	ch = make(chan *core.Event)
	_, err = vs.enConn.BindRecvChan(topic, ch)
	return
}

func (vs *VindaluSubscriber) SubscribeQueueGroup(topic, qGroup string) (ch chan *core.Event, err error) {

	ch = make(chan *core.Event)
	_, err = vs.enConn.BindRecvQueueChan(topic, qGroup, ch)
	return
}

func (vs *VindaluSubscriber) Close() {
	vs.enConn.Close()
	vs.conn.Close()
}
