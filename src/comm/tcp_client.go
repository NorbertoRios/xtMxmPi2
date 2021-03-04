package comm

import (
	"controller"
	"dto"
	"fmt"
	"interfaces"
	"net"
	"time"
)

//Client struct for tcp connection
type Client struct {
	Connection     net.Conn
	ConnectedAt    time.Time
	LastActivityTs time.Time
	Received       int64
	Transmitted    int64
	PageBuffer     *interface{}
	Device         *interfaces.Device
	Session        string
	VideoHandler   *dto.VideoHandler
}

// Send text message to client
func (c *Client) Send(message string) error {
	count, err := c.Connection.Write([]byte(message))
	ServerCounters.AddFloat("Transmitted", float64(count))
	c.Transmitted += int64(count)
	return err
}

//SendBytes packet to client
func (c *Client) SendBytes(message []byte) error {
	count, err := c.Connection.Write(message)
	ServerCounters.AddFloat("Transmitted", float64(count))
	c.Transmitted += int64(count)
	return err
}

func (c *Client) CloseConnection() {
	c.Connection.Close()
}

//ConnectedAtTs returns connected timestamp
func (c *Client) ConnectedAtTs() time.Time {
	return c.ConnectedAt
}

//ReceivedBytes total received bytes
func (c *Client) ReceivedBytes() int64 {
	return c.Received
}

//TransmittedBytes total transmitted bytes
func (c *Client) TransmittedBytes() int64 {
	return c.Transmitted
}

//RemoteAddr client's ip address
func (c *Client) RemoteAddr() string {
	return c.Connection.RemoteAddr().String()
}

//RemoteIP indicates device remote IP address
func (c *Client) RemoteIP() string {
	return fmt.Sprintf("%v", c.Connection.RemoteAddr().String())
}

//RemotePort indicates device remote port
func (c *Client) RemotePort() int {
	return 0
}

//Listen client data from channel
func (c *Client) Listen() {
	buffer := make([]byte, 4096)
	tso := &dto.TSOBuffer{}
	for {
		count, err := c.Connection.Read(buffer)
		if err != nil {
			c.Connection.Close()
			return
		}
		c.LastActivityTs = time.Now().UTC()
		c.Received = c.Received + int64(count)
		ServerCounters.AddFloat("Received", float64(count))
		tb := buffer[:count]
		//justPrint(tb)
		segmentBuffer := dto.HandlePackageWithSO(tso, tb, c, &dto.GeneralPackageHeader{}, controller.HandleTCPPacket)
		for segmentBuffer != nil {
			segmentBuffer = dto.HandlePackageWithSO(tso, segmentBuffer, c, &dto.GeneralPackageHeader{}, controller.HandleTCPPacket)
		}
	}
}

func justPrint(tb []byte) {
	fmt.Printf("New packet: %X", tb)
	fmt.Println()
	fmt.Printf("New packet as text: %s", tb)
	fmt.Println()
}

//LastActivity indicates last device activity
func (c *Client) LastActivity() time.Time {
	return c.LastActivityTs
}

func (c *Client) GetPageBuffer() *interface{} {
	return c.PageBuffer
}

func (c *Client) SetPageBuffer(buffer *interface{}) {
	c.PageBuffer = buffer
}

func (c *Client) GetDevice() *interfaces.Device {
	return c.Device
}

func (c *Client) SetDevice(device *interfaces.Device) {
	c.Device = device
}

func (c *Client) GetCurrentSession() string {
	return c.Session
}

func (c *Client) SetCurrentSession(s string) {
	c.Session = s
}

func (c *Client) SetVideoHandler(vh interface{}) {
	c.VideoHandler = vh.(*dto.VideoHandler)
}
func (c *Client) GetVideoHandler() interface{} {
	return c.VideoHandler
}
