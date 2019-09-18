#!/usr/bin/env bash

HOME=`pwd`

STR="{$HOME}/reptile"

PID=`ps -ef | grep "reptile" | grep -v grep | awk '{print $2}'`

case "$1" in
    "-c")

        $STR = "{$STR} -c {$2}"

    ;;

    "start")

        $STP

    ;;

    "stop")

        kill -15 $PID

    ;;

    "restart")

        kill -s SIGUSR1  $PID

    ;;

esac


case "$3" in
    "start")

        $STP

    ;;

    "stop")

        kill -15 $PID

    ;;

    "restart")

        kill -s SIGUSR1  $PID

    ;;

esac
