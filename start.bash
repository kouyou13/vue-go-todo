#!/usr/bin/env bash

# frontend を起動
(
  cd ./frontend || exit 1
  npm run dev
) &

# backend を起動
(
  cd ./backend || exit 1
  go run main.go
) &

# 両方のプロセスが終了するまで待つ
wait