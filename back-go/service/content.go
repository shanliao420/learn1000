package service

import (
	"github.com/shanliao420/learn1000/back-go/content"
	"github.com/shanliao420/learn1000/back-go/do"
	"github.com/shanliao420/learn1000/back-go/vo"
)

type ContentService struct {
	contentManager *content.Content
}

func NewContentService() *ContentService {
	return &ContentService{
		contentManager: content.NewContent(),
	}
}

// get list
func (s *ContentService) List() vo.List {
	s.NewItem()
	idList := s.contentManager.GetIdList()
	result := vo.List{
		IdList: idList,
		Total:  len(idList),
	}
	return result
}

func (s *ContentService) Get(id string) *do.Item {
	item := s.contentManager.GetItem(id)
	return item
}

func (s *ContentService) Add(item do.Item) error {
	s.contentManager.AddItem(item)
	return nil
}

func (s *ContentService) Delete(id string) error {
	s.contentManager.DeleteItem(id)
	return nil
}

func (s *ContentService) Update(item do.Item) {
	s.contentManager.ModifyItem(item)
}

func (s *ContentService) NewItem() {
	s.contentManager.NewBlankItem()
}
