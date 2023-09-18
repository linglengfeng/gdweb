#/bin/bash

cfg_init(){
	cfg['workspace']=`pwd`
	cfg['svr_list']='web3Server webServer'
}

init(){
    cfg_init
    ROOT=${cfg['workspace']}
    cd ${ROOT} || exit 0
}

fun_startall() {
	svr_list=${cfg['svr_list']}
	for svr_name in ${svr_list[*]};do
		echo "startall ${svr_name}"
		fun_start "${svr_name}"
	done
	exit 0
}

fun_stopall() {
	svr_list=${cfg['svr_list']}
	for svr_name in ${svr_list[*]};do
		echo "stopall ${svr_name}"
		fun_stop "${svr_name}"
	done
	exit 0
}

fun_start() {
	svr_name=$1
	echo "start === ${svr_name} ==="
	if ps aux | grep -v grep | grep -v ctl | grep "${svr_name}"; then
		echo "start error: ${svr_name} is running."
	else
		cd ${ROOT}
		cd ../${svr_name}
		go build
		nohup ./${svr_name} &
		echo "start ${svr_name} successed."
	fi
}

fun_stop() {
	svr_name=$1
	echo "stop === ${svr_name} ==="
	if ps aux | grep -v grep | grep -v ctl | grep "${svr_name}"; then
		pkill -TERM -f "$svr_name"
		# 检查是否成功关闭进程
		if [ $? -eq 0 ]; then
			echo "$svr_name has been gracefully stopped."
		else
			echo "$svr_name could not be stopped."
		fi
	else
		echo "stop ${svr_name} already stoped."
	fi
}

help(){
    echo "startall 启动"
	echo "stopall 停止"
	echo "start name 启动一个"
	echo "stop name 停止一个"
}

# ------------------------------------------------------
# 执行入口
# ------------------------------------------------------
declare -A cfg
init
case $1 in
    startall) fun_startall;;
	stopall) fun_stopall;;
	start) fun_start $2;;
	stop) fun_stop $2;;
    *)
        echo "未知指令，请使用以下有效指令"
        echo "----------------------------------------------------------"
        help
        exit 1
        ;;
esac
exit 0

