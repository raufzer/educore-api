package repository

import (
    "context"
    "educore-api/internal/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
    collection *mongo.Collection
}

// Create a new user in the database
func (r *MongoUserRepository) Create(user *models.User) error {
    _, err := r.collection.InsertOne(context.Background(), user)
    return err
}

// Get a user by name from the database
func (r *MongoUserRepository) GetByName(name string) (*models.User, error) {
    var user models.User
    err := r.collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&user)
    return &user, err
}

// Get all users from the database
func (r *MongoUserRepository) GetAll() ([]*models.User, error) {
    var users []*models.User
    cursor, err := r.collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, &user)
    }
    return users, nil
}

// Update an existing user in the database
func (r *MongoUserRepository) Update(user *models.User) error {
    _, err := r.collection.UpdateOne(context.Background(), bson.M{"name": user.Name}, bson.M{"$set": user})
    return err
}

// Delete a user from the database by name
func (r *MongoUserRepository) Delete(name string) error {
    _, err := r.collection.DeleteOne(context.Background(), bson.M{"name": name})
    return err
}
