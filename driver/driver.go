package driver

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/kubeedge/mapper-framework/pkg/common"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/udp"
	"k8s.io/klog/v2"
)

func NewClient(protocol ProtocolConfig) (*CustomizedClient, error) {
	client := &CustomizedClient{
		ProtocolConfig: protocol,
		deviceMutex:    sync.Mutex{},
		// TODO initialize the variables you added
	}
	return client, nil
}

func (c *CustomizedClient) InitDevice() error {
	// TODO: add init operation
	// you can use c.ProtocolConfig
	return nil
}

func (c *CustomizedClient) GetDeviceData(visitor *VisitorConfig) (interface{}, error) {
	ip := c.ProtocolConfig.ConfigData.MoteAddr
	query := visitor.QueryPath
	klog.Infof("GET request to %s%s", ip, query)

	co, err := udp.Dial(ip)
	if err != nil {
		klog.Errorf("Connection error : %v", err)
		return nil, err
	}
	defer co.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := co.Get(ctx, query)
	if err != nil {
		klog.Errorf("GET Error : %v", err)
		return nil, err
	}

	fmt.Println(res.Code())
	if res.Code() == codes.Content {
		payload, err := io.ReadAll(res.Body())
		if err != nil {
			return nil, err
		}
		value := string(payload)
		return value, nil
	}
	return nil, nil
}

func (c *CustomizedClient) DeviceDataWrite(visitor *VisitorConfig, deviceMethodName string, propertyName string, data interface{}) error {
	// TODO: add the code to write device's data
	// you can use c.ProtocolConfig and visitor to write data to device
	return nil
}

func (c *CustomizedClient) SetDeviceData(data interface{}, visitor *VisitorConfig) error {
	// TODO: set device's data
	// you can use c.ProtocolConfig and visitor
	return nil
}

func (c *CustomizedClient) StopDevice() error {
	// TODO: stop device
	// you can use c.ProtocolConfig
	return nil
}

func (c *CustomizedClient) GetDeviceStates() (string, error) {
	ip := c.ProtocolConfig.ConfigData.MoteAddr
	co, err := udp.Dial(ip)
	if err != nil {
		klog.Errorf("Connection error : %v", err)
		return common.DeviceStatusDisCONN, err
	}
	defer co.Close()
	return common.DeviceStatusOK, nil
}
