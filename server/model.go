package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"sync"
)

type Board struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Items []Item `json:"items""`
}

type Item struct {
	Id       string    `json:"id"`
	Type     ItemType  `json:"type"`
	Text     string    `json:"text"`
	VoteUp   int       `json:"vote_up"`
	VoteDown int       `json:"vote_down"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type ItemType int

const (
	WellWent ItemType = iota
	ToImprove
	ActionItems
)

type Model struct {
	storage *Storage
	boards  map[string]*Board
}

var lock = &sync.Mutex{}

func NewModel(storage *Storage) (*Model, error) {
	return &Model{
		storage: storage,
		boards:  make(map[string]*Board),
	}, nil
}

func (model *Model) HandleMessage(boardId string, payload []byte) (DispatchTo, []byte, error) {
	cmd := make(map[string]string)
	err := json.Unmarshal(payload, &cmd)
	if err != nil {
		return ToClient, nil, err
	}
	board := model.getOrCreate(boardId)

	switch cmd["action"] {
	case "init":
		return model.initBoard(board, cmd)
	//case "updateBoard":
	//	return model.updateBoard(board, cmd)
	case "addItem":
		return model.addItem(board, cmd)
	//case "updateItem":
	//	return model.updateItem(board, cmd)
	//case "removeItem":
	//	return model.removeItem(board, cmd)
	//case "voteUp":
	//	return model.voteUp(board, cmd)
	//case "voteDown":
	//	return model.voteDown(board, cmd)
	//case "addComment":
	//	return model.addComment(board, cmd)
	//case "removeComment":
	//	return model.removeComment(board, cmd)
	default:
		log.Printf("Cannot do %s", cmd["action"])
		return ToClient, nil, nil //todo
	}
}

func (model *Model) getOrCreate(boardId string) *Board {
	board, exists := model.boards[boardId]
	if !exists {
		lock.Lock()
		defer lock.Unlock()
		board, exists = model.boards[boardId]
		if !exists {
			newBoard, err := model.storage.FindBoard(boardId)
			if err != nil {
				log.Fatal("Cannot find board", err)
			}
			if newBoard == nil {
				newBoard = &Board{
					Id:    boardId,
					Name:  "New board",
					Items: []Item{},
				}
				err = model.storage.CreateBoard(newBoard)
				if err != nil {
					log.Fatal("Cannot create board", err) //todo
				}
			}
			model.boards[boardId] = newBoard
			board = newBoard
		}
	}
	return board
}

func (model *Model) initBoard(board *Board, _ map[string]string) (DispatchTo, []byte, error) {
	js, err := json.Marshal(board)
	return ToClient, js, err
}

//func (model *Model) updateBoard(board *Board, cmd map[string]string) (DispatchTo, []byte, error) {
//
//}

func (model *Model) addItem(board *Board, cmd map[string]string) (DispatchTo, []byte, error) {
	item := &Item{
		Id:       uuid.NewString(),
		Type:     WellWent, //(cmd["type"], //todo
		Text:     cmd["text"],
		VoteUp:   0,
		VoteDown: 0,
		Comments: []Comment{},
	}
	board.Items = append(board.Items, *item)
	//todo save
	js, err := json.Marshal(item)
	return ToRoom, js, err
}

//func (model *Model) updateItem(board *Board, cmd map[string]string) (DispatchTo, []byte, error) {
//
//}

//func (model *Model) removeItem(board *Board, cmd map[string]string) (DispatchTo, []byte, error) {
//	board.Items.re
//}

//func (model *Model) voteUp(board *Board, cmd map[string]string) (DispatchTo, []byte, error) {
//
//}
//
//func (model *Model) voteDown(board *Board, cmd map[string]string) (DispatchTo, []byte, error) {
//
//}
//
//func (model *Model) addComment(board *Board, cmd map[string]string) (DispatchTo, []byte, error) {
//
//}
//
//func (model *Model) removeComment(board *Board, cmd map[string]string) (DispatchTo, []byte, error) {
//
//}
