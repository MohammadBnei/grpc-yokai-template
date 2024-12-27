package internal

import (
	"github.com/ankorstore/yokai/fxgrpcserver"
	"go.uber.org/fx"

	"github.com/mohammadbnei/vos-personal-words/internal/service"
	"github.com/mohammadbnei/vos-personal-words/proto"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fxgrpcserver.AsGrpcServerService(service.NewPersonalService, &proto.PersonalWordService_ServiceDesc),
	)
}
