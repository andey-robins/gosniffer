#!/bin/bash
echo "This may hang for some time, be patient."
# Name is specific, if the adapter doesn't default to something nice you've gotta change it
# These commands will set the adapter to monitor mode (v2 TL-WN722N)
sudo ifconfig wlxb0a7b9582f8a down
sudo airmon-ng check kill
iwconfig wlxb0a7b9582f8a mode monitor
ifconfig wlxb0a7b9582f8a up

# Start listening for wifi beacons
sudo airodump-ng -w tcpdumpoutput/myOutput --output-format csv wlxb0a7b9582f8a