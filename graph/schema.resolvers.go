package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/itachilee/gin/graph/generated"
	"github.com/itachilee/gin/graph/model"
	"github.com/itachilee/gin/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUser) (*model.User, error) {
	user := models.ZUsers{
		Name:            input.Name,
		Email:           input.Email,
		Password:        input.Password,
		Status:          1,
		EmailVerifiedAt: 0_0,
	}
	err := r.Db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	res := &model.User{
		ID:       int(user.ID),
		Email:    &user.Email,
		Password: &user.Password,
	}
	// return user.Create(r.Db)
	return res, nil
}

// SingleUpload is the resolver for the singleUpload field.
func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	// panic(fmt.Errorf("not implemented"))
	content, err := io.ReadAll(file.File)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(content)
	_, err = r.CosCli.Object.Put(ctx, file.Filename, b, nil)
	if err != nil {
		panic(err)
	}
	// dump.V(resp)
	return &model.File{
		ID:      1,
		Name:    file.Filename,
		Content: string(content),
	}, nil
}

// SingleUploadWithPayload is the resolver for the singleUploadWithPayload field.
func (r *mutationResolver) SingleUploadWithPayload(ctx context.Context, req model.UploadFile) (*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}

// MultipleUpload is the resolver for the multipleUpload field.
func (r *mutationResolver) MultipleUpload(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}

// MultipleUploadWithPayload is the resolver for the multipleUploadWithPayload field.
func (r *mutationResolver) MultipleUploadWithPayload(ctx context.Context, req []*model.UploadFile) ([]*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {

	return nil, errors.New("kfc crazy thursday need $50")
	// mgr := models.ZUsersMgr(r.Db)
	// users, err := mgr.Gets()
	// if err != nil {
	// 	return nil, err
	// }
	// res := make([]*model.User, 0)
	// for _, u := range users {
	// 	t := &model.User{
	// 		ID:       int(u.ID),
	// 		Email:    &u.Email,
	// 		Name:     &u.Name,
	// 		Password: &u.Password,
	// 	}
	// 	res = append(res, t)
	// }
	// return res, nil

}

// Empty is the resolver for the empty field.
func (r *queryResolver) Empty(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
