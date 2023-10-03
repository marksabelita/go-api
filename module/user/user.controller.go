package user

import (
	"context"
	"go-api/config"
	user_model "go-api/module/user/model"
	user_response "go-api/module/user/response"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")

// @Summary Lists all users details.
// @Description Lists all users details.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} []user_model.User
// @Param        name    query     string  false  "name"  
// @Router /users [get]
func GetUser(c *fiber.Ctx) error {
    query := bson.M{}
    name := c.Query("name")

    if name != "" { query["name"] = name  }
    
    users, err := FindService(query)

    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(user_response.ErrorResponse{Message: "error"})
    }
   
    return c.Status(http.StatusOK).JSON(
        users,
    )
}

// @Summary Display user details
// @Description Display user details
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} user_model.User
// @Param        id   path      int  false  "Account ID"

// @Router /users/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	var user user_model.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
	if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(user_response.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(user_response.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

func EditUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    userId := c.Params("userId")
    var user user_model.User
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex(userId)

    //validate the request body
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(user_response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    //use the validator library to validate required fields
		validate := validator.New()
    if validationErr := validate.Struct(&user); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(user_response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }

    update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}

    result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(user_response.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
    //get updated user details
    var updatedUser user_model.User
    if result.MatchedCount == 1 {
        err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

        if err != nil {
            return c.Status(http.StatusInternalServerError).JSON(user_response.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }
    }

    return c.Status(http.StatusOK).JSON(user_response.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedUser}})
}

func CreateUser(c *fiber.Ctx) error {
	var user user_model.User
	ctx, cancel := context.WithTimeout(context.Background(), config.DEFAULT_TIMEOUT * time.Second)
		defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(user_response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	validate := validator.New()

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(user_response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUser := user_model.User{
		Id: primitive.NewObjectID(),
		Name: user.Name,
		Location: "test",
		Title: "test",
	}

	result, err := userCollection.InsertOne(ctx, newUser)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(user_response.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(user_response.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    userId := c.Params("userId")
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex(userId)

    result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(user_response.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    if result.DeletedCount < 1 {
        return c.Status(http.StatusNotFound).JSON(
            user_response.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
        )
    }

    return c.Status(http.StatusOK).JSON(
        user_response.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
    )
}