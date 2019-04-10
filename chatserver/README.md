# chatserver
A chat server implementation using golang.

## Project details:
1. Create a multi-threaded server which accepts requests from different clients.
2. Client Properties: every client will have ID and Group ID.
   Note: we can maintain Max-Group a global configurable value.
3. Server will Accept clients count of Max-Group for each Group ID
4. Whenever a message/packet is received from one client X with group ID Y, it should be Processed (rev the message and send to complete group Y except sender X)
5. Also implement the Wheel Timers for ping (time can be configured). Here server will be sending the ping for every 10 sec to all the clients, if response is not received from the client then packet will be dropped thereafter, from that specific client (that means server will be disconnect from that client).