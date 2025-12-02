package grpc

type ServerService struct {
	AuthClient *AuthGRPCClient
}

func NewServerService(auth *AuthGRPCClient) *ServerService {
	return &ServerService{
		AuthClient: auth,
	}
}
