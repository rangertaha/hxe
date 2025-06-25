package __

import (
	"github.com/rangertaha/hxe/internal/models/db"
	pb "github.com/rangertaha/hxe/internal/models/pb"
)

func ToDb(p *pb.Service) *db.Service {
	return &db.Service{
		ID:          uint(p.Id),
		Name:        p.Name,
		Description: p.Description,
		User:        int(p.User),
		Group:       int(p.Group),
	}
}

func ToPb(db *db.Service) *pb.Service {
	return &pb.Service{
		Id:          uint32(db.ID),
		Name:        db.Name,
		Description: db.Description,
		User:        int32(db.User),
		Group:       int32(db.Group),
	}
}
