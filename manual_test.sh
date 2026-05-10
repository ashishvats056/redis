#!/bin/bash

echo "Connecting to Redis clone on localhost:6379..."
echo "Make sure server is running in another terminal (go run .)"
echo ""

exec 3<>/dev/tcp/localhost/6379

sleep_cmd() {
  echo "Sleeping $1 seconds..."
  sleep "$1"
  echo ""
}

send() {
  echo ">> $1"
  echo "$1" >&3
  read response <&3
  echo "$response"
  sleep 0.1
  echo ""
}

# =====================
# Strings
# =====================
echo "------------------"
echo "---- STRINGS ----"
echo "------------------"
send "SET name john"
send "GET name"
send "DEL name"

echo "---- TTL TEST ----"
send "SET temp value 5"
send "GET temp"
sleep_cmd 5
send "GET temp"

echo "---- MULTIPLE GET ----"
send "SET a 1"
send "SET b 2"
send "SET c 3"
send "MGET a b c d"

# =====================
# Hashes
# =====================

echo "-----------------"
echo "---- HASHES ----"
echo "-----------------"
send "HSET user name alice"
send "HGET user name"

send "HSET user age 30"
send "HGET user age"

echo "---- OVERWRITE FIELD ----"
send "HSET user name bob"
send "HGET user name"

# =====================
# Lists
# =====================

echo "----------------"
echo "---- LISTS ----"
echo "----------------"
send "LADD mylist a"
send "LADD mylist b"
send "LADD mylist c"

echo "---- READ RANGE ----"
send "LRANGE mylist 0 2"

echo "---- INDEX ACCESS ----"
send "INDEX mylist 0"
send "INDEX mylist 2"
echo "---- OUT OF BOUNDS INDEX ----"
send "INDEX mylist 10"

# =====================
# TTL commands
# =====================

echo "---------------------------------------"
echo "---- TTL (EXPIRE / TTL / PERSIST) ----"
echo "---------------------------------------"
echo ""
echo "---- SET + EXPIRE ----"
send "SET key1 value1"
send "EXPIRE key1 5"
send "TTL key1"

sleep_cmd 5

send "GET key1"
send "TTL key1"

echo "---- PERSIST TEST ----"
send "SET key2 value2"
send "EXPIRE key2 10"
send "PERSIST key2"
send "TTL key2"

# =====================
# Mixed stress test
# =====================

echo "---- STRESS TEST ----"
send "SET a 1"
send "SET b 2"
send "SET c 3"

send "HSET user name john"
send "HSET user age 25"

send "LADD list x"
send "LADD list y"
send "LRANGE list 0 10"

send "MGET a b c"

send "EXPIRE a 5"
send "TTL a"

# =====================
# Expiration batch
# =====================

echo "---- EXPIRATION VALIDATION BATCH ----"
send "SET temp1 v1 2"
send "SET temp2 v2 4"
send "SET temp3 v3 6"

send "GET temp1"
send "GET temp2"
send "GET temp3"

sleep_cmd 2

send "GET temp1"
send "GET temp2"
send "GET temp3"

sleep_cmd 2

send "GET temp1"
send "GET temp2"
send "GET temp3"

sleep_cmd 2

send "GET temp1"
send "GET temp2"
send "GET temp3"

# =====================
# Cleanup test
# =====================

echo "---- CLEANUP SANITY CHECK ----"
send "SET cleanup1 a 2"
send "SET cleanup2 b 2"
send "SET cleanup3 c 2"

sleep_cmd 4

send "GET cleanup1"
send "GET cleanup2"
send "GET cleanup3"

echo "DONE"