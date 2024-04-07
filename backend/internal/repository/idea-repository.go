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

func (r *IdeaRepository) convertToDbObject(obj *types.Idea) types.DbIdea {
	category_id, err := r.categoryStorage.GetObjectsByField("name", obj.Category)
	if err != nil {
		if _, ok := err.(storage.NotFoundError); ok {
			category_id[0], err = r.categoryStorage.AddObject(&types.DbCategory{Name: obj.Category})
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	return types.DbIdea{
		Idea_id:     obj.ID,
		Name:        obj.Name,
		Category_id: category_id[0].Category_id,
	}
}

func (r *IdeaRepository) convertFromDbObject(obj *types.DbIdea) types.Idea {
	category, err := r.categoryStorage.GetObject(obj.Category_id)
	if err != nil {
		panic(err)
	}

	attributes, err := r.attributeStorage.GetObjectsByField("idea_id", string(obj.Idea_id))
	if err != nil {
		panic(err)
	}

	var ideaAttributes []types.Attribute
	for _, attribute := range attributes {
		ideaAttributes = append(ideaAttributes, types.Attribute{
			ID:     attribute.Attribute_id,
			IdeaId: attribute.Idea_id,
			Name:   attribute.Name,
			Value:  attribute.Value,
		})
	}

	return types.Idea{
		ID:         obj.Idea_id,
		Name:       obj.Name,
		Category:   category.Name,
		Attributes: ideaAttributes,
	}
}

func (r *IdeaRepository) GetObject(id int) (*types.Idea, error) {
	dbIdea, err := r.ideaStorage.GetObject(id)
	if err != nil {
		return nil, NotFoundError{Err: "Idea with ID: " + string(id) + " Not Found"}
	}

	idea := r.convertFromDbObject(dbIdea)

	return &idea, nil
}

func (r *IdeaRepository) GetAllObjects() ([]*types.Idea, error) {
	dbIdeas, err := r.ideaStorage.GetAllObjects()
	if err != nil {
		return nil, NotFoundError{Err: "No Ideas Found"}
	}

	var ideas []*types.Idea
	for _, dbIdea := range dbIdeas {
		idea := r.convertFromDbObject(dbIdea)

		ideas = append(ideas, &idea)
	}

	return ideas, nil
}

func (r *IdeaRepository) AddObject(idea *types.Idea) (*types.Idea, error) {
	_dbIdea := r.convertToDbObject(idea)
	dbIdea, err := r.ideaStorage.AddObject(&_dbIdea)
	if err != nil {
		return nil, err
	}

	for _, attribute := range idea.Attributes {
		attribute.IdeaId = dbIdea.Idea_id
		dbAttribute, err := r.attributeStorage.AddObject(&types.DbAttribute{
			Name:    attribute.Name,
			Value:   attribute.Value,
			Idea_id: attribute.IdeaId,
		})
		if err != nil {
			return nil, err
		}
		attribute.ID = dbAttribute.Attribute_id
	}

	idea.ID = dbIdea.Idea_id
	return idea, nil
}

func (r *IdeaRepository) UpdateObject(idea *types.Idea) (*types.Idea, error) {
	_dbIdea := r.convertToDbObject(idea)
	_, err := r.ideaStorage.UpdateObject(&_dbIdea)
	if err != nil {
		return nil, err
	}

	for _, attribute := range idea.Attributes {
		dbAttribute, err := r.attributeStorage.GetObject(attribute.ID)
		if err != nil {
			return nil, err
		}

		dbAttribute.Name = attribute.Name
		dbAttribute.Value = attribute.Value
		_, err = r.attributeStorage.UpdateObject(dbAttribute)
		if err != nil {
			return nil, err
		}
	}

	return idea, nil
}

func (r *IdeaRepository) DeleteObject(id int) error {
	err := r.ideaStorage.DeleteObject(id)
	if err != nil {
		return err
	}

	attributes, err := r.attributeStorage.GetObjectsByField("idea_id", string(id))
	if err != nil {
		return err
	}

	for _, attribute := range attributes {
		err := r.attributeStorage.DeleteObject(attribute.Attribute_id)
		if err != nil {
			return err
		}
	}

	return nil
}
