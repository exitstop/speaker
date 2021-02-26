# speaker

# Как установить Golang

```bash
GOVERSION="1.16"
wget -P /tmp -q https://dl.google.com/go/go$GOVERSION.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf /tmp/go$GOVERSION.linux-amd64.tar.gz
```

# Вариант использующий голос gTTS, не требующий голосового движка
# Самый простой способ запустить

```bash
sudo -H pip3 install gTTS; sudo apt install -y mpg123
go run cmd/gtts/main.go -t
```

# Вариант с отдельным приложение на android
## Описание

`cmd/android/main.go` Переводит текст при нажатии ctr+c, и читает его в слух с помощью приложение для android использую tts установленные там.

- Перед использованием нужно получить приложение для android https://github.com/exitstop/speakerandroid
- Скомпилировать и установить на android

# Как работает переводчик

Перевод происходит через https://translate.google.ru с помощью https://github.com/mxschmitt/playwright-go, запуская браузер в headless режиме

# requirement

```bash
sudo apt-get install -y libx11-dev libxtst-dev libxt-dev libxinerama-dev libx11-xcb-dev libxkbcommon-dev libxkbcommon-x11-dev libxkbfile-dev
```

# Как запустить

```bash
Вариант 1.
# берет голос gtts
sudo -H pip3 install gTTS; sudo apt install -y mpg123
go run cmd/gtts/main.go -t

Вариант 2.
# берет голос из android приложения 
go run cmd/android/main.go -ip 192.168.0.133 -t
```

# Компиляция под windows

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

# Бесплатные голоса для ubuntu

```bash
sudo apt-get install gnustep-gui-runtime
say "hello"

sudo apt-get install festival
echo "hello" | festival --tts

sudo apt-get install speech-dispatcher
spd-say "hello"

sudo apt-get install espeak
espeak "hello"

# https://gtts.readthedocs.io/en/latest/
sudo -H pip3 install gTTS
(Google Text to Speech / github.com/pndurette/gTTS )
gtts-cli -h
gtts-cli -l ru "привет мир" | mpg123 -
gtts-cli -l ru "Бесплатный сервис Google позволяет мгновенно переводить слова, фразы и веб-страницы с английского более чем на 100 языков и обратно." | mpg123 -d 3 --pitch 0 -
# https://linux.die.net/man/1/mpg123
 

```

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

# Аналоги

- https://github.com/soimort/translate-shell

```bash
git clone https://github.com/soimort/translate-shell
cd translate-shell
sudo apt install -y gawk
make
sudo make install
trans en:ru -brief 'Hello world' -p 

# https://github.com/soimort/translate-shell/wiki/Narrator-Selection
trans -e yandex "Ничего, были бы кости, а мясо будет" -sp -n jane,evil
```
