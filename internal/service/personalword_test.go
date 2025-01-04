package service_test

import (
	"context"
	"testing"

	"github.com/ankorstore/yokai/grpcserver/grpcservertest"
	"github.com/ankorstore/yokai/log/logtest"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/mohammadbnei/vos-template/internal"
	"github.com/mohammadbnei/vos-template/proto"
)

func TestSaveWord(t *testing.T) {
	var connFactory grpcservertest.TestBufconnConnectionFactory
	var logBuffer logtest.TestLogBuffer

	internal.RunTest(t, fx.Populate(&connFactory, &logBuffer))

	// conn preparation
	conn, err := connFactory.Create(
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)

	defer func() {
		err = conn.Close()
		assert.NoError(t, err)
	}()

	// client preparation
	client := proto.NewPersonalWordServiceClient(conn)

	// Prepare request with valid data
	req := &proto.SaveWordRequest{
		Word:     "example",
		UserId:   "user123",
		ListName: "favorites",
	}

	// call
	response, err := client.SaveWord(context.Background(), req)
	assert.NoError(t, err)

	// response assertions
	assert.NotNil(t, response.GetWord())
	assert.NotNil(t, response.GetList())
	assert.Equal(t, "example", response.GetWord().GetWord())
	assert.Equal(t, "user123", response.GetWord().GetUserId())
	assert.Equal(t, "favorites", response.GetList().GetName())
}

func TestSaveWord_InvalidArgument_EmptyWord(t *testing.T) {
	var connFactory grpcservertest.TestBufconnConnectionFactory
	var logBuffer logtest.TestLogBuffer

	internal.RunTest(t, fx.Populate(&connFactory, &logBuffer))

	// conn preparation
	conn, err := connFactory.Create(
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)

	defer func() {
		err = conn.Close()
		assert.NoError(t, err)
	}()

	// client preparation
	client := proto.NewPersonalWordServiceClient(conn)

	// Prepare request with empty word
	req := &proto.SaveWordRequest{
		UserId:   "user123",
		ListName: "favorites",
	}

	// call
	_, err = client.SaveWord(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestSaveWord_InvalidArgument_EmptyUserID(t *testing.T) {
	var connFactory grpcservertest.TestBufconnConnectionFactory
	var logBuffer logtest.TestLogBuffer

	internal.RunTest(t, fx.Populate(&connFactory, &logBuffer))

	// conn preparation
	conn, err := connFactory.Create(
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)

	defer func() {
		err = conn.Close()
		assert.NoError(t, err)
	}()

	// client preparation
	client := proto.NewPersonalWordServiceClient(conn)

	// Prepare request with empty user ID
	req := &proto.SaveWordRequest{
		Word:     "example",
		ListName: "favorites",
	}

	// call
	_, err = client.SaveWord(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}
