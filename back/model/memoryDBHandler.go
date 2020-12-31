package model

import (
	"errors"
)


type memoryHandler struct {
	toDoMap map[int]*ToDo
}

func (h *memoryHandler) GetToDos() []*ToDo {
	slice := []*ToDo{}
	for _, v := range h.toDoMap {
		slice = append(slice, v)
	}
	return slice
}

func (h *memoryHandler) AddToDo(newToDo *ToDo) (*ToDo, error) {
	_, ok := h.toDoMap[newToDo.ID]
	if ok == true {
		return nil, errors.New("Same ID Error")
	}
	h.toDoMap[newToDo.ID] = newToDo
	return newToDo, nil
}

func (h *memoryHandler) DeleteToDo(id int) error {
	_, ok := h.toDoMap[id]
	if ok == false {
		return errors.New("ID does not exist")
	}
	delete(h.toDoMap, id)
	return nil
}

func (h *memoryHandler) CompleteToDo(id int) error {
	toDo, ok := h.toDoMap[id]
	if ok == false {
		return errors.New("ID does not exist")
	}
	toDo.Complete = !toDo.Complete
	return nil
}

func (h *memoryHandler) GetDetail(id int) (*ToDo, error) {
	toDo, ok := h.toDoMap[id]
	if ok == false {
		return nil, errors.New("ID does not exist")
	}
	return toDo, nil
}

func (h *memoryHandler) Close() {
	
}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.toDoMap = make(map[int]*ToDo)
	return m
}