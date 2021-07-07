package main

import (
	"context"
	"encoding/json"
	"feedPosts/pkg/dtos"
	"feedPosts/pkg/models"
	"feedPosts/tracer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"io"
	"net/http"
)

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	span := tracer.StartSpanFromContext(ctx, "renderJSON")
	defer span.Finish()

	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func decodeBody(ctx context.Context, r io.Reader) (*dtos.FeedPostDTO, error) {
	span := tracer.StartSpanFromContext(ctx, "decodeBody")
	defer span.Finish()

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt dtos.FeedPostDTO
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func createPost(ctx context.Context, rt *dtos.FeedPostDTO, id primitive.ObjectID, app application) (mongo.InsertOneResult, error) {
	span := tracer.StartSpanFromContext(ctx, "CreatePost")
	defer span.Finish()
	listTagged := taggedUsersToPrimitiveObject(ctx, *rt)
	var post = models.Post{
		User:        id,
		DateTime:    time.Now(),
		Tagged:      listTagged,
		Description: rt.Description,
		Hashtags:    parseHashTags(rt.Hashtags),
		Location:    rt.Location,
		Blocked:     false,
	}
	var feedPost = models.FeedPost{
		Post:     post,
		Likes:    []primitive.ObjectID{},
		Dislikes: []primitive.ObjectID{},
		Comments: []models.Comment{},
	}

	insertResult, err := app.feedPosts.Insert(feedPost)


	return *insertResult, err
}

func taggedUsersToPrimitiveObject(ctx context.Context, m dtos.FeedPostDTO) []primitive.ObjectID {
	span := tracer.StartSpanFromContext(ctx, "taggedUsersToPrimitiveObject")
	defer span.Finish()
	listTagged := []primitive.ObjectID{}
	for _, tag := range m.Tagged {
		primitiveTag, _ := primitive.ObjectIDFromHex(tag)

		listTagged = append(listTagged, primitiveTag)
	}
	return listTagged
}