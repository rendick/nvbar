# nvbar - a bar

![image](https://github.com/rendick/nvbar/assets/91757099/aff0f6fa-ecba-4a02-9640-c97308b7572b)


```bash
git clone https://github.com/rendick/nvbar.git
sh install.sh
nvbar
```

```bash
bar {
        status_command nvbar 
	position top
	mode dock
	separator_symbol  " | "
        modifier None
        colors {

                separator #FF00FF
                background #19132c
                statusline #dddddd
                focused_workspace #ffffff #b05883
                active_workspace  #9400D3 #9400D3 #ffffff
                inactive_workspace #1B1B1B #1B1B1B #443365
                urgent_workspace #2f343a #696969 #000000
    }
}
```

# To-Do

- [x] Connection
- [x] Memory
- [x] Time
- [ ] CPU usage
- [x] Keyboard layout
- [x] Battery

- [ ] Colors
