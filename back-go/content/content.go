package content

import (
	"fmt"
	"os"

	"github.com/shanliao420/learn1000/back-go/do"
)

const (
	DataPath = "./data/"
)


type Content struct {
    data []do.Item
}

func NewContent() *Content {
    return &Content{
        data: make([]do.Item, 0),
    }
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
    
}