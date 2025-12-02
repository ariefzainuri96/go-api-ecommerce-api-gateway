package grpc

import (
    authpb "github.com/ariefzainuri96/go-api-ecommerce-api-gateway/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

type AuthGRPCClient struct {
    Client authpb.AuthServiceClient
	Conn   *grpc.ClientConn
}

func NewAuthGRPCClient(addr string) (*AuthGRPCClient, error) {
    conn, err := grpc.NewClient(
        addr,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
        grpc.WithUnaryInterceptor(MetadataInterceptor()),
    )
    if err != nil {
        return nil, err
    }

    client := authpb.NewAuthServiceClient(conn)
    return &AuthGRPCClient{
        Client: client,
        Conn:   conn,
    }, nil
}

func (c *AuthGRPCClient) Close() error {
    return c.Conn.Close()
}
