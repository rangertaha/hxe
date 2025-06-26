package __

import (
	"github.com/rangertaha/hxe/internal/models/db"
	pb "github.com/rangertaha/hxe/internal/models/pb"
)

func ToDatabase(p *pb.Service) *db.Service {
	return &db.Service{
		ID:          uint(p.Id),
		Name:        p.Name,
		Description: p.Description,
		User:        int(p.User),
		Group:       int(p.Group),
		Directory:   p.Directory,
		PreExec:     p.PreExec,
		CmdExec:     p.CmdExec,
		PostExec:    p.PostExec,
		Autostart:   p.Autostart,
		Retries:     int(p.Retries),
	}
}

func ToProto(d *db.Service) *pb.Service {
	return &pb.Service{
		Id:          uint32(d.ID),
		Name:        d.Name,
		Description: d.Description,
		User:        int32(d.User),
		Group:       int32(d.Group),
		Directory:   d.Directory,
		PreExec:     d.PreExec,
		CmdExec:     d.CmdExec,
		PostExec:    d.PostExec,
		Autostart:   d.Autostart,
		Retries:     int32(d.Retries),
	}
}
