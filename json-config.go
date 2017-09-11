package core

import (
	"encoding/json"
)


type Jsoner interface {
	
}

type JsonParser func(Jsoner)

func (loader JsonParser)Unmarshal(buf []byte, jsoner Jsoner) error {

    
	err := json.Unmarshal(buf, jsoner)
	return err
}

