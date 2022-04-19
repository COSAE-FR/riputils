package helpers

import (
	"fmt"
	"gopkg.in/hlandau/configurable.v1"
	"gopkg.in/hlandau/easyconfig.v1"
	"gopkg.in/hlandau/easyconfig.v1/cflag"
	"gopkg.in/hlandau/service.v2/daemon"
	"gopkg.in/hlandau/svcutils.v1/passwd"
	"os/user"
	"strconv"
)

var (
	parsed bool
	uidFlag string
	gidFlag string
	converted bool
	uid int = -1
	gid int = -1
)

func populateIds() {
	if parsed {
		return
	}
	easyconfig.Parse(nil, nil)
	configurable.Visit(func(c configurable.Configurable) error {
		fmt.Printf("%+v\n", c)
		obj, ok := c.(*cflag.Group)
		if ok && obj.String() == "service" {
			for _, sub := range obj.CfChildren() {
				flg, ok := sub.(*cflag.StringFlag)
				if ok {
					switch flg.String() {
					case "uid":
						uidFlag = flg.Value()
					case "gid":
						gidFlag = flg.Value()
					}
				}
			}
		}
		return nil
	})
	parsed = true
}

func convertIds() {
	if converted {
		return
	}
	var err error
	if uidFlag != "" && gidFlag == "" {
		gid, err = passwd.GetGIDForUID(uidFlag)
		if err == nil {
			gidFlag = strconv.FormatInt(int64(gid), 10)
		}
	}
	if uidFlag != "" {
		uid, _ = passwd.ParseUID(uidFlag)
		gid, _ = passwd.ParseGID(gidFlag)
	}
	converted = true
}

func RunWithPrivileges() bool {
	current, err := user.Current()
	if err != nil {
		return false
	}
	currentUid, err := passwd.ParseUID(current.Uid)
	if err != nil {
		return false
	}
	if currentUid != 0 {
		return false
	}
	populateIds()
	convertIds()
	if uidFlag != "" && gidFlag == "" {
		gid, err := passwd.GetGIDForUID(uidFlag)
		if err != nil {
			return true
		}
		gidFlag = strconv.FormatInt(int64(gid), 10)
	}
	if uid < 1 || gid < 0 {
		return true
	}
	return false
}

func DropPrivileges() error {
	if !RunWithPrivileges() {
		_, err := daemon.DropPrivileges(uid, gid, "/")
		if err != nil {
			return fmt.Errorf("failed to drop privileges: %v", err)
		}
	}
	return nil
}

