# LogFilesObfuscator
Basic log files obfuscation functionality. 2 API endpoints corresponding to 'user' and 'admin' to depict encryption of sensitive information.

INSTRUCTIONS

The functional endpoints are:
1. localhost:3000/user/:key/:value
2. localhost:3000/admin/:key/:value

'key' can be 'FileID' or 'OrderID' or 'Date', and the 'value' can be checked from the 'data.json' file inside 'logs'.
