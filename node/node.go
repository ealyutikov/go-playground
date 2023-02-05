package node

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/elyutikov/goblockchain/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version    string
	listenAddr string
	logger     *zap.SugaredLogger

	peerLock sync.RWMutex
	peers    map[proto.NodeClient]*proto.Version

	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.EncoderConfig.TimeKey = ""
	logger, _ := loggerConfig.Build()
	return &Node{
		peers:   make(map[proto.NodeClient]*proto.Version),
		version: "blcoker-0.1",
		logger:  logger.Sugar(),
	}
}

func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()

	n.logger.Debugw("new peer connected", "addr", v.ListenAddr, "height", v.Height)

	n.peers[c] = v
}

func (n *Node) deletePeer(c proto.NodeClient) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
	delete(n.peers, c)
}

func (n *Node) BootstrapNetwork(addrs []string) error {
	for _, a := range addrs {
		c, err := makeNodeClient(a)
		if err != nil {
			return err
		}

		v, err := c.Handshake(context.Background(), n.getVersion())
		if err != nil {
			n.logger.Error("handshake error:", err)
			continue
		}

		n.addPeer(c, v)
	}
	return nil
}

func (n *Node) Start(address string) error {
	n.listenAddr = address
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	proto.RegisterNodeServer(server, n)

	n.logger.Infow("node started...", "port", n.listenAddr)

	return server.Serve(ln)
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	c, err := makeNodeClient(v.ListenAddr)
	if err != nil {
		return nil, err
	}

	n.addPeer(c, v)

	return n.getVersion(), nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	log.Println("recieved tx from:", peer)
	return &proto.Ack{}, nil
}

func (n *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version:    "goblockchain-0.1",
		Height:     0,
		ListenAddr: n.listenAddr,
	}
}

func makeNodeClient(address string) (proto.NodeClient, error) {
	c, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return proto.NewNodeClient(c), nil
}
