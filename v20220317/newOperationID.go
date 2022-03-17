package v20220317

import (
	"github.com/bwmarrin/snowflake"
)

func NewOperationID() string {
	return <-operationIdChan
}

func init() {
	go loop()
}

var (
	NodId           int64
	nd              *snowflake.Node
	operationIdChan = make(chan string)
)

func generateId(nodId int64) (string, error) {
	if nd != nil {
		return nd.Generate().String(), nil
	}
	node, err := snowflake.NewNode(nodId)
	if err != nil {
		return "0", err
	}
	defer func() {
		nd = node
	}()
	return node.Generate().String(), nil
}
func loop() {
	for {
		/*if NodId != 0 */ {
			id, _ := generateId(NodId)
			operationIdChan <- id
			//time.Sleep(time.Microsecond * 100)
			//time.Sleep(time.Millisecond)
		} /*else {
			initNodId()
		}*/
	}
}
