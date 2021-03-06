package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sekky0905/nuxt-vue-go-chat/server/domain/model"
	"github.com/sekky0905/nuxt-vue-go-chat/server/domain/repository"
	mock_repository "github.com/sekky0905/nuxt-vue-go-chat/server/domain/repository/mock"
	"github.com/sekky0905/nuxt-vue-go-chat/server/domain/service"
	mock_service "github.com/sekky0905/nuxt-vue-go-chat/server/domain/service/mock"
	"github.com/sekky0905/nuxt-vue-go-chat/server/infra/db/query"
	mock_query "github.com/sekky0905/nuxt-vue-go-chat/server/infra/db/query/mock"
	"github.com/sekky0905/nuxt-vue-go-chat/server/testutil"
)

func Test_commentService_ListComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testutil.SetFakeTime(time.Now())

	type fields struct {
		m        query.DBManager
		repo     repository.CommentRepository
		service  service.CommentService
		txCloser CloseTransaction
	}
	type args struct {
		ctx      context.Context
		threadID uint32
		limit    int
		cursor   uint32
	}

	type mockReturns struct {
		list *model.CommentList
		err  error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		mockReturns
		want    *model.CommentList
		wantErr bool
	}{
		{
			name: "When appropriate args given, ListComments returns CommentList and nil",
			fields: fields{
				m:       mock_query.NewMockDBManager(ctrl),
				repo:    mock_repository.NewMockCommentRepository(ctrl),
				service: mock_service.NewMockCommentService(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx:      context.Background(),
				threadID: model.ThreadValidIDForTest,
				limit:    20,
				cursor:   21,
			},
			mockReturns: mockReturns{
				list: &model.CommentList{
					Comments: testutil.GenerateCommentHelper(21, 40),
					HasNext:  true,
					Cursor:   41,
				},
				err: nil,
			},
			want: &model.CommentList{
				Comments: testutil.GenerateCommentHelper(21, 40),
				HasNext:  true,
				Cursor:   41,
			},
			wantErr: false,
		},
		{
			name: "When some error occurs at repository layer, ListComments returns nil and error",
			fields: fields{
				m:    mock_query.NewMockDBManager(ctrl),
				repo: mock_repository.NewMockCommentRepository(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx:      context.Background(),
				threadID: model.ThreadValidIDForTest,
				limit:    20,
				cursor:   21,
			},
			mockReturns: mockReturns{
				list: nil,
				err:  errors.New(model.ErrorMessageForTest),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, ok := tt.fields.repo.(*mock_repository.MockCommentRepository)
			if !ok {
				t.Fatal("failed to assert MockCommentRepository")
			}
			tr.EXPECT().ListComments(tt.args.ctx, tt.fields.m, tt.args.threadID, tt.args.limit, tt.args.cursor).Return(tt.mockReturns.list, tt.mockReturns.err)

			a := &commentService{
				m:        tt.fields.m,
				service:  tt.fields.service,
				repo:     tt.fields.repo,
				txCloser: tt.fields.txCloser,
			}
			got, err := a.ListComments(tt.args.ctx, tt.args.threadID, tt.args.limit, tt.args.cursor)
			if (err != nil) != tt.wantErr {
				t.Errorf("commentService.ListComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commentService.ListComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentService_GetComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testutil.SetFakeTime(time.Now())

	type fields struct {
		m        query.DBManager
		repo     repository.CommentRepository
		txCloser CloseTransaction
	}
	type args struct {
		ctx context.Context
		id  uint32
	}

	type mockReturns struct {
		comment *model.Comment
		err     error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		mockReturns
		want    *model.Comment
		wantErr bool
	}{
		{
			name: "When appropriate args given, GetComment returns Comment and nil",
			fields: fields{
				m:    mock_query.NewMockDBManager(ctrl),
				repo: mock_repository.NewMockCommentRepository(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			mockReturns: mockReturns{
				comment: &model.Comment{
					ID:       model.CommentInValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
				err: nil,
			},
			want: &model.Comment{
				ID:       model.CommentInValidIDForTest,
				ThreadID: model.ThreadValidIDForTest,
				User: &model.User{
					ID:   model.UserValidIDForTest,
					Name: model.UserNameForTest,
				},
				Content: model.CommentContentForTest,
			},
			wantErr: false,
		},
		{
			name: "When some error occurs at repository layer, GetComment returns nil and error",
			fields: fields{
				m:    mock_query.NewMockDBManager(ctrl),
				repo: mock_repository.NewMockCommentRepository(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				id:  model.CommentInValidIDForTest,
			},
			mockReturns: mockReturns{
				comment: nil,
				err:     errors.New(model.ErrorMessageForTest),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, ok := tt.fields.repo.(*mock_repository.MockCommentRepository)
			if !ok {
				t.Fatal("failed to assert MockCommentRepository")
			}
			tr.EXPECT().GetCommentByID(tt.args.ctx, tt.fields.m, tt.args.id).Return(tt.mockReturns.comment, tt.mockReturns.err)

			a := &commentService{
				m:        tt.fields.m,
				repo:     tt.fields.repo,
				txCloser: tt.fields.txCloser,
			}
			got, err := a.GetComment(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("commentService.GetComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commentService.GetComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commentService_CreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testutil.SetFakeTime(time.Now())

	type fields struct {
		m        query.DBManager
		service  service.CommentService
		repo     repository.CommentRepository
		txCloser CloseTransaction
	}
	type args struct {
		ctx   context.Context
		param *model.Comment
	}

	type mockArgsInsertComment struct {
		ctx   context.Context
		tx    query.DBManager
		param *model.Comment
	}

	type mockReturnsInsertComment struct {
		id  uint32
		err error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		mockArgsInsertComment
		mockReturnsInsertComment
		wantComment *model.Comment
		wantErr     bool
	}{
		{
			name: "When appropriate args given, CreateComment returns id and nil",
			fields: fields{
				m:       mock_query.NewMockDBManager(ctrl),
				repo:    mock_repository.NewMockCommentRepository(ctrl),
				service: mock_service.NewMockCommentService(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				param: &model.Comment{
					ID:       model.CommentInValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockArgsInsertComment: mockArgsInsertComment{
				ctx: context.Background(),
				param: &model.Comment{
					ID:       model.CommentInValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockReturnsInsertComment: mockReturnsInsertComment{
				id:  model.CommentValidIDForTest,
				err: nil,
			},
			wantComment: &model.Comment{
				ID:       model.CommentValidIDForTest,
				ThreadID: model.ThreadValidIDForTest,
				User: &model.User{
					ID:   model.UserValidIDForTest,
					Name: model.UserNameForTest,
				},
				Content: model.CommentContentForTest,
			},
			wantErr: false,
		},
		{
			name: "When some error occurs at repository layer, CreateComment returns nil and error",
			fields: fields{
				m:       mock_query.NewMockDBManager(ctrl),
				repo:    mock_repository.NewMockCommentRepository(ctrl),
				service: mock_service.NewMockCommentService(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockArgsInsertComment: mockArgsInsertComment{
				ctx: context.Background(),
				tx:  mock_query.NewMockDBManager(ctrl),
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockReturnsInsertComment: mockReturnsInsertComment{
				id:  model.CommentValidIDForTest,
				err: errors.New(model.ErrorMessageForTest),
			},
			wantComment: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, ok := tt.fields.m.(*mock_query.MockDBManager)
			if !ok {
				t.Fatal("failed to assert MockDBManager")
			}
			m.EXPECT().Begin().Return(mock_query.NewMockTxManager(ctrl), nil)

			if tt.mockArgsInsertComment.param != nil {
				tr, ok := tt.fields.repo.(*mock_repository.MockCommentRepository)
				if !ok {
					t.Fatal("failed to assert MockCommentRepository")
				}

				txM := mock_query.NewMockTxManager(ctrl)

				tr.EXPECT().InsertComment(tt.mockArgsInsertComment.ctx, txM, tt.args.param).Return(tt.mockReturnsInsertComment.id, tt.mockReturnsInsertComment.err)
			}

			a := &commentService{
				m:        tt.fields.m,
				repo:     tt.fields.repo,
				service:  tt.fields.service,
				txCloser: tt.fields.txCloser,
			}
			gotComment, err := a.CreateComment(tt.args.ctx, tt.args.param)
			if gotComment != nil {
				gotComment.CreatedAt = tt.wantComment.CreatedAt
				gotComment.UpdatedAt = tt.wantComment.UpdatedAt
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("commentService.CreateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotComment, tt.wantComment) {
				t.Errorf("commentService.CreateComment() = %v, want %v", gotComment, tt.wantComment)
			}
		})
	}
}

func Test_commentService_UpdateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testutil.SetFakeTime(time.Now())

	type fields struct {
		m        query.DBManager
		service  service.CommentService
		repo     repository.CommentRepository
		txCloser CloseTransaction
	}
	type args struct {
		ctx   context.Context
		id    uint32
		param *model.Comment
	}

	type mockArgsIsAlreadyExistID struct {
		ctx context.Context
		id  uint32
	}

	type mockReturnsIsAlreadyExistID struct {
		found bool
		err   error
	}

	type mockArgsUpdateComment struct {
		ctx   context.Context
		param *model.Comment
	}

	type mockReturnsUpdateComment struct {
		err error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		mockArgsIsAlreadyExistID
		mockReturnsIsAlreadyExistID
		mockArgsUpdateComment
		mockReturnsUpdateComment
		wantComment *model.Comment
		wantErr     bool
	}{
		{
			name: "When appropriate args given, UpdateComment returns Comment and err",
			fields: fields{
				m:       mock_query.NewMockDBManager(ctrl),
				service: mock_service.NewMockCommentService(ctrl),
				repo:    mock_repository.NewMockCommentRepository(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockArgsIsAlreadyExistID: mockArgsIsAlreadyExistID{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
			},
			mockReturnsIsAlreadyExistID: mockReturnsIsAlreadyExistID{
				found: true,
				err:   nil,
			},
			mockArgsUpdateComment: mockArgsUpdateComment{
				ctx: context.Background(),
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockReturnsUpdateComment: mockReturnsUpdateComment{
				err: nil,
			},
			wantComment: &model.Comment{
				ID:       model.CommentValidIDForTest,
				ThreadID: model.ThreadValidIDForTest,
				User: &model.User{
					ID:   model.UserValidIDForTest,
					Name: model.UserNameForTest,
				},
				Content: model.CommentContentForTest,
			},
			wantErr: false,
		},
		{
			name: "When given id has not existed, UpdateComment returns nil and error",
			fields: fields{
				m:       mock_query.NewMockDBManager(ctrl),
				service: mock_service.NewMockCommentService(ctrl),
				repo:    mock_repository.NewMockCommentRepository(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockArgsIsAlreadyExistID: mockArgsIsAlreadyExistID{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
			},
			mockReturnsIsAlreadyExistID: mockReturnsIsAlreadyExistID{
				found: false,
				err:   nil,
			},
			mockArgsUpdateComment: mockArgsUpdateComment{
				param: nil,
			},
			wantComment: nil,
			wantErr:     true,
		},
		{
			name: "When some error occurs at repository layer, UpdateComment returns nil and error",
			fields: fields{
				m:       mock_query.NewMockDBManager(ctrl),
				service: mock_service.NewMockCommentService(ctrl),
				repo:    mock_repository.NewMockCommentRepository(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockArgsIsAlreadyExistID: mockArgsIsAlreadyExistID{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
			},
			mockReturnsIsAlreadyExistID: mockReturnsIsAlreadyExistID{
				found: true,
				err:   nil,
			},
			mockArgsUpdateComment: mockArgsUpdateComment{
				ctx: context.Background(),
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockReturnsUpdateComment: mockReturnsUpdateComment{
				err: errors.New(model.ErrorMessageForTest),
			},
			wantComment: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, ok := tt.fields.m.(*mock_query.MockDBManager)
			if !ok {
				t.Fatal("failed to assert MockDBManager")
			}
			m.EXPECT().Begin().Return(mock_query.NewMockTxManager(ctrl), nil)

			ts, ok := tt.fields.service.(*mock_service.MockCommentService)
			if !ok {
				t.Fatal("failed to assert MockCommentService")
			}

			ts.EXPECT().IsAlreadyExistID(tt.mockArgsIsAlreadyExistID.ctx, gomock.Any(), tt.mockArgsIsAlreadyExistID.id).Return(tt.mockReturnsIsAlreadyExistID.found, tt.mockReturnsIsAlreadyExistID.err)

			if tt.mockArgsUpdateComment.param != nil {

				tr, ok := tt.fields.repo.(*mock_repository.MockCommentRepository)
				if !ok {
					t.Fatal("failed to assert MockCommentRepository")
				}

				txM := mock_query.NewMockTxManager(ctrl)

				tr.EXPECT().UpdateComment(tt.mockArgsUpdateComment.ctx, txM, tt.args.id, tt.args.param).Return(tt.mockReturnsUpdateComment.err)

			}

			a := &commentService{
				m:        tt.fields.m,
				service:  tt.fields.service,
				repo:     tt.fields.repo,
				txCloser: tt.fields.txCloser,
			}

			gotComment, err := a.UpdateComment(tt.args.ctx, tt.args.id, tt.args.param)
			if gotComment != nil {
				gotComment.UpdatedAt = tt.wantComment.UpdatedAt
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("commentService.UpdateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotComment, tt.wantComment) {
				t.Errorf("commentService.UpdateComment() = %+v, want %+v", gotComment, tt.wantComment)
			}
		})
	}
}

func Test_commentService_DeleteComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testutil.SetFakeTime(time.Now())

	type fields struct {
		m        query.DBManager
		service  service.CommentService
		repo     repository.CommentRepository
		txCloser CloseTransaction
	}
	type args struct {
		ctx   context.Context
		id    uint32
		param *model.Comment
	}

	type mockArgsIsAlreadyExistID struct {
		ctx context.Context
		id  uint32
	}

	type mockReturnsIsAlreadyExistID struct {
		found bool
		err   error
	}

	type mockReturnsDeleteComment struct {
		err error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		mockArgsIsAlreadyExistID
		mockReturnsIsAlreadyExistID
		mockReturnsDeleteComment
		wantComment *model.Comment
		wantErr     bool
	}{
		{
			name: "When appropriate args given, DeleteComment returns Comment and err",
			fields: fields{
				m:       mock_query.NewMockDBManager(ctrl),
				service: mock_service.NewMockCommentService(ctrl),
				repo:    mock_repository.NewMockCommentRepository(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockArgsIsAlreadyExistID: mockArgsIsAlreadyExistID{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
			},
			mockReturnsIsAlreadyExistID: mockReturnsIsAlreadyExistID{
				found: true,
				err:   nil,
			},
			mockReturnsDeleteComment: mockReturnsDeleteComment{
				err: nil,
			},
			wantComment: &model.Comment{
				ID:       model.CommentValidIDForTest,
				ThreadID: model.ThreadValidIDForTest,
				User: &model.User{
					ID:   model.UserValidIDForTest,
					Name: model.UserNameForTest,
				},
				Content: model.CommentContentForTest,
			},
			wantErr: false,
		},
		{
			name: "When given id has not existed, DeleteComment returns nil and error",
			fields: fields{
				m:       mock_query.NewMockDBManager(ctrl),
				service: mock_service.NewMockCommentService(ctrl),
				repo:    mock_repository.NewMockCommentRepository(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockArgsIsAlreadyExistID: mockArgsIsAlreadyExistID{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
			},
			mockReturnsIsAlreadyExistID: mockReturnsIsAlreadyExistID{
				found: false,
				err:   nil,
			},
			wantComment: nil,
			wantErr:     true,
		},
		{
			name: "When some error occurs at repository layer, DeleteComment returns nil and error",
			fields: fields{
				m:       mock_query.NewMockDBManager(ctrl),
				service: mock_service.NewMockCommentService(ctrl),
				repo:    mock_repository.NewMockCommentRepository(ctrl),
				txCloser: func(tx query.TxManager, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
				param: &model.Comment{
					ID:       model.CommentValidIDForTest,
					ThreadID: model.ThreadValidIDForTest,
					User: &model.User{
						ID:   model.UserValidIDForTest,
						Name: model.UserNameForTest,
					},
					Content: model.CommentContentForTest,
				},
			},
			mockArgsIsAlreadyExistID: mockArgsIsAlreadyExistID{
				ctx: context.Background(),
				id:  model.CommentValidIDForTest,
			},
			mockReturnsIsAlreadyExistID: mockReturnsIsAlreadyExistID{
				found: true,
				err:   nil,
			},
			mockReturnsDeleteComment: mockReturnsDeleteComment{
				err: errors.New(model.ErrorMessageForTest),
			},
			wantComment: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, ok := tt.fields.m.(*mock_query.MockDBManager)
			if !ok {
				t.Fatal("failed to assert MockDBManager")
			}
			m.EXPECT().Begin().Return(mock_query.NewMockTxManager(ctrl), nil)

			ts, ok := tt.fields.service.(*mock_service.MockCommentService)
			if !ok {
				t.Fatal("failed to assert MockCommentService")
			}

			ts.EXPECT().IsAlreadyExistID(tt.mockArgsIsAlreadyExistID.ctx, gomock.Any(), tt.mockArgsIsAlreadyExistID.id).Return(tt.mockReturnsIsAlreadyExistID.found, tt.mockReturnsIsAlreadyExistID.err)

			if tt.mockReturnsIsAlreadyExistID.found {
				tr, ok := tt.fields.repo.(*mock_repository.MockCommentRepository)
				if !ok {
					t.Fatal("failed to assert MockCommentRepository")
				}

				txM := mock_query.NewMockTxManager(ctrl)

				tr.EXPECT().DeleteComment(tt.args.ctx, txM, tt.args.id).Return(tt.mockReturnsDeleteComment.err)
			}

			a := &commentService{
				m:        tt.fields.m,
				service:  tt.fields.service,
				repo:     tt.fields.repo,
				txCloser: tt.fields.txCloser,
			}

			err := a.DeleteComment(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("commentService.DeleteComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
