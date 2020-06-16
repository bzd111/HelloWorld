package gin_graphql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
	"gin_graphql/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var db = GetInstance()

type Resolver struct{}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	// DisplayName string `json:"display_name"`
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*models.Todo, error) {
	todo := models.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}
	db.Create(&todo)
	return &todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input EditTodo) (*models.Todo, error) {
	todo := models.Todo{ID: input.ID}
	db.First(&todo)
	todo.Text = input.Text
	db.Model(&models.Todo{}).Update(&todo)
	return &todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input int) (*models.Todo, error) {
	todo := models.Todo{ID: input}
	db.First(&todo)
	db.Delete(&todo)
	return &todo, nil
}

func (r *mutationResolver) UserLogin(ctx context.Context, name string, password string) (*models.User, error) {
	user := models.User{}
	db.Where("name = ?", name).First(&user)
	if user.Name == "" {
		// return nil, errors.Wrap("try to insert post got error")
		return nil, errors.New("no such user name")
	}
	// hashedPassword, _ := GeneratePasswordHash([]byte("123"))
	// fmt.Println(string(hashedPassword))
	// if !ValidatePasswordHash([]byte(user.Password), []byte(password)) {
	//     return nil, errors.New("user password isn't correct")
	// }

	// uc := &UserClaims{
	//     StandardClaims: jwt.StandardClaims{
	//         Subject:   string(user.ID),
	//         IssuedAt:  utils.Clock2.GetUTCNow().Unix(),
	//         ExpiresAt: utils.Clock.GetUTCNow().Add(7 * 24 * time.Hour).Unix(),
	//     },
	//     Username: user.Name,
	//     // DisplayName: user.Username,
	// }
	// if err := auth.SetLoginCookie(ctx, uc); err != nil {
	//     return nil, errors.Wrap(err, "try to set cookies got error")
	/* } */

	return &user, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*models.User, error) {
	user := models.User{
		Name: input.Name,
	}
	db.Create(&user)
	return &user, nil
}

func (r *queryResolver) Todos(ctx context.Context, page *Pagination) ([]*models.Todo, error) {
	var todos []*models.Todo
	db.Preload("User").First(&todos).Limit(page.Size).Offset(page.Page)
	fmt.Println(todos[0].User)
	return todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	db.Find(&users)
	return users, nil
}

func (r *queryResolver) Todo(ctx context.Context, input *FetchTodo) (*models.Todo, error) {
	var todo models.Todo
	// todo.Done = false
	// todo.Text = "123123"
	db.Preload("User").First(&todo, input.ID)
	return &todo, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
