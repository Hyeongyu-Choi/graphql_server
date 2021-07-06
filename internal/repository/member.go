package repository

import (
	"errors"
	"github.com/Hyeongyu-Choi/graphql_server/internal/repository/entity"
)

func (r *repository) Join(member *entity.Member) (uint, error) {
	if err := r.validateDuplicateMember(member); err != nil {
		return 0, err
	}
	if err := r.db.Create(&member).Error; err != nil {
		return 0, err
	}

	return member.ID, nil
}

func (r *repository) validateDuplicateMember(member *entity.Member) error {
	var count int64
	r.db.Model(&entity.Member{}).Where("name = ?", member.Name).Count(&count)

	if count > 0 {
		return errors.New("이미 등록된 이름입니다")
	}

	return nil
}

func (r *repository) FindMembers() (members []*entity.Member, err error) {
	return members, r.db.Find(&members).Error
}

func (r *repository) FindMemberByID(id uint) (member *entity.Member, err error) {
	return member, r.db.First(&member, id).Error
}

func (r *repository) UpdateMember(id uint, name string) error {
	return r.db.Model(&entity.Member{}).Where(id).Update("name", name).Error
}
