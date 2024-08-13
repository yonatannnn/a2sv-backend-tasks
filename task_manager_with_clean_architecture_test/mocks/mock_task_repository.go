package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SingleResultMock is a custom mock for mongo.SingleResult
type SingleResultMock struct {
    Result interface{}
    Err    error
}

// Decode mocks the Decode method of mongo.SingleResult
// Decode mocks the Decode method of mongo.SingleResult
func (s *SingleResultMock) Decode(v interface{}) error {
    if s.Err != nil {
        return s.Err
    }
    // Marshal the Result into BSON bytes and then Unmarshal into the provided interface
    bsonBytes, _ := bson.Marshal(s.Result)
    return bson.Unmarshal(bsonBytes, v)
}


// MockTaskCollection is the mock collection
type MockTaskCollection struct {
    mock.Mock
}

func (m *MockTaskCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	result := args.Get(0).(*mongo.SingleResult)
	return result
}


func (m *MockTaskCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockTaskCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockTaskCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTaskCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
    args := m.Called(ctx, filter)
    return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockTaskCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
    args := m.Called(ctx, filter)
    return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}




