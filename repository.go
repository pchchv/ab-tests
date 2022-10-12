package main

import "errors"

type Repository interface {
	Create(todo *Hypothesis)
	GetAll() []*Hypothesis
	GetById(id int) (h *Hypothesis, err error)
	GetByTitle(key string) (h *Hypothesis, err error)
	Update(*Hypothesis) (err error)
	DeleteAll()
	Delete(id int) (err error)
}

type InMemoryRepository struct {
	Hypothesis []*Hypothesis
	nextId     int
}

func NewInMemoryRepository() Repository {
	h := new(InMemoryRepository)
	h.Hypothesis = make([]*Hypothesis, 0)
	h.nextId = 1
	return h
}

func (r *InMemoryRepository) Create(hypothesis *Hypothesis) {
	hypothesis.Id = r.nextId
	r.Hypothesis = append(r.Hypothesis, hypothesis)
	r.nextId++
}

func (r *InMemoryRepository) GetAll() []*Hypothesis {
	return r.Hypothesis
}

func (r *InMemoryRepository) DeleteAll() {
	r.Hypothesis = make([]*Hypothesis, 0)
}

func (r *InMemoryRepository) GetById(id int) (h *Hypothesis, err error) {
	for _, h = range r.Hypothesis {
		if h.Id == id {
			return h, nil
		}
	}
	return nil, errors.New("hypothesis not found")
}

func (r *InMemoryRepository) GetByTitle(key string) (h *Hypothesis, err error) {
	for _, h = range r.Hypothesis {
		if h.Key == key {
			return h, nil
		}
	}
	return nil, errors.New("hypothesis not found")
}

func (r *InMemoryRepository) Delete(id int) (err error) {
	for i, h := range r.Hypothesis {
		if h.Id == id {
			r.Hypothesis = append(r.Hypothesis[:i], r.Hypothesis[i+1:]...)
			return nil
		}
	}
	return errors.New("hypothesis not found")
}

func (r *InMemoryRepository) Update(hypothesis *Hypothesis) (err error) {
	for i, h := range r.Hypothesis {
		if h.Id == hypothesis.Id {
			r.Hypothesis[i] = hypothesis
			return nil
		}
	}
	return errors.New("hypothesis not found")
}
