# speaker

# Как запустить

```bash
go run cmd/speaker/main.go -ip 192.168.0.133 -t
```

# find only dir

```bash
find . -type d
```

# requirement

```bash
sudo apt-get install -y libx11-dev libxtst-dev libxt-dev libxinerama-dev libx11-xcb-dev libxkbcommon-dev libxkbcommon-x11-dev libxkbfile-dev
```

# Always on bottom
- https://askubuntu.com/questions/351720/how-to-start-an-application-with-bottom-most-property
```bash
sudo apt-get install devilspie4
vim ~/.config/devilspie2/chromium.lua
```

```bash
sudo apt-get install devilspie4
debug_print("Window Name: " .. get_window_name())
debug_print("Application name: " .. get_application_name())
debug_print("WM_CLASS: " .. get_class_instance_name())
debug_print("Window Class: " .. get_window_class())
if (string.match(get_application_name(),"Chrome$")) then
   set_window_below();
end
```

devilspie2 --debug

# chromedriver скачать такую же версию как google chrome

- https://chromedriver.chromium.org/
```bash
wget https://chromedriver.storage.googleapis.com/86.0.4240.22/chromedriver_linux64.zip
extract extract chromedriver_linux64.zip
cd chromedriver_linux64
sudo mv chromdriver /usr/local/bin/
```
# Fix go mod build private repo

bash
git config --global url."ssh://git@github.com".insteadOf "https://github.com"

##### ~/.zshrc
```bash
export GO111MODULE=on
export GOPROXY=direct
export GOSUMDB=off
export GOPRIVATE="gitlab.com/FaceChainTeam"
source ~/.zshrc
```
