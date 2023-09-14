package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"test.com/project-user/internal/data/member"
	"test.com/project-user/internal/database"
	"test.com/project-user/internal/database/gorms"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m *MemberDao) FindMember(ctx context.Context, name *string, pwd *string) (*member.Member, error) {
	var mem *member.Member
	err := m.conn.Session(ctx).Where("name=? and pwd=?", name, pwd).First(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return mem, err
}

func (m *MemberDao) SaveMember(conn database.DbConn, ctx context.Context, mem member.Member) error {
	m.conn = conn.(*gorms.GormConn)
	return m.conn.Tx(ctx).Create(&mem).Error
}

func (m *MemberDao) GetMemberByEmail(ctx context.Context, username *string) (bool, error) {
	if m.conn == nil {
		return false, fmt.Errorf("database connection is nil")
	}
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("name=?", username).Count(&count).Error
	return count > 0, err
}

func NewMemberDao() *MemberDao {
	return &MemberDao{
		conn: gorms.New(),
	}
}
