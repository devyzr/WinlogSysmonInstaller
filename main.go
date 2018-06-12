package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	installDir := "C:\\SysLogBeat\\"
	var err error

	err = makeInstallDir(installDir)
	if err == nil {
		log.Printf("Created '%v' successfully.\n\n", installDir)
	} else {
		log.Println(err)
		os.Exit(1)
	}

	err = unpackSysmon(installDir)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = installSysmon(installDir)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = unpackWinlogbeat(installDir)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = installWinlogbeat(installDir)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Winlogbeat installation finalized. Check kibana server for logs.")
}

func makeInstallDir(installDir string) (err error) {
	_, err = os.Open(installDir)
	if err == nil {
		log.Printf("%v already exists, check if already installed.\n", installDir)
		return
	}

	err = os.Mkdir(installDir, 0761)
	if err != nil {
		log.Printf("Error creating %v.\n", installDir)
		return
	}

	_, err = os.Open(installDir)
	if err != nil {
		log.Println("Error accessing %v after creation.\n", installDir)
		return
	}

	return
}

func loadAsset(installPath string, assetName string) (err error) {
	// Unpack asset
	asset_install_path := installPath + "\\" + assetName
	log.Printf("Loading %v from assets.\n", assetName)
	data, err := Asset("assets/" + assetName)
	if err != nil {
		log.Println("Problem loading %v from assets.", assetName)
		return
	}

	// Write asset
	err = ioutil.WriteFile(asset_install_path, []byte(data), 0644)
	if err != nil {
		log.Println("There was a problem writing %v.\n", assetName)
		return
	} else {
		log.Printf("%v written successfully.\n\n", assetName)
		return
	}
}

func unpackSysmon(installDir string) (err error) {
	// Create Sysmon dir
	sysmonDir := installDir + "Sysmon"
	err = os.Mkdir(sysmonDir, 0761)
	if err != nil {
		log.Printf("Error creating %v.\n", sysmonDir)
		return
	} else {
		log.Printf("Successfully created %v.\n", sysmonDir)
	}

	// Unpack Sysmon and config file
	err = loadAsset(sysmonDir, "Sysmon.exe")
	if err != nil {
		return
	}
	err = loadAsset(sysmonDir, "sysmonconfig-export.xml")
	return
}

func installSysmon(installDir string) (err error) {
	sysmonExecutable := installDir + "Sysmon\\Sysmon.exe"
	sysmonConfig := installDir + "Sysmon\\sysmonconfig-export.xml"
	// Start sysmon service.
	// Here we call cmd.exe and pass the "/C" flag to indicate we're passing a command.
	runSysmon := exec.Command("cmd", "/C", sysmonExecutable, "-i", sysmonConfig, "-accepteula", "-h", "md5,sha256", "-n", "-l")
	err = runSysmon.Run()
	if err != nil {
		log.Println("Error starting Sysmon.")
		log.Println("If exit status is 1242, it probably means you need to uninstall sysmon with 'Sysmon.exe -u'.")
		return
	} else {
		log.Println("Successfully started Sysmon.")
	}

	// Install sysmon service.
	// sc config Sysmon start= auto
	installSysmon := exec.Command("cmd", "/C", "sc", "config", "Sysmon", "start=", "auto")
	err = installSysmon.Run()
	if err != nil {
		log.Println("Error installing Sysmon service.")
		return
	} else {
		log.Println("Successfully installed Sysmon service.\n")
	}
	return
}

func unpackWinlogbeat(installDir string) (err error) {
	// Create Winlogbeat dir
	winlogDir := installDir + "Winlogbeat"
	err = os.Mkdir(winlogDir, 0761)
	if err != nil {
		log.Printf("Error creating %v.\n", winlogDir)
		return
	} else {
		log.Printf("Successfully created %v.\n", winlogDir)
	}

	err = loadAsset(winlogDir, "winlogbeat.exe")
	if err != nil {
		return
	}
	err = loadAsset(winlogDir, "winlogbeat.yml")
	if err != nil {
		return
	}
	err = loadAsset(winlogDir, "fields.yml")
	if err != nil {
		return
	}
	err = loadAsset(winlogDir, "install-service-winlogbeat.ps1")
	if err != nil {
		return
	}
	err = loadAsset(winlogDir, "ELK-Stack.crt")
	return
}

func installWinlogbeat(installDir string) (err error) {
	winlogDir := installDir + "Winlogbeat"
	winlogSvc := winlogDir + "\\install-service-winlogbeat.ps1"
	// Run Winlogbeat installer
	installWLB := exec.Command("powershell", "-Exec", "bypass", "-File", winlogSvc)
	err = installWLB.Run()
	if err != nil {
		log.Println("Error installing Winlogbeat.")
		return
	} else {
		log.Println("Winlogbeat installed successfully.")
	}

	// Install winlogbeat service
	installWLBSvc := exec.Command("powershell", "Set-Service", "-Name", "\"winlogbeat\"", "-StartupType", "automatic")
	err = installWLBSvc.Run()
	if err != nil {
		log.Println("Error installing Winlogbeat service.")
		return
	} else {
		log.Println("Winlogbeat service installed successfully.")
	}

	// Start winlogbeat service
	startWLBSvc := exec.Command("powershell", "Start-Service", "-Name", "\"winlogbeat\"")
	err = startWLBSvc.Run()
	if err != nil {
		log.Println("Error starting Winlogbeat service.")
		return
	} else {
		log.Println("Winlogbeat service started successfully.\n")
	}
	return
}
