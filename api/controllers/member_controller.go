package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/codinomello/weebie-go/api/models"
	"github.com/codinomello/weebie-go/api/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MemberController struct {
	MemberRepository repositories.MemberRepository
}

func NewMemberController(memberRepo repositories.MemberRepository) *MemberController {
	return &MemberController{
		MemberRepository: memberRepo,
	}
}

func (c *MemberController) CreateMember(member *models.Member) (*models.Member, error) {
	if member.ProjectID.IsZero() || member.UserID.IsZero() {
		return nil, fmt.Errorf("projectID e userID são obrigatórios")
	}

	member.JoinedAt = time.Now()
	member.Status = "active"

	return c.MemberRepository.CreateMember(context.Background(), member)
}

func (c *MemberController) GetMembersByProject(projectID primitive.ObjectID) ([]models.Member, error) {
	return c.MemberRepository.GetMembersByProject(context.Background(), projectID)
}

func (c *MemberController) UpdateMemberRole(memberID primitive.ObjectID, role string) error {
	return c.MemberRepository.UpdateMemberRole(context.Background(), memberID, role)
}

func (c *MemberController) DeleteMember(memberID primitive.ObjectID) error {
	return c.MemberRepository.DeleteMember(context.Background(), memberID)
}
