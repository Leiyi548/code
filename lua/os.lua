local uname = vim.loop.os_uname()
local os = uname.sysname
local release = uname.release
-- wsl2
print(os) -- linux
print(release) -- 5.10.16.3-microsoft-standard-WSL2
