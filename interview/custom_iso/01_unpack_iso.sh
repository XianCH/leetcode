#!/bin/bash
set -euo pipefail

#ISO_ORIG="${1:-arch-custom-luks.iso}"

ISO_ORIG="${1:-archlinux-x86_64.iso}"
WORKDIR="${2:-$(pwd)}"
OUT="$WORKDIR/output"
LOG="./build.log"
DATE=$(date '+%F_%T')

MNT="$WORKDIR/iso"
ISO_NEW="$WORKDIR/iso-new"
ROOTFS="$WORKDIR/airootfs"
INITFS_DIR="$WORKDIR/initfs"
NEW_INITRAMFS="$ISO_NEW/arch/boot/x86_64/initramfs-linux.img"
OVERLAY="$WORKDIR/overlay"

log() { echo "[$(date '+%T')] $*" | tee -a "$LOG"; }

require_cmd() {
    command -v "$1" >/dev/null 2>&1 || { echo "缺少依赖: $1"; exit 1; }
}

log "[*] 检查依赖..."
for cmd in unsquashfs mksquashfs xorriso cpio zcat gzip sha256sum; do
    require_cmd "$cmd"
done

log "[*] 创建目录..."
mkdir -p "$MNT" "$ISO_NEW" "$ROOTFS" "$OUT" "$INITFS_DIR"

LOG="$OUT/build.log"
touch "$LOG"
log "[*] 初始化日志文件完成"

log "[*] 校验原始ISO..."
sha256sum "$ISO_ORIG" | tee "$OUT/iso.sha256"

log "[*] 挂载 ISO..."
sudo mount -o loop "$ISO_ORIG" "$MNT"

# log "[*] 解包 squashfs..."
# unsquashfs -d "$ROOTFS" "$MNT/arch/x86_64/airootfs.sfs"

# log "[*] 备份并劫持 dmidecode..."
# echo "[*] 劫持 dmidecode"
# mv "$ROOTFS/usr/bin/dmidecode" "$ROOTFS/usr/bin/dmidecode.real"
# cp "$WORKDIR/overlay/usr/bin/dmidecode" "$ROOTFS/usr/bin/dmidecode"
# chmod +x "$ROOTFS/usr/bin/dmidecode"

log "[*] 合并 overlay 配置..."
rsync -a "$OVERLAY/etc/" "$ROOTFS/etc/" || log "[!] overlay/etc 为空或不存在"

log "[*] 解包 initramfs..."
cd "$INITFS_DIR"
cpio -idmv < "$MNT/arch/boot/x86_64/initramfs-linux.img"



