package node

import (
	"context"
	"encoding/hex"
	"net"
	"sync"
	"time"

	"github.com/elyutikov/goblockchain/crypto"
	"github.com/elyutikov/goblockchain/proto"
	"github.com/elyutikov/goblockchain/types"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const blockTime = 5 * time.Second

type ServerConfig struct {
	Version    string
	ListenAddr string
	PrivateKey *crypto.PrivateKey
}

type Node struct {
	ServerConfig
	logger *zap.SugaredLogger

	peerLock sync.RWMutex
	peers    map[proto.NodeClient]*proto.Version
	mempool  *Mempool

	proto.UnimplementedNodeServer
}

func NewNode(config ServerConfig) *Node {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.EncoderConfig.TimeKey = ""
	logger, _ := loggerConfig.Build()
	return &Node{
		peers:        make(map[proto.NodeClient]*proto.Version),
		logger:       logger.Sugar(),
		mempool:      NewMempool(),
		ServerConfig: config,
	}
}

func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()

	// Handle the logic where we decide to accept or drop
	// the incoming connection.

	n.peers[c] = v

	// Connect to all peers in the received list of peers.
	if len(v.PeerList) > 0 {
		go n.bootstrapNetwork(v.PeerList)
	}

	n.logger.Debugw("new peer successfully connected",
		"we", n.ListenAddr,
		"remote", v.ListenAddr,
		"height", v.Height)
}

func (n *Node) deletePeer(c proto.NodeClient) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
	delete(n.peers, c)
}

func (n *Node) bootstrapNetwork(addrs []string) error {
	for _, a := range addrs {
		if !n.canConnectWith(a) {
			continue
		}
		n.logger.Debugf("dialing remote node", "we", n.ListenAddr, "remote", a)
		c, v, err := n.dialRemoteNode(a)
		if err != nil {
			return err
		}
		n.addPeer(c, v)
	}
	return nil
}

func (n *Node) Start(address string, bootstrapNodes []string) error {
	n.ListenAddr = address
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	proto.RegisterNodeServer(server, n)

	n.logger.Infow("node started...", "port", n.ListenAddr)

	// bootstrap the network with a list of already known nodes
	// in the network
	if len(bootstrapNodes) > 0 {
		go n.bootstrapNetwork(bootstrapNodes)
	}

	if n.PrivateKey != nil {
		go n.validateLoop()
	}

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
	hash := hex.EncodeToString(types.HashTransaction(tx))

	if n.mempool.Add(tx) {
		n.logger.Debugw("received tx", "from", peer.Addr, "hash", hash, "we", n.ListenAddr)
		go func() {
			err := n.broadcast(tx)
			if err != nil {
				n.logger.Errorw("broadcast error", "err", err)
			}
		}()
	}

	return &proto.Ack{}, nil
}

func (n *Node) validateLoop() {
	n.logger.Infow("starting validating loop", "pubkey", n.PrivateKey.Public(), "blockTime", blockTime)
	ticker := time.NewTicker(blockTime)
	for {
		<-ticker.C
		n.logger.Debugw("time to create new block", "lenTx", len(n.mempool.txx))

		for hash := range n.mempool.txx {
			delete(n.mempool.txx, hash)
		}
	}
}

func (n *Node) broadcast(msg any) error {
	for peer := range n.peers {
		switch v := msg.(type) {
		case *proto.Transaction:
			_, err := peer.HandleTransaction(context.Background(), v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version:    "goblockchain-0.1",
		Height:     0,
		ListenAddr: n.ListenAddr,
		PeerList:   n.getPeerList(),
	}
}

func (n *Node) getPeerList() []string {
	n.peerLock.RLock()
	defer n.peerLock.RUnlock()

	peers := []string{}
	for _, version := range n.peers {
		peers = append(peers, version.ListenAddr)
	}
	return peers
}

func (n *Node) dialRemoteNode(addr string) (proto.NodeClient, *proto.Version, error) {
	c, err := makeNodeClient(addr)
	if err != nil {
		return nil, nil, err
	}

	v, err := c.Handshake(context.Background(), n.getVersion())
	if err != nil {
		return nil, nil, err
	}

	return c, v, nil
}

func (n *Node) canConnectWith(addr string) bool {
	if n.ListenAddr == addr {
		return false
	}

	for _, connectedAddr := range n.getPeerList() {
		if addr == connectedAddr {
			return false
		}
	}

	return true
}

func makeNodeClient(address string) (proto.NodeClient, error) {
	client, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return proto.NewNodeClient(client), nil
}
