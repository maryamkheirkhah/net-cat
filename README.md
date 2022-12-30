## NET-CAT, grit:lab golang project

### PROJECT OVERVIEW

The aim of this project was to recreate the NetCat command-line utility in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server. The utility reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports and many more.

### REQUIREMENTS

The following features for the utility were specified:

> __01.__
TCP connection between server and multiple clients (relation of 1 to many)<br>

> __02.__
A name requirement to the client.<br>

> __03.__
Control connections quantity.<br>

> __04.__
Clients must be able to send messages to the chat.<br>

> __05.__
Do not broadcast EMPTY messages from a client.<br>

> __06.__
Messages sent, must be identified by the time that was sent and the user name of who sent the message, example : [2020-01-20 15:48:41][client.name]:[client.message] .<br>

> __07.__
If a Client joins the chat, all the previous messages sent to the chat must be uploaded to the new Client.<br>

> __08.__
If a Client connects to the server, the rest of the Clients must be informed by the server that the Client joined the group.<br>

> __09.__
If a Client exits the chat, the rest of the Clients must be informed by the server that the Client left.<br>

> __10.__
All Clients must receive the messages sent by other Clients.<br>

> __11.__
If a Client leaves the chat, the rest of the Clients must not disconnect.<br>

> __12.__
If there is no port specified, then set as default the port 8989. Otherwise, program must respond with usage message: [USAGE]: ./TCPChat $port<br>


### INITIAL STARTUP

Clone the repo by running the following via terminal:  
> git clone "https://01.gritlab.ax/git/Steve/net-cat.git"  

Change root directory to this repo, then run the bash script "startup.sh" by entering the following in the terminal:  
> sh startup.sh 

If one prefers to specify a different port at startup, instead from the directory run:  

> go build -o TCPChat
> ./TCPChat [port]  

If usage does not correspond with "./TCPChat $port", a reminder on the terminal will be displayed.

After the above (which automatically starts the server on port 8989), the terminal should display a confirmation that the TCP server is up and running.  

Get the local computer's IP-address (via settings, preferences or similar, depending on OS). Now open a new terminal (either on the host computer, or another computer) and run:

> nc <IP-address of host computer> 8989

In order to change one's username at any time during the chat, simply type the command below into the terminal followed by ENTER.  

> /cn  

The user will then by prompted with [ENTER YOUR NAME]: . Duplicate names are not accepted, and will force a re-prompt for a new name to be entered.
