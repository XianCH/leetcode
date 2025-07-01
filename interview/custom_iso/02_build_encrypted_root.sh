
#!/bin/bash
set -euo pipefail

WORKDIR="${1:-$(pwd)}"
IMG="$WORKDIR/root.img"
KEYFILE="$WORKDIR/overlay/authorized/keyfile.bin"
MAPPER_NAME="luks-root"
MOUNTPOINT="$WORKDIR/mnt-rootfs"

mkdir -p "$MOUNTPOINT"


if [ -e "/dev/mapper/$MAPPER_NAME" ]; then
  echo "[*] 检测到已存在映射 $MAPPER_NAME，正在关闭..."
  cryptsetup close "$MAPPER_NAME"
fi

if [ ! -f "$KEYFILE" ]; then
  echo "[*] 生成 keyfile: $KEYFILE"
  dd if=/dev/urandom of="$KEYFILE" bs=1 count=32
  chmod 600 "$KEYFILE"
else

  echo "[*] 使用已存在的 keyfile: $KEYFILE"
fi

echo "[*] 创建 root.img..."
dd if=/dev/zero of="$IMG" bs=1M count=2048

echo "[*] LUKS 格式化 root.img（使用 keyfile）..."
cryptsetup luksFormat "$IMG" --batch-mode --key-file "$KEYFILE"

echo "[*] 打开加密 root.img..."
cryptsetup open "$IMG" "$MAPPER_NAME" --key-file "$KEYFILE"

echo "[*] 格式化加密设备为 ext4..."
mkfs.ext4 /dev/mapper/$MAPPER_NAME

echo "[*] 构建 Arch 根文件系统..."
mount /dev/mapper/$MAPPER_NAME "$MOUNTPOINT"
pacstrap -c -G -M "$MOUNTPOINT" base linux linux-firmware cryptsetup mkinitcpio

echo "[*] 拷贝自定义 init（含解密逻辑）..."
cp "$WORKDIR/initfs-init" "$MOUNTPOINT/sbin/init"
chmod +x "$MOUNTPOINT/sbin/init"

umount "$MOUNTPOINT"
cryptsetup close "$MAPPER_NAME"


echo "[*] 验证生成的 root.img 是否是 LUKS 加密卷..."
cryptsetup isLuks "$ROOT_IMG" && echo " 是有效的 LUKS 镜像"
