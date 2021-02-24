# speaker

Переводит текст при нажатии ctr+c, и читает его в слух с помощью приложение для android использую tts установленные там.

- Перед использованием нужно получить приложение для android https://github.com/exitstop/speakerandroid
- Скомпилировать и установить на android

# Как работает переводичк

Перевод происходит через https://translate.google.ru с помощью https://github.com/mxschmitt/playwright-go, запуская браузер в headless режиме

# requirement

```bash
sudo apt-get install -y libx11-dev libxtst-dev libxt-dev libxinerama-dev libx11-xcb-dev libxkbcommon-dev libxkbcommon-x11-dev libxkbfile-dev
```

# Как запустить

```bash
go run cmd/voice/main.go -ip 192.168.0.133 -t
```

# Копиляция под windows

```bash
make docker_prebuild_image
make windows
# -> build/speaker.exe
```

# Глобальные горячие клавиши

- `ctrl+c` перевести и прочитать текст
- `alt+t` читать текст но не переводить(вкл/выкл)
- `ctrl+alt+p` pause/resume.
- `ctr+shift+q` завершить приложение

# find only dir

```bash
find . -type d | cut -c 3- | sed 's/^/!*/' | grep -v ".git"
```


# Fix go mod build private repo

bash
git config --global url."ssh://git@github.com".insteadOf "https://github.com"

##### ~/.zshrc
```bash
export GO111MODULE=on
export GOPROXY=direct
export GOSUMDB=off
source ~/.zshrc
```
