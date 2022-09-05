package main

import (
	"fmt"
	"strings"

	"github.com/jessevdk/go-flags"
)

type Opts struct {
	X []string `short:"D"`
	D []string `short:"X"`
}

var exam = "/usr/local/jdk1.8.0_321-amd64//bin/java -Dproc_nodemanager -Dhdp.version=3.1.5.0-152 -Djava.net.preferIPv4Stack=true -Dhdp.version=3.1.5.0-152 -Dyarn.id.str= -Dyarn.policy.file=hadoop-policy.xml -Djava.io.tmpdir=/var/lib/ambari-agent/tmp/hadoop_java_io_tmpdir -Dnm.audit.logger=INFO,NMAUDIT -Dyarn.log.dir=/var/logs/hadoop-yarn/yarn -Dyarn.log.file=hadoop-yarn-nodemanager-office001135.log -Dyarn.home.dir=/usr/hdp/3.1.5.0-152/hadoop-yarn -Dyarn.root.logger=INFO,console -Djava.library.path=:/usr/hdp/3.1.5.0-152/hadoop/lib/native/Linux-amd64-64:/usr/hdp/3.1.5.0-152/hadoop/lib/native/Linux-amd64-64:/var/lib/ambari-agent/tmp/hadoop_java_io_tmpdir:/usr/hdp/3.1.5.0-152/hadoop/lib/native -Xmx1024m -Dhadoop.log.dir=/var/logs/hadoop-yarn/yarn -Dhadoop.log.file=hadoop-yarn-nodemanager-office001135.log -Dhadoop.home.dir=/usr/hdp/3.1.5.0-152/hadoop -Dhadoop.id.str=yarn -Dhadoop.root.logger=INFO,RFA -Dhadoop.policy.file=hadoop-policy.xml -Dhadoop.security.logger=INFO,NullAppender org.apache.hadoop.yarn.server.nodemanager.NodeManager"
// var exam = "/usr/local/jdk1.8.0_321-amd64//bin/java -Dproc_nodemanager org.apache.hadoop.yarn.server.nodemanager.NodeManager"

func main() {
	opts := &Opts{}
	// fmt.Println(os.Args)
	args := strings.Fields(exam)
	r, err := flags.ParseArgs(opts, args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Println(err)
		return
	}
	// Finally print the collected string
	fmt.Println(*opts)
	fmt.Println(r)
}
