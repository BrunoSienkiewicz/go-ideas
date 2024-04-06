package repository

import (
	db "github.com/BrunoSienkiewicz/go_ideas/internal/db"
	storage "github.com/BrunoSienkiewicz/go_ideas/internal/storage"
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

type IdeaRepository struct {
	categoryStorage  storage.Storage[types.DbCategory]
	attributeStorage storage.Storage[types.DbAttribute]
	ideaStorage      storage.Storage[types.DbIdea]
}

func NewIdeaRepository(postgres *db.Postgres) *IdeaRepository {
	return &IdeaRepository{
		categoryStorage:  storage.NewCategoryStorage(postgres),
		attributeStorage: storage.NewAttributeStorage(postgres),
		ideaStorage:      storage.NewIdeaStorage(postgres),
	}
}

func (r *IdeaRepository) handleCategory(name string) (*types.DbCategory, error) {
	categories, err := r.categoryStorage.GetObjectsByField("name", name)
	if err != nil {
		return nil, err
	}

	if categories == nil {
		dbCategory := &types.DbCategory{
			Name: name,
		}

		err = r.categoryStorage.AddObject(dbCategory)
		if err != nil {
			return nil, err
		}

		categories, err = r.categoryStorage.GetObjectsByField("name", name)
		if err != nil {
			return nil, err
		}
	}

	return categories[0], nil
}

func (r *IdeaRepository) GetIdea(id int) (*types.Idea, error) {
	dbIdea, err := r.ideaStorage.GetObject(id)
	if err != nil {
		return nil, err //NotFoundError{Err: "Idea with ID: " + string(id) + " Not Found"}
	}

	ideaCategory, err := r.categoryStorage.GetObject(dbIdea.Category_id)
	ideaDbAttributes, err := r.attributeStorage.GetObjectsByField("idea_id", string(id))

	var ideaAttributes []types.Attribute
	for _, dbAttribute := range ideaDbAttributes {
		ideaAttributes = append(ideaAttributes, types.Attribute{
			Name:  dbAttribute.Name,
			Value: dbAttribute.Value,
		})
	}

	idea := &types.Idea{
		ID:         dbIdea.Idea_id,
		Name:       dbIdea.Name,
		Category:   ideaCategory.Name,
		Attributes: ideaAttributes,
	}

	return idea, nil
}

func (r *IdeaRepository) GetAllIdeas() ([]*types.Idea, error) {
	dbIdeas, err := r.ideaStorage.GetAllObjects()
	if err != nil {
		return nil, NotFoundError{Err: "No Ideas Found"}
	}

	var ideas []*types.Idea
	for _, dbIdea := range dbIdeas {
		idea, err := r.GetIdea(dbIdea.Idea_id)
		if err != nil {
			return nil, err
		}

		ideas = append(ideas, idea)
	}

	return ideas, nil
}

func (r *IdeaRepository) AddIdea(idea *types.Idea) (*types.Idea, error) {
	category, err := r.handleCategory(idea.Category)

	dbIdea := &types.DbIdea{
		Name:        idea.Name,
		Category_id: category.Category_id,
	}

	err = r.ideaStorage.AddObject(dbIdea)
	if err != nil {
		return nil, err
	}

	for _, attribute := range idea.Attributes {
		dbAttribute := &types.DbAttribute{
			Idea_id: dbIdea.Idea_id,
			Name:    attribute.Name,
			Value:   attribute.Value,
		}

		err := r.attributeStorage.AddObject(dbAttribute)
		if err != nil {
			return nil, err
		}
	}

	return idea, nil
}

func (r *IdeaRepository) UpdateIdea(idea *types.Idea) error {
	dbIdea, err := r.ideaStorage.GetObject(idea.ID)
	if err != nil {
		return NotFoundError{Err: "Idea with ID: " + string(idea.ID) + " Not Found"}
	}

	category, err := r.handleCategory(idea.Category)

	dbIdea.Name = idea.Name
	dbIdea.Category_id = category.Category_id

	err = r.ideaStorage.UpdateObject(dbIdea)
	if err != nil {
		return err
	}

	// TODO: Update attributes

	return nil
}

func (r *IdeaRepository) DeleteIdea(id int) error {
	_, err := r.ideaStorage.GetObject(id)
	if err != nil {
		return NotFoundError{Err: "Idea with ID: " + string(id) + " Not Found"}
	}

	err = r.ideaStorage.DeleteObject(id)
	if err != nil {
		return err
	}

	return nil
}
