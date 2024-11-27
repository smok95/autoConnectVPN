package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/micmonay/keybd_event"
	"gopkg.in/ini.v1"
)

func main() {
	// 설정 파일 열기
	cfg, err := ini.Load("autoConnectVPN.ini")
	if err != nil {
		fmt.Println("Error loading config file:", err)
		os.Exit(1)
	}

	// VPN 섹션에서 사용자 정보 가져오기
	section := cfg.Section("vpn")
	username := section.Key("username").String()
	password := section.Key("password").String()
	profile := section.Key("profile").String()

	section = cfg.Section("options")
	exitAfterSeconds, err := section.Key("exit_after_seconds").Uint()
	if err != nil {
		exitAfterSeconds = 0
	}

	disconnectBeforeConnect, err := section.Key("disconnect_before_connect").Uint()
	if err != nil {
		disconnectBeforeConnect = 1
	}

	// ProgramFiles(x86) 환경 변수 값 가져오기
	programFiles := os.Getenv("ProgramFiles(x86)")
	if programFiles == "" {
		fmt.Println("Error: ProgramFiles(x86) environment variable not set.")
		os.Exit(1)
	}

	// VPN 클라이언트 경로 설정
	vpnClientPath := programFiles + `\Cisco Systems\VPN Client`

	// vpnclient.exe가 지정된 경로에 있는지 확인
	_, err = os.Stat(vpnClientPath + `\vpnclient.exe`)
	if err != nil {
		fmt.Printf("Error: vpnclient.exe not found in %s.\nExiting...\n", vpnClientPath)
		os.Exit(1)
	}

	if disconnectBeforeConnect == 1 {
		// 기존 VPN 연결 끊기
		cmd := exec.Command(vpnClientPath+`\vpnclient`, "disconnect")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println("Error disconnecting from VPN:", err)
		}
	}

	cmd := exec.Command(vpnClientPath+`\vpnclient`, "connect", "cliauth", profile, "user", username, "pwd", password)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// vpnclient를 시작
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error connecting to VPN:", err)
		os.Exit(1)
	}

	// "y" 키 누르기
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Println("Error creating keybonding:", err)
		os.Exit(1)
	}
	kb.SetKeys(keybd_event.VK_Y)
	err = kb.Launching()
	if err != nil {
		fmt.Println("Error pressing key:", err)
		os.Exit(1)
	}

	if exitAfterSeconds > 0 {
		time.Sleep(time.Duration(exitAfterSeconds) * time.Second)

		kb.Clear()
		kb.SetKeys(keybd_event.VK_C)
		kb.HasCTRL(true)
		err = kb.Launching()
		if err != nil {
			fmt.Println("Error pressing key:", err)
			os.Exit(1)
		}
	}
	time.Sleep(15 * time.Second)

	// vpnclient 프로세스 종료
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting vpnclient process:", err)
		os.Exit(1)
	}
}
