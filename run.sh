#!/bin/bash
npm run client
python3 ./googlekeep/keep_endpoint.py
go run ./server/cmd/main.gonpm