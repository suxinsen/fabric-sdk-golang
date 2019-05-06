package device

import (
	"device-data-server/fabric"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func Start() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>> start camera data sync >>>>>>>>>>>>>>>>>>>>>>")
	timer()
}

var (
	pathMap = make(map[string]string)
	deviceId = beego.AppConfig.String("device.id")
	deviceDataPath = beego.AppConfig.String("device.list.cmd")
)

func syncNewFile() {
	pathBytes := Command("bash","-c", deviceDataPath)
	regexp := regexp.MustCompile("/opt/camera/motion/[0-9_-]{15,30}.avi")
	params := regexp.FindAll(pathBytes, -1)
	for _, path := range params {
		id := deviceId + Between(string(path), "-", ".")
		if _, ok := pathMap[id]; !ok {
			log.Printf("add new movie file %s to map \n", path)
			pathMap[id] = string(path)
			//movie file to md5sum or hash
			md5sumByte := Command("bash", "-c", "md5sum " + string(path))
			if md5sumByte != nil {
				md5sum := strings.Split(string(md5sumByte), " ")
				log.Printf("get movie file md5sum is %s", md5sum[0])
				var args [][]byte
				args = append(args, []byte(id))
				args = append(args, []byte(md5sum[0]))
				result := fabric.ObtainSdkUtil().Invoke("upload", args)
				log.Println(result)
			}
		}

	}
}

func Command(name string, args ...string) []byte{
	cmd := exec.Command(name,args...)
	pathBytes, err := cmd.Output()
	if err != nil {
		log.Printf("exec linux cmd err, info: %s \n", err.Error())
		return nil
	}
	return pathBytes
}

func Between(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}

func timer( ) {

	//300秒出发一次
	tick := time.NewTicker( 300 * time.Second)

	for {
		select {
		//此处在等待channel中的信号，因此执行此段代码时会阻塞120秒
		case <- tick.C :
			fmt.Println(">>>>>>>>>>>>  camera data sync  <<<<<<<<<<<<")
			syncNewFile()
		}
	}
}
