#########################################################################
#     Wifiscanner.py - A simple python script which records and logs wifi probe requests.
#     Author - D4rKP01s0n
#     Requirements - Scapy and Datetime
#     Inspiration - Tim Tomes (LaNMaSteR53)'s WUDS https://bitbucket.org/LaNMaSteR53/wuds/
#     Reminder - Change mon0 (around line 65) to your monitor-mode enabled wifi interface
# https://gist.github.com/dropmeaword/42636d180d52e52e2d8b6275e79484a0
#########################################################################

import time
from datetime import datetime
from scapy.all import sniff, Dot11
#import numpy
import sqlite3
import time
#Devices which are known to be constantly probing
IGNORE_LIST = set(['00:00:00:00:00:00', '01:01:01:01:01:01'])
SEEN_DEVICES = set() #Devices which have had their probes recieved
d = {'00:00:00:00:00:00':'Example MAC Address'} #Dictionary of all named devices
database_name = "../wifi_signals.db"

import aiofiles
import asyncio

async def writeout(data):
    async with aiofiles.open('outputfile.txt', mode='a') as f:
        await f.write(data)

def handle_packet(pkt):
    if not pkt.haslayer(Dot11):
        return
    if pkt.type == 0 and pkt.subtype == 4: #subtype used to be 8 (APs) but is now 4 (Probe Requests)
    #logging.debug('Probe Recorded with MAC ' + curmac)
        curmac = pkt.addr2
        curmac = curmac.upper() #Assign variable to packet mac and make it uppercase
        SEEN_DEVICES.add(curmac) #Add to set of known devices (sets ignore duplicates so it is not a problem)
        if curmac not in IGNORE_LIST: #If not registered as ignored
            if curmac in d:
                print(f"{d[curmac]}--{curmac}")
            else:
                print(f'Device MAC: {pkt.addr2} with SSID: {pkt.info}')
                return pkt.sprintf("%pkt.info%")
                # return pkt.info
        #print SEEN_DEVICES #Just for debug, prints all known devices
        #dump()

def main():
    print('start main sniffing program')
    interface = 'wlan0mon'
    print('begin sniffing interface')
    pkts = []

    # cur.execute("create table lang (name, first_appeared)")

    # This is the qmark style:
    while 1:
        now = datetime.now().strftime("%H:%M:%S")
        res = sniff(iface=interface, prn=handle_packet)
        name = res[0].info.decode('utf-8')
        if name == '':
            continue
        else:
            asyncio.run(writeout(name))
        # cur.execute("insert into wifi_signal values (?, ?, ?, ?)", (None,res[0].info.decode('utf-8'), datetime.now().strftime("%H:%M:%S"),int(time.time())))
        # with open(f'outputfile-{now}','w') as outfl:
        #     outfl.write(res[0].info.decode('utf-8'))
        # print(res[0].info.decode('utf-8'))
        # pkts.append(res[0].info.decode('utf-8')) #start sniffin
        # if len(pkts)>5:
        #     print([x for x in pkts])
        # print(f"after the sniff: {pkts[0].info.decode('utf-8')}")
        time.sleep(1) # Supposed to make an infinite loop, but for some reason it stops after a while



if __name__ == '__main__':
	main()