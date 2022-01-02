#!/bin/bash
export HTTP_ADDR=localhost:2727
export SECRET_CODE=8109612acafb4ae0ff34f5f1fa549577f4ca3a4a294f559498c111cc7d92973e5dde4eb64f086b49e063708705338f29b662047c09c850f5bb21da65f37036b4

# DB settings
export INIT_DB="false"

# APP settings
export FOLDERDATA="data"
export APPURL="https://quasar.m0ta.ru"
#export APPEMAIL="Benefy team <info@benefy.ru>"
export EMAILFROM="Benefy team <info@benefy.ru>"
export EMAILHOST="smtp.yandex.ru"
export EMAILPORT=465
export EMAILUSER="info@benefy.ru"
export EMAILPASS="fkFK44oG8Z"

# API Exchange settings
export FIXERIO="false"
export FIXERIO_PERIOD="59 */6 * * 1,2,3,4,5"
export FIXERIO_API="http://data.fixer.io/api"
export FIXERIO_TOKEN="3220a8f2dc5be11491d4e770ae3d39c8"

export YAHOO="false"
#export YAHOO_PERIOD="05 02 * * 1,2,3,4,5"
export YAHOO_PERIOD="2/30 4-22 * * 1,2,3,4,5"
export YAHOO_DELAY="10"
export YAHOO_API="https://query1.finance.yahoo.com/v11/finance/quoteSummary"
export YAHOO_TOKEN=""

export MARKETSTACK="false"
export MARKETSTACK_PERIOD_S="05 02 * * 1,2,3,4,5"
export MARKETSTACK_PERIOD_I="58 */6 * * 1,2,3,4,5"
#export MARKETSTACK_PERIOD_I="58 * * * 1,2,3,4,5"
export MARKETSTACK_PERIOD_E="00 22 * * 1,2,3,4,5"
export MARKETSTACK_API="http://api.marketstack.com/v1"
export MARKETSTACK_TOKEN="5f5a447111243c8e88919f3b453d84cc"

export ISSMOEX="false"
export ISSMOEX_PERIOD_S="05 02 * * 1,2,3,4,5" #строка cron - обновление информации о ценных бумагах
export ISSMOEX_PERIOD_I="*/30 4-22 * * 1,2,3,4,5" #строка cron - обновление информации intraday
export ISSMOEX_PERIOD_E="01 22 * * 1,2,3,4,5" #строка cron - обновление информации end of day
export ISSMOEX_API="http://iss.moex.com/iss"
export ISSMOEX_TOKEN=""

export CBR="false"
export CBR_PERIOD_S="05 02 * * 1,2,3,4,5" #строка cron - обновление информации о ценных бумагах
export CBR_PERIOD_I="*/30 4-22 * * 1,2,3,4,5" #строка cron - обновление информации intraday
export CBR_PERIOD_E="01 22 * * 1,2,3,4,5" #строка cron - обновление информации end of day
export CBR_API="http://www.cbr.ru/scripts"
export CBR_TOKEN=""

# Postgres settings
#export PG_URL="postgresql://mota:motamota@rc1c-03nho7rtlgd40ai2.mdb.yandexcloud.net:6432/doxod?connect_timeout=30&sslmode=require&sslrootcert=/home/mota/.postgresql/root.crt"
#export PGX_URL="host=rc1c-03nho7rtlgd40ai2.mdb.yandexcloud.net port=6432 dbname=doxod user=mota password=motamota connect_timeout=30 sslmode=require sslrootcert=/home/mota/.postgresql/root.crt"
#export PG_URL="postgres://mota:motamota@rc1c-3d9q0p5x421j77vl.mdb.yandexcloud.net:6432/benefy?sslmode=disable"
export PG_URL="postgres://mota:motamota@rc1c-pd00n5y6rttquslp.mdb.yandexcloud.net:6432/benefy?sslmode=disable"
#export PG_MIGRATIONS_PATH="file://../../app/store/pg/migrations/"

# Logger settings
export LOG_LEVEL=debug

# Token settings
export TOKENEXP=720h
export TOKENKEY=mSv87L7KPfPDDyjvmqf9t8RVyrc-6wLJeGtCSv5fMjCSd2dTE6kX3pg-Ny3QyFh5G9x6ZkgdNseseZZpDcj