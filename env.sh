#!/bin/bash
export HTTP_ADDR=localhost:2727

# DB settings
export INIT_DB="false"

# APP settings
export FOLDERDATA="data"
export PG_URL="postgres://postres:postres@dbserver:5432/database?sslmode=disable"

# Logger settings
export LOG_LEVEL=debug

# Token settings
export TOKENEXP=720h
export TOKENKEY=625da44e4eaf58d61cf048d168aa6f5e492dea166d8bb54ec06c30de07db57e1