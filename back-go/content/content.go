package content

import (
	"fmt"
	"log"
	"os"

	"github.com/shanliao420/learn1000/back-go/do"
	"github.com/shanliao420/learn1000/back-go/utils"
)

const (
	DataPath                = "./data/"
	SourceFileBase          = "source.data"
	TranslateFileBase       = "translate.data"
	FirstRecordAddrFileBase = "firstRecordAddr.data"
	LastRecordAddrFileBase  = "lastRecordAddr.data"
)

type Content struct {
	idm *utils.IDMaker

	data map[string]do.Item
}

func NewContent() *Content {
	c := &Content{
		idm:  utils.NewIDMaker(),
		data: make(map[string]do.Item),
	}
	c.Init()
	return c
}

// init
func (c *Content) Init() {

	files, err := os.ReadDir(DataPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		c.addItem(file.Name())
	}
}

func (c *Content) addItem(id string) {
	root := getItemRoot(id)
	source := root + SourceFileBase
	translate := root + TranslateFileBase
	firstRecordAddr := root + FirstRecordAddrFileBase
	lastRecordAddr := root + LastRecordAddrFileBase
	sourceString, err := utils.ReadFile(source)
	if err != nil {
		log.Println("Error reading source file:", err, "id: ", id)
		return
	}
	translateString, err := utils.ReadFile(translate)
	if err != nil {
		log.Println("Error reading file:", err, "id: ", id)
		return
	}
	item := &do.Item{
		Id:              id,
		Source:          sourceString,
		Translation:     translateString,
		FirstRecordAddr: firstRecordAddr,
		LastRecordAddr:  lastRecordAddr,
	}
	c.data[id] = *item
}

func (c *Content) GetItem(id string) *do.Item {
	item, ok := c.data[id]
	if !ok {
		return nil
	}
	return &item
}

func (c *Content) GetIdList() []string {
	idList := make([]string, 0, len(c.data))
	for id := range c.data {
		date := c.idm.GetDateFromID(id)
		idList = append(idList, date)
	}
	return idList
}

// delete item
func (c *Content) DeleteItem(id string) {
	delete(c.data, id)
}

// modify item
func (c *Content) ModifyItem(item do.Item) {
	c.data[item.Id] = item
	root := getItemRoot(item.Id)
	source := root + SourceFileBase
	translate := root + TranslateFileBase
	utils.SaveFile(source, item.Source)
	utils.SaveFile(translate, item.Translation)
}


// add item
func (c *Content) AddItem(item do.Item) {
    id := c.idm.MakeID()
	item.Id = id
	c.data[id] = item
	root := getItemRoot(id)
	source := root + SourceFileBase
	translate := root + TranslateFileBase
	utils.SaveFile(source, item.Source)
	utils.SaveFile(translate, item.Translation)
}


func getItemRoot(id string) string {
	return DataPath + id + "/"
}