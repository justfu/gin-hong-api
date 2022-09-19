xiangmu='gin-hong-api'
xiangmupath='/golang/go/src/'
xiangmubinpath='/golang/go/bin'
cd /data/git/$xiangmu/
cd $xiangmupath$xiangmu/
go install .
cd $xiangmupath$xiangmu/cmd/queue
go install .

#运行任务
#nohup $xiangmubinpath/gin -env pro > /tmp/gin.log &
#nohup $xiangmubinpath/queue -env pro > /tmp/queue.log &

#重启所有命令
supervisorctl reload
