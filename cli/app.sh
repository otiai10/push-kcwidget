#!/bin/sh

app="github.com/otiai10/push-kcwidget"

get_pids() {
    local mode=$1
    echo `ps aux | grep revel | grep $mode | awk '{print $2}'`
}
is_running() {
    local mode=$1
    for pid in `get_pids $mode`; do
        echo 1
        return
    done
    echo 0
}
kill_by_pid() {
    local pid=$1
    kill -9 $pid
}
kill_by_mode() {
    local mode=$1
    for pid in `get_pids $mode`; do
        kill_by_pid $pid
    done
}
show_help() {
    echo "USAGE"
    echo "\trun {dev|prod}"
    echo "\tstop {dev|prod}"
    echo "\trestart {dev|prod}"
}

if [ "$1" = "run" ]; then
    mode="dev"
    if [ "$2" = "prod" ]; then
        mode="prod"
    fi
    if [ `is_running $mode` = 1 ]; then
        echo "Already running in $mode mode"
        exit 1
    fi
    # process is clean
    revel run $app $mode
elif [ "$1" = "stop" ]; then
    mode="dev"
    if [ "$2" = "prod" ]; then
        mode="prod"
    fi
    kill_by_mode $mode   
elif [ "$1" = "restart" ]; then
    mode="dev"
    if [ "$2" = "prod" ]; then
        mode="prod"
    fi
    kill_by_mode $mode
    revel run $app $mode
else
    show_help
fi
