### Usage

1. Configure connection information in the `autoConnectVPN.ini` file:
    ```ini
    [vpn]
    ; Your VPN username
    username=user
    ; Your VPN password
    password=pwd
    ; The "Connection Entry" name in the VPN Client program
    profile=TEST_VPN
    
    [options]
    ; Program will exit after n seconds
    ; Number of seconds before program exits
    ; Default is 0 seconds (program will not exit)
    exit_after_seconds=0

    ; Disconnect existing VPN connection before connecting
    ; Set to 1 to enable, 0 to disable
    ; Default is 1
    disconnect_before_connect=1
    ```
    
2. Run the `autoConnectVPN.exe` program.
