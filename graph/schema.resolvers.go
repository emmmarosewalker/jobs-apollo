package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jobs-apollo/graph/generated"
	"jobs-apollo/graph/model"
	"log"
	"net/http"
)

func (r *mutationResolver) AddListing(ctx context.Context, input model.ListingInput) (*model.Listing, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListingsByCompany(ctx context.Context, companyID string) ([]*model.Listing, error) {
	resp, err := http.Get("http://localhost:7000/listings/company/" + companyID)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if len(body) == 0 {
		return nil, nil
	}

	var listings []*model.Listing
	err = json.Unmarshal(body, &listings)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listings, err
}

func (r *queryResolver) Listings(ctx context.Context) ([]*model.Listing, error) {
	resp, err := http.Get("http://localhost:7000/listings")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if len(body) == 0 {
		return nil, nil
	}

	var listings []*model.Listing
	err = json.Unmarshal(body, &listings)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listings, err
}

func (r *queryResolver) ListingByID(ctx context.Context, id int) (*model.Listing, error) {
	endpoint := fmt.Sprintf("%s%d", "http://localhost:7000/listings/", id)
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if len(body) == 0 {
		return nil, nil
	}

	var listing *model.Listing
	err = json.Unmarshal(body, &listing)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listing, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
