package service

import (
	"context"
	"errors"
	"project/internal/auth"
	"project/internal/models"
	"project/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_ViewJobById(t *testing.T) {
	type args struct {
		ctx context.Context
		jid uint64
	}
	tests := []struct {
		name string
		// s       *Service
		args             args
		want             models.Jobs
		wantErr          bool
		mockRepoResponse func() (models.Jobs, error)
	}{
		{name: "error",
			want: models.Jobs{},
			args: args{
				ctx: context.Background(),
				jid: 5,
			},
			wantErr: true,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, errors.New("test error")
			},
		},
		{name: "success",
			want: models.Jobs{
				Company: models.Company{
					Name:     "tcs",
					Location: "bang",
					Field:    "software",
				},
				Cid:          2,
				Name:         "developer",
				Salary:       "30000",
				NoticePeriod: "3 weeks",
			},
			args: args{
				ctx: context.Background(),
				jid: 15,
			},
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Company: models.Company{
						Name:     "tcs",
						Location: "bang",
						Field:    "software",
					},
					Cid:          2,
					Name:         "developer",
					Salary:       "30000",
					NoticePeriod: "3 weeks",
				}, nil
			},
		},
		{
			name: "invalid id",
			want: models.Jobs{},
			args: args{
				ctx: context.Background(),
				jid: 5,
			},
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, nil
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobDetailsBy(tt.args.ctx, tt.args.jid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewJobById(tt.args.ctx, tt.args.jid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJobById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJobById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewAllJobs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		// s       *Service
		args             args
		want             []models.Jobs
		wantErr          bool
		mockRepoResponse func() ([]models.Jobs, error)
	}{
		// {name: "error",
		// 	want: []models.Jobs{},
		// 	args: args{
		// 		ctx: context.Background(),
		// 	},
		// 	wantErr: true,
		// 	mockRepoResponse: func() ([]models.Jobs, error) {
		// 		return []models.Jobs{}, errors.New("test error")
		// 	},
		// },
		{
			name: "success",
			want: []models.Jobs{
				{
					Cid:          2,
					Name:         "tcs",
					Salary:       "30000",
					NoticePeriod: "3 weeks",
				},
				{Cid: 2,
					Name:         "developer",
					Salary:       "30000",
					NoticePeriod: "3 weeks",
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return []models.Jobs{
					{
						Cid:          2,
						Name:         "tcs",
						Salary:       "30000",
						NoticePeriod: "3 weeks",
					},
					{
						Cid:          2,
						Name:         "developer",
						Salary:       "30000",
						NoticePeriod: "3 weeks",
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)

			mockRepo.EXPECT().FindAllJobs(tt.args.ctx).Return(tt.mockRepoResponse()).AnyTimes()

			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewAllJobs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewJob(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name string
		// s       *Service
		args             args
		want             []models.Jobs
		wantErr          bool
		mockRepoResponse func() ([]models.Jobs, error)
	}{
		{
			name: "success",
			want: []models.Jobs{
				// {
				// 	Cid:          2,
				// 	Name:         "tcs",
				// 	Salary:       "30000",
				// 	NoticePeriod: "3 weeks",
				// },
				{Cid: 2,
					Name:         "assosiate",
					Salary:       "50000",
					NoticePeriod: "3 days",
				},
			},
			args: args{
				ctx: context.Background(),
				cid: 4,
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return []models.Jobs{
					{
						Cid:          2,
						Name:         "assosiate",
						Salary:       "50000",
						NoticePeriod: "3 days",
					},
				}, nil
			},
		},
	}
	// {name: "failure",
	// 	args: args{
	// 		ctx: context.Background(),
	// 		cid: 0,
	// 	},
	// 	want: []models.Jobs{},
	// 	},
	// 	wantErr: true,
	// 	mockRepoResponse: func() ([]models.Jobs, error) {
	// 		return []models.Jobs{}, errors.New("id is zero")
	// 	},

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().FindJob(tt.args.ctx, tt.args.cid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewJob(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
