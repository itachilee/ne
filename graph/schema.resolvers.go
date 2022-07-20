package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/itachilee/gin/graph/generated"
	"github.com/itachilee/gin/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	collection := r.MgoCli.Database("blog").Collection("todo")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{
		{Key: "userId", Value: input.UserID},
		{Key: "text", Value: input.Text}})
	if err != nil {
		fmt.Print(err)
	}
	id := res.InsertedID.(primitive.ObjectID)

	todo := &model.Todo{
		ID:   id.String(), // fmt.Sprint((rand.Int())),
		Text: input.Text,
	}

	return todo, nil
}

type T struct {
	ReaderSeeker io.ReadSeeker
	L            int64
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

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	findOption := options.Find()
	findOption.SetLimit(5)
	collection := r.MgoCli.Database("blog").Collection("todo")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOption)
	if err != nil {
		fmt.Print(err)
	}
	var todos []*model.Todo
	for cur.Next(context.TODO()) {
		var todo model.Todo
		err = cur.Decode(&todo)

		if err != nil {
			fmt.Print(err)
		}
		todos = append(todos, &todo)
	}
	if err := cur.Err(); err != nil {
		fmt.Print(err)
	}
	cur.Close(context.TODO())
	return todos, nil
	// panic(fmt.Errorf("not implemented"))
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	findOption := bson.M{"_id": id}
	collection := r.MgoCli.Database("blog").Collection("todo")

	todo := model.Todo{}
	err := collection.FindOne(context.TODO(), findOption).Decode(&todo)
	if err != nil {
		fmt.Println(id)
		fmt.Println(err)
	}
	return &todo, err
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
