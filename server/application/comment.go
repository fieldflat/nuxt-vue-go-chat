package application

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/sekky0905/nuxt-vue-go-chat/server/domain/model"
	. "github.com/sekky0905/nuxt-vue-go-chat/server/domain/repository"
	"github.com/sekky0905/nuxt-vue-go-chat/server/domain/service"
)

// CommentService is interface of CommentService.
type CommentService interface {
	ListComments(ctx context.Context, commentId uint32, limit int, cursor uint32) (*model.CommentList, error)
	GetComment(ctx context.Context, id uint32) (*model.Comment, error)
	CreateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	UpdateComment(ctx context.Context, id uint32, comment *model.Comment) (*model.Comment, error)
	DeleteComment(ctx context.Context, id uint32) error
}

// commentService is application service of comment.
type commentService struct {
	m        DBManager
	service  service.CommentService
	repo     CommentRepository
	txCloser CloseTransaction
}

// NewCommentService generates and returns CommentService.
func NewCommentApplication(m DBManager, service service.CommentService, repo CommentRepository, txCloser CloseTransaction) CommentService {
	return &commentService{
		m:        m,
		service:  service,
		repo:     repo,
		txCloser: txCloser,
	}
}

// ListThreads gets ThreadList.
func (cs *commentService) ListComments(ctx context.Context, threadId uint32, limit int, cursor uint32) (*model.CommentList, error) {
	comments, err := cs.repo.ListComments(ctx, cs.m, threadId, limit, cursor)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list comments")
	}

	return comments, nil
}

// GetComment gets Comment.
func (cs *commentService) GetComment(ctx context.Context, id uint32) (*model.Comment, error) {
	comment, err := cs.repo.GetCommentByID(ctx, cs.m, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get comment by id")
	}

	return comment, nil
}

// CreateComment creates Comment.
func (cs *commentService) CreateComment(ctx context.Context, param *model.Comment) (comment *model.Comment, err error) {
	tx, err := cs.m.Begin()
	if err != nil {
		return nil, beginTxErrorMsg(err)
	}

	defer func() {
		if err := cs.txCloser(tx, err); err != nil {
			err = errors.Wrap(err, "failed to close tx")
		}
	}()

	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()

	id, err := cs.repo.InsertComment(ctx, tx, param)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert comment")
	}
	param.ID = id

	return param, nil
}

// UpdateComment updates Comment.
func (cs *commentService) UpdateComment(ctx context.Context, id uint32, param *model.Comment) (comment *model.Comment, err error) {
	return param, nil
}

// DeleteComment deletes Comment.
func (cs *commentService) DeleteComment(ctx context.Context, id uint32) (err error) {

	return nil
}
