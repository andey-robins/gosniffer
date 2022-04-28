// This project was inspired and based on the work from [this repo](This project is based on [this repo](https://github.com/stevemcquaid/War-Walker)

/**The MIT License (MIT)

Copyright (c) 2016 by Daniel Eichhorn

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

See more at http://blog.squix.ch
*/

#include <Arduino.h>
#include <algorithm>
#include <vector>
#include <stdlib.h>
#include <string>
#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <Ticker.h>
#include <JsonListener.h>
#include "Wire.h"
#include "TimeClient.h"

//the structure
class Network {
public:
  String StationMac;
  String BSSID;
  String SSID;
  Network(){};
};

//the initialize function
Network constructNet(String StationMac, String BSSID, String SSID){
  Network foo;
  foo.StationMac  = StationMac;
  foo.BSSID       = BSSID;
  foo.SSID       = SSID;

  return foo;
}

Network constructNet(String StationMac, String BSSID, String SSID);
bool networkSorter(Network const& lhs, Network const& rhs);
std::vector<Network> scan();
void setupScan();

void setupScan(){
  Serial.begin(115200);

  // Set WiFi to station mode and disconnect from an AP if it was previously connected
  WiFi.mode(WIFI_STA);
  WiFi.disconnect();
  delay(100);

  Serial.println("Setup done");
}

std::vector<Network> populateNetArray(int n){
  std::vector<Network> netArray(n);
  for (int i=0; i < n; i++){
    netArray[i] = constructNet(WiFi.macAddress(), WiFi.BSSIDstr(), WiFi.SSID());
  }
  return netArray;
}

std::vector<Network> scan() {
  Serial.println("Starting Scan...");
  // WiFi.scanNetworks will return the number of networks found
  bool async = false;
  bool show_hidden = false;
  int n = WiFi.scanNetworks(async, show_hidden);

  serialDebugScan(n);

  std::vector<Network> netArray = populateNetArray(n);
  std::sort(netArray.begin(), netArray.end(), networkSorter);
  return(netArray);
}

void serialDebugScan(int n){
  Serial.println("scan done");
  if (n == 0){
    Serial.println("no networks found");
  } else
  {
    Serial.print(n);
    Serial.println(" networks found");
  }
  Serial.println("");
}

// WiFi connection and POST request were inspired by [this](https://techtutorialsx.com/2016/07/21/esp8266-post-requests/)
void WiFi_Connect() {
  // Establish a connection to a WiFi network
  Serial.begin(115200);                 //Serial connection
  WiFi.begin("yourSSID", "yourPASS");   //WiFi connection
 
  while (WiFi.status() != WL_CONNECTED) {  //Wait for the WiFI connection completion
 
    delay(500);
    Serial.println("Waiting for connection");
 
  }
}

void POST(Network network) {
    if (WiFi.status() == WL_CONNECTED) { //Check WiFi connection status
 
    HTTPClient http;    //Declare object of class HTTPClient
 
    http.begin("http://seekdanyouwillbefound.org/register");      //Specify request destination
    http.addHeader("Content-Type", "application/json");  //Specify content-type header

    //int stationmaclen = network.StationMac.length() + 1; // Add one for the empty string we'll concat with

    char emptyStr[] = "";
 
    strcat(emptyStr, "mac: ");
    strcat(emptyStr, network.StationMac.c_str());
    strcat(emptyStr, "power: 0");
    strcat(emptyStr, "packetCount: 0");
    strcat(emptyStr, "bssid: ");
    strcat(emptyStr, network.BSSID.c_str());
    strcat(emptyStr, "essid: ");
    strcat(emptyStr, network.SSID.c_str());

    int httpCode = http.POST(emptyStr);   //Send the request
    String payload = http.getString();                  //Get the response payload
 
    Serial.println(httpCode);   //Print HTTP return code
    Serial.println(payload);    //Print request response payload
 
    http.end();  //Close connection
 
  } else {
 
    Serial.println("Error in WiFi connection");
 
  }
 
  delay(30000);  //Send a request every 30 seconds
}

void loop()
{
  std::vector<Network> vec;
  vec = scan();
  for (auto i: vec) {
    POST(i);
  }

}