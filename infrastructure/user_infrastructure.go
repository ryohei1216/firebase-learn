package infrastructure

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/ryohei1216/firebase-learn/domain/entity/user"
	"github.com/ryohei1216/firebase-learn/domain/repository"
	"google.golang.org/api/iterator"
)

type userRepository struct {
	sc *firestore.Client
}

func NewUserRepository(storeClient *firestore.Client) repository.UserRepository {
	return &userRepository{
		sc: storeClient,
	}
}

func (ur userRepository) Create(ctx context.Context, u user.User) (*user.User, error) {
	ref, result, err := ur.sc.Collection("users").Add(ctx, u)
	if err != nil {
		log.Printf("failed to create user: %v", err)
	}
	log.Println("ref", ref)
	log.Println("result", result)

	return nil, nil
}

func (ur userRepository) Get(ctx context.Context, uid string) (*user.User, error) {
	var user user.User
	
	iter := ur.sc.Collection("users").Where("ID", "==", uid).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		jsonData, err := json.Marshal(doc.Data())
		if err != nil {
			log.Println(err)
			return nil, err
		}

		err = json.Unmarshal(jsonData, &user)
		if err != nil {
			log.Println(err)
			return nil, err
	  }
	}
	
	return &user, nil
}
