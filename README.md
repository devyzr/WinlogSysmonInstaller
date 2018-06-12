# Sysmon and Winlogbeat Installer

This is a Go program that will let you easily install a working Sysmon/Winlogbeat setup on an endpoint. This works with ELK set up how [Ryan Watsons' fantastic tutorial](https://silentbreaksecurity.com/windows-events-sysmon-elk/) shows us.

Credit goes to Mathias Lafeldt for his awesome [Embedding Assets in Go](https://blog.codeship.com/embedding-assets-in-go/) tutorial.

## Requierements:
- Go Lang
- [go-bindata](https://github.com/go-bindata/go-bindata) (make sure the go/bin/ directory is in your PATH after installing if you're using windows)

## Usage:

- Update line 107 in the winlogbeat.yml file in the assets directory so that your IP matches your ELK server and make any other necessary adjustments.
- Customize [Swift on Security](https://github.com/SwiftOnSecurity/sysmon-config)'s sysmonconfig-export.xml that's in the assets folder according to your needs, it's already set up very well though and isn't strictly necessary.
- Run `go generate` and then `go build` in the project folder.
- Test on an endpoint!

## To Do:

- Make an Uninstaller.