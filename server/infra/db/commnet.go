package db

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sekky0905/nuxt-vue-go-chat/server/domain/model"
	"github.com/sekky0905/nuxt-vue-go-chat/server/domain/repository"
	. "github.com/sekky0905/nuxt-vue-go-chat/server/domain/repository"
	"github.com/sekky0905/nuxt-vue-go-chat/server/infra/logger"
	"go.uber.org/zap"
)

// commentRepository is repository of comment.
type commentRepository struct {
}

// NewCommentRepository generates and returns CommentRepository.
func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}

// ErrorMsg generates and returns error message.
func (repo *commentRepository) ErrorMsg(method model.RepositoryMethod, err error) error {
	return &model.RepositoryError{
		BaseErr:                     err,
		RepositoryMethod:            method,
		DomainModelNameForDeveloper: model.DomainModelNameCommentForDeveloper,
		DomainModelNameForUser:      model.DomainModelNameCommentForUser,
	}
}

// ListThreads lists ThreadList.
func (repo *commentRepository) ListComments(ctx context.Context, m SQLManager, threadID uint32, limit int, cursor uint32) (*model.CommentList, error) {
	query := `SELECT c.id, c.content, u.id, u.name, c.thread_id, c.created_at, c.updated_at
	FROM comments AS c
	INNER JOIN users AS u
	ON c.user_id = u.id
	WHERE c.id > ?
	AND c.thread_id = ?
	ORDER BY c.id ASC
	LIMIT ?;`

	limitForCheckHasNext := readyLimitForHasNext(limit)
	comments, err := repo.list(ctx, m, model.RepositoryMethodREAD, query, cursor, threadID, limitForCheckHasNext)

	length := len(comments)

	if length == 0 {
		err = &model.NoSuchDataError{
			BaseErr:                     err,
			DomainModelNameForDeveloper: model.DomainModelNameCommentForDeveloper,
			DomainModelNameForUser:      model.DomainModelNameCommentForUser,
		}
		return nil, err
	}

	if err != nil {
		return nil, repo.ErrorMsg(model.RepositoryMethodLIST, errors.WithStack(err))
	}

	hasNext := checkHasNext(length, limit)
	if hasNext {
		cursor = comments[limitForCheckHasNext-1].ID
	} else {
		cursor = 0
	}

	if length == limitForCheckHasNext {
		// exclude thread for cursor
		return &model.CommentList{
			Comments: comments[:limitForCheckHasNext-1],
			HasNext:  hasNext,
			Cursor:   cursor,
		}, nil
	}

	return &model.CommentList{
		Comments: comments,
		HasNext:  hasNext,
		Cursor:   cursor,
	}, nil
}

// GetThreadByID gets and returns a record specified by id.
func (repo *commentRepository) GetCommentByID(ctx context.Context, m SQLManager, id uint32) (*model.Comment, error) {
	query := `SELECT c.id, c.content, u.id, u.name, c.thread_id, c.created_at, c.updated_at
	FROM comments AS c
	INNER JOIN users AS u
	ON c.user_id = u.id
	WHERE c.id=?
	LIMIT 1;`

	comments, err := repo.list(ctx, m, model.RepositoryMethodREAD, query, id)

	if len(comments) == 0 {
		err = &model.NoSuchDataError{
			BaseErr:                     err,
			PropertyNameForDeveloper:    model.IDPropertyForDeveloper,
			PropertyNameForUser:         model.IDPropertyForUser,
			PropertyValue:               id,
			DomainModelNameForDeveloper: model.DomainModelNameCommentForDeveloper,
			DomainModelNameForUser:      model.DomainModelNameCommentForUser,
		}
		return nil, err
	}

	if err != nil {
		return nil, repo.ErrorMsg(model.RepositoryMethodLIST, errors.WithStack(err))
	}

	return comments[0], nil
}

// list gets and returns list of records.
func (repo *commentRepository) list(ctx context.Context, m repository.DBManager, method model.RepositoryMethod, query string, args ...interface{}) (comments []*model.Comment, err error) {
	stmt, err := m.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.WithStack(repo.ErrorMsg(method, err))
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Logger.Error("stmt.Close", zap.String("error message", err.Error()))
		}
	}()

	rows, err := stmt.QueryContext(ctx, args...)

	if err != nil {
		return nil, repo.ErrorMsg(method, errors.WithStack(err))
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			logger.Logger.Error("rows.Close", zap.String("error message", err.Error()))
		}
	}()

	list := make([]*model.Comment, 0)
	for rows.Next() {
		comment := &model.Comment{
			User: &model.User{},
		}

		err = rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.User.ID,
			&comment.User.Name,
			&comment.ThreadID,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)

		if err != nil {
			return nil, repo.ErrorMsg(method, errors.WithStack(err))
		}

		list = append(list, comment)
	}

	return list, nil
}

// InsertThread insert a record.
func (repo *commentRepository) InsertComment(ctx context.Context, m SQLManager, comment *model.Comment) (uint32, error) {
	query := "INSERT INTO comments (content, user_id, thread_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?);"
	stmt, err := m.PrepareContext(ctx, query)
	if err != nil {
		return model.InvalidID, errors.WithStack(repo.ErrorMsg(model.RepositoryMethodInsert, err))
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Logger.Error("stmt.Close", zap.String("error message", err.Error()))
		}
	}()

	result, err := stmt.ExecContext(ctx, comment.Content, comment.User.ID, comment.ThreadID, comment.CreatedAt, comment.UpdatedAt)
	if err != nil {
		return model.InvalidID, repo.ErrorMsg(model.RepositoryMethodInsert, errors.WithStack(err))
	}

	affect, err := result.RowsAffected()
	if affect != 1 {
		err = fmt.Errorf("total affected id: %d ", affect)
		return model.InvalidID, repo.ErrorMsg(model.RepositoryMethodInsert, errors.WithStack(err))
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.InvalidID, repo.ErrorMsg(model.RepositoryMethodInsert, errors.WithStack(err))
	}

	return uint32(id), nil
}

// UpdateComment updates a record.
func (repo *commentRepository) UpdateComment(ctx context.Context, m SQLManager, id uint32, comment *model.Comment) error {
	query := "UPDATE comments SET content=?, updated_at=? WHERE id=?;"

	stmt, err := m.PrepareContext(ctx, query)
	if err != nil {
		return errors.WithStack(repo.ErrorMsg(model.RepositoryMethodUPDATE, err))
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Logger.Error("stmt.Close", zap.String("error message", err.Error()))
		}
	}()

	if err != nil {
		return repo.ErrorMsg(model.RepositoryMethodUPDATE, errors.WithStack(err))
	}

	result, err := stmt.ExecContext(ctx, comment.Content, comment.UpdatedAt, id)
	if err != nil {
		return repo.ErrorMsg(model.RepositoryMethodUPDATE, errors.WithStack(err))
	}

	affect, err := result.RowsAffected()
	if affect != 1 {
		err = fmt.Errorf("total affected id: %d ", affect)
		return repo.ErrorMsg(model.RepositoryMethodUPDATE, errors.WithStack(err))
	}

	return nil
}

// DeleteComment delete a record.
func (repo *commentRepository) DeleteComment(ctx context.Context, m SQLManager, id uint32) error {
	query := "DELETE FROM comments WHERE id=?;"

	stmt, err := m.PrepareContext(ctx, query)
	if err != nil {
		return repo.ErrorMsg(model.RepositoryMethodDELETE, errors.WithStack(err))
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Logger.Error("stmt.Close", zap.String("error message", err.Error()))
		}
	}()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return repo.ErrorMsg(model.RepositoryMethodDELETE, errors.WithStack(err))
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return repo.ErrorMsg(model.RepositoryMethodDELETE, errors.WithStack(err))
	}
	if affect != 1 {
		err = fmt.Errorf("total affected id: %d ", affect)
		return repo.ErrorMsg(model.RepositoryMethodDELETE, errors.WithStack(err))
	}

	return nil
}