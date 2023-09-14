package repo

import (
	"context"
	"test.com/project-user/internal/data/member"
	"test.com/project-user/internal/database"
)

type MemberRepo interface {
	FindMember(ctx context.Context, name *string, pwd *string) (mem *member.Member, err error)
	SaveMember(conn database.DbConn, ctx context.Context, mem member.Member) error
	GetMemberByEmail(ctx context.Context, username *string) (bool, error)
}
