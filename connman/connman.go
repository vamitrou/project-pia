package connman

import (
	"fmt"
	"time"
)

var live_connections map[string]interface{} = make(map[string]interface{})

func GetRConnection(app_id string, live bool) (*rconn, error) {
	if live {
		if conn, ok := live_connections[app_id]; ok {
			fmt.Println("reusing connection")
			c, _ := conn.(*rconn)
			c.last_accessed = time.Now()
			return c, nil
		} else {
			rc, err := NewRConnection()
			live_connections[app_id] = rc
			return rc, err
		}
	} else {
		rc, err := NewRConnection()
		return rc, err
	}
}

func Recycle(rc *rconn) {
	keepAlive := 30.0
	accessed_before := time.Since(rc.last_accessed).Seconds()
	fmt.Println(accessed_before)
	if accessed_before >= keepAlive {
		fmt.Println("Closing connection")
		rc.Close()
	} else {
		dur, _ := time.ParseDuration(fmt.Sprintf("%fs", keepAlive-accessed_before))
		time.Sleep(dur)
		go Recycle(rc)
	}
}
