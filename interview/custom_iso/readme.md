## 内容

 1.注入启动 Logo启动过程中显示 ASCII 图形 Logo（如品牌标识或欢迎信息）

 2.硬件识别绑定自动识别设备 SN（序列号）和 MAC 地址，生成唯一标识

 3.SN + MAC 双重绑定授权机制x 使用 dmidecode 获取 SN，ip link 获取 MAC，计算并对比 Hash，判断是否为授权设备未授权设备拒绝启动，进入维护 shell

 4.自定义 initramfs 启动脚本



## 准备

```shell
sudo pacman -S  squashfs-tools xorriso cpio gzip dmidecode rsync
```

```shel
weget https://mirror.rackspace.com/archlinux/iso/latest/archlinux-x86_64.iso
```



## 启动

```shell
chmod +x ./*.sh
./build.sh
```

```
.
├── archlinux-x86_64.iso
├── build.log
├── build.sh
├── clean.sh
├── initfs-init
├── overlay
│   ├── authorized
│   │   └── auth.hash
│   ├── boot-logo
│   │   └── logo.txt
│   ├── etc
│   ├── issue
│   └── usr
│       └── bin
│           └── dmidecode
└── readme.md
```

