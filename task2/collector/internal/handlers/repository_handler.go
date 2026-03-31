package handlers

import (
	"collector/internal/domain"

	"collector/api/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type RepositoryHandler struct {
	repositoryUsecases RepositoryUsecases
}

type RepositoryUsecases interface {
	GetRepository(authorName, repoName string) (*domain.Repository, error)
}

func NewRepositoryHandler(repositoryUsecases RepositoryUsecases) *RepositoryHandler {
	return &RepositoryHandler{repositoryUsecases}
}

func (rh *RepositoryHandler) GetRepository(request *proto.GetRepositoryRequest) (*proto.GetRepositoryReply, error) {
	repo, err := rh.repositoryUsecases.GetRepository(request.OwnerName, request.RepoName)
	if err != nil {
		return &proto.GetRepositoryReply{}, err
	}

	reply := &proto.GetRepositoryReply{
		RepoName:   repo.Name,
		OwnerName:  repo.Owner.Login,
		Id:         int32(repo.Id),
		StarsCount: int32(repo.StarsCount),
		ForksCount: int32(repo.ForksCount),
		CreatedAt:  timestamppb.New(repo.CreatedAt),
		IsFork:     repo.IsFork,
		ForkedFrom: &repo.Parent.FullName,
		OpenIssues: int32(repo.OpenIssuesCount),
	}
	return reply, nil
}
