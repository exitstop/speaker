# speaker

# find only dir

```bash
find . -type d
```

# requirement

```bash
sudo apt-get install -y libx11-dev libxtst-dev libxt-dev libxinerama-dev libx11-xcb-dev libxkbcommon-dev libxkbcommon-x11-dev libxkbfile-dev
```

# Always on bottom
# https://askubuntu.com/questions/351720/how-to-start-an-application-with-bottom-most-property
sudo apt-get install devilspie4
vim ~/.config/devilspie2/chromium.lua
```bash
debug_print("Window Name: " .. get_window_name())
debug_print("Application name: " .. get_application_name())
debug_print("WM_CLASS: " .. get_class_instance_name())
debug_print("Window Class: " .. get_window_class())
if (string.match(get_application_name(),"Chrome$")) then
   set_window_below();
end
```

devilspie2 --debug
