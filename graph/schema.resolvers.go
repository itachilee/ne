package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/itachilee/gin/global"
	"github.com/itachilee/gin/graph/generated"
	"github.com/itachilee/gin/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	collection := global.MgoCli.Database("blog").Collection("todo")

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

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	findOption := options.Find()
	findOption.SetLimit(5)
	collection := global.MgoCli.Database("blog").Collection("todo")
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
	collection := global.MgoCli.Database("blog").Collection("todo")

	todo := model.Todo{}
	err := collection.FindOne(context.TODO(), findOption).Decode(&todo)
	if err != nil {
		fmt.Println(id)
		fmt.Println(err)
	}
	return &todo, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
