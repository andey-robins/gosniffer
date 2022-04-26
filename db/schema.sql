CREATE TABLE networks (
    uid INTEGER PRIMARY KEY AUTOINCREMENT,
    mac TEXT NOT NULL,
    power TEXT NOT NULL,
    packetCount TEXT NOT NULL,
    bssid TEXT NOT NULL,
    essid TEXT NOT NULL 
);