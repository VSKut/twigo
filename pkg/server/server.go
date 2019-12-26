package server

import (
	"context"
	"database/sql"
	"errors"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vskut/twigo/pkg/common/entity"
	"github.com/vskut/twigo/pkg/common/repository"
	"github.com/vskut/twigo/pkg/common/token"
	proto "github.com/vskut/twigo/pkg/grpc"
	"google.golang.org/grpc/health"
	hlzpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server ...
type Server struct {
	repo *repository.Repository
}

// NewServer constructs the Server struct
func NewServer(db *sql.DB) *Server {
	return &Server{
		repo: repository.NewPostgreSQLRepository(db),
	}
}

// AuthFuncOverride allows a given gRPC service implementation to override the global `JwtMiddleware`.
func (s *Server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	switch fullMethodName {
	case "/tweet.TweetService/ListTweet", "/tweet.TweetService/CreateTweet", "/user.UserService/Subscribe":
		return token.JwtMiddleware(ctx)
	default:
		return ctx, nil
	}
}

// Run starts the grpc server
func (s *Server) Run(addr string) error {
	tcp, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(token.JwtMiddleware),
		)),
	)

	hlzpb.RegisterHealthServer(grpcServer, health.NewServer())

	proto.RegisterAuthServiceServer(grpcServer, s)
	proto.RegisterTweetServiceServer(grpcServer, s)
	proto.RegisterUserServiceServer(grpcServer, s)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(tcp); err != nil {
		return err
	}

	return nil
}

// Login implements logic of authentication
func (s *Server) Login(_ context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	if err := request.Validate(); err != nil {
		return &proto.LoginResponse{}, err
	}

	user, err := s.repo.Users.GetByEmail(request.GetEmail())
	if err != nil {
		return &proto.LoginResponse{}, errors.New("invalid LoginRequest.Email: wrong email provided")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.GetPassword())); err != nil {
		return &proto.LoginResponse{}, errors.New("invalid LoginRequest.Password: wrong password provided")
	}

	jwtToken, err := token.CreateToken(user)
	if err != nil {
		return &proto.LoginResponse{}, errors.New("invalid LoginRequest: jwt-token creation error")
	}

	return &proto.LoginResponse{Token: jwtToken}, nil
}

// Register implements logic of registration
func (s *Server) Register(_ context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	if err := request.Validate(); err != nil {
		return &proto.RegisterResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return &proto.RegisterResponse{}, err
	}

	user, err := s.repo.Users.Save(entity.User{
		Email:    request.GetEmail(),
		Password: string(hash),
		Username: request.GetUsername(),
	})
	if err != nil {
		return &proto.RegisterResponse{}, err
	}

	return &proto.RegisterResponse{
		Id:       user.ID,
		Username: user.Username,
	}, nil
}

// Subscribe implements logic of subscription
func (s *Server) Subscribe(ctx context.Context, request *proto.SubscribeRequest) (*proto.SubscribeResponse, error) {
	if err := request.Validate(); err != nil {
		return &proto.SubscribeResponse{}, err
	}

	authUser, ok := ctx.Value(token.ValueTokenContextKey).(entity.User)
	if !ok {
		return &proto.SubscribeResponse{}, status.Error(codes.Unauthenticated, "Authentication required")
	}

	if err := s.repo.Users.Subscribe(authUser, request.GetUsername()); err != nil {
		return &proto.SubscribeResponse{}, err
	}

	return &proto.SubscribeResponse{}, nil
}

// CreateTweet implements logic of creating new tweet
func (s *Server) CreateTweet(ctx context.Context, request *proto.CreateTweetRequest) (*proto.CreateTweetResponse, error) {
	if err := request.Validate(); err != nil {
		return &proto.CreateTweetResponse{}, err
	}

	authUser, ok := ctx.Value(token.ValueTokenContextKey).(entity.User)
	if !ok {
		return &proto.CreateTweetResponse{}, status.Error(codes.Unauthenticated, "Authentication required")
	}

	tweet, err := s.repo.Tweets.Save(entity.Tweet{
		Message: request.GetMessage(),
		UserID:  authUser.ID,
	})
	if err != nil {
		return &proto.CreateTweetResponse{}, err
	}

	return &proto.CreateTweetResponse{Id: tweet.ID, Message: tweet.Message}, nil
}

// ListTweet implements logic of listing user's tweets
func (s *Server) ListTweet(ctx context.Context, _ *proto.ListTweetRequest) (*proto.ListTweetResponse, error) {
	authUser, ok := ctx.Value(token.ValueTokenContextKey).(entity.User)
	if !ok {
		return &proto.ListTweetResponse{}, status.Error(codes.Unauthenticated, "Authentication required")
	}

	res, err := s.repo.Tweets.ListAllByUser(authUser)
	if err != nil {
		return &proto.ListTweetResponse{}, err
	}

	var tweets []*proto.ListTweetResponse_Tweet
	for _, v := range res {
		tweets = append(tweets, &proto.ListTweetResponse_Tweet{
			Id:      v.ID,
			Message: v.Message,
		})
	}

	return &proto.ListTweetResponse{Tweets: tweets}, nil
}
