package eureka

import (
	"encoding/json"
	"strings"
	"fmt"
)

func (c *Client) RegisterInstance(appId string, instanceInfo *InstanceInfo) error {
	values := []string{"apps", appId}
	path := strings.Join(values, "/")
	instance := &Instance{
		Instance: instanceInfo,
	}
	body, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	var v *RawResponse
	v, err = c.Post(path, body)
	fmt.Println(v)
	return err
}
