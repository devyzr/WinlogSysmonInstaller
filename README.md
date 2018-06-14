# Sysmon and Winlogbeat Installer

This is a Go program that will let you easily install a working Sysmon/Winlogbeat setup on an endpoint. This works with ELK set up how [Ryan Watsons' fantastic tutorial](https://silentbreaksecurity.com/windows-events-sysmon-elk/) shows us.

Credit goes to Mathias Lafeldt for his awesome [Embedding Assets in Go](https://blog.codeship.com/embedding-assets-in-go/) tutorial.

## Requierements:
- Go Lang
- [go-bindata](https://github.com/go-bindata/go-bindata) (make sure the go/bin/ directory is in your PATH after installing if you're using windows)
- Powershell on the endpoint you'll be installing Sysmon/Winlogbeat on. If you're using XP you'll probably have to install it.
- Sysmon.exe, it's not included due to licencing reasons. You can download it from [here](https://live.sysinternals.com/Sysmon.exe).

## Usage:

- Clone this repository using `git clone` or download and uncompress the zip.
- Download Sysmon.exe from https://live.sysinternals.com/Sysmon.exe, and put it into the assets folder. It's not included due to licensing reasons.
- Update line 107 in the winlogbeat.yml file in the assets directory so that your IP matches your ELK server and make any other necessary adjustments.
- Customize [Swift on Security](https://github.com/SwiftOnSecurity/sysmon-config)'s sysmonconfig-export.xml that's in the assets folder according to your needs, it's already set up very well though and isn't strictly necessary.
- Copy the cert you generated in the tutorial into the assets directory, make sure it's called 'ELK-Stack.crt'.
- Run `go generate` and then `go build` in the project folder. (Note: I'm building this in windows, if you're doing it in linux you're probably going to have to tell Go to output an .exe)
- Test on an endpoint!

## To Do:

- Make an Uninstaller.