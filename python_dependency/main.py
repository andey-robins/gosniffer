'''
file: main.py
auth: rafer cooley
desc: this file is the main entry point for the sniffer program of the wifi signal sniffer project
configure network card:
> sudo ifconfig wlan0 down
> sudo iwconfig wlan0 mode monitor
'''

from scapy.all import *
from threading import Thread
import pandas, sqlite3, netifaces
import time, os

database_name = "../wifi_signals.db"

# initialize the networks dataframe that will contain all access points nearby
networks = pandas.DataFrame(columns=["BSSID", "SSID", "dBm_Signal", "Channel", "Crypto"])
# set the index BSSID (MAC address of the AP)
networks.set_index("BSSID", inplace=True)

def create_database():
    '''
    create the database structure if not already exists
    '''
    if not os.path.exists(database_name):
        print(f"Database {database_name} does not exist, making now")
        conn = sqlite3.connect(database_name)
        conn.execute('''
        create table wifi_signal
        (ID INTEGER PRIMARY KEY,
        NAME TEXT NOT NULL,
        DATESTR TEXT NOT NULL,
        EPOCHTIME INTEGER NOT NULL);
        ''')
        print('Database created')
        conn.close()
    else:
        print(f"Database {database_name} Exists")

def print_all():
    '''
    Now we need a way to visualize this dataframe. Since we're going to use sniff() function (which blocks and start sniffing in the main thread), we need to use a separate thread to print the content of networks dataframe, the below code does that:
    '''
    while True:
        os.system("clear")
        print(networks)
        time.sleep(0.5)

def callback(packet):
    '''
    This callback makes sure that the sniffed packet has a beacon layer on it, if it is the case, then it will extract the BSSID, SSID (name of access point), signal and some stats. the Scapy's Dot11Beacon class has the awesome network_stats() function that extracts some useful information from the network, such as the channel, rates and the encryption type. Finally, we add these information to the dataframe with the BSSID as the index.

    You will encounter some networks that doesn't have the SSID (ssid equals to ""), this is an indicator that it's a hidden network. In hidden networks, the access point leaves the info field blank to hide the discovery of the network name, you will still find them using this tutorial's script, but without a network name.
    '''
    if packet.haslayer(Dot11Beacon):
        # extract the MAC address of the network
        bssid = packet[Dot11].addr2
        # get the name of it
        ssid = packet[Dot11Elt].info.decode()
        try:
            dbm_signal = packet.dBm_AntSignal
        except:
            dbm_signal = "N/A"
        # extract network stats
        stats = packet[Dot11Beacon].network_stats()
        # get the channel of the AP
        channel = stats.get("channel")
        # get the crypto
        crypto = stats.get("crypto")
        networks.loc[bssid] = (ssid, dbm_signal, channel, crypto)

def change_channel():
    '''
    Now if you execute this, you will notice not all nearby networks are available, that's because we're listening on one WLAN channel only. We can use the iwconfig command to change the channel, here is the Python function for it:
    '''
    ch = 1
    interface = 'wlan0'
    while True:
        os.system(f"iwconfig {interface} channel {ch}")
        # switch channel from 1 to 14 each 0.5s
        ch = ch % 14 + 1
        time.sleep(0.5)

if __name__ == "__main__":
    create_database()

    net_interfaces = netifaces.interfaces()
    # start the channel changer
    # channel_changer = Thread(target=change_channel)
    # channel_changer.daemon = True
    # channel_changer.start()

    # # interface name, check using iwconfig
    # interface = "wlan0mon"
    # # start the thread that prints all the networks
    # printer = Thread(target=print_all)
    # printer.daemon = True
    # printer.start()
    # # start sniffing
    # sniff(prn=callback, iface=interface)

## References
#- main structure of program: https://www.thepythoncode.com/article/building-wifi-scanner-in-python-scapy
