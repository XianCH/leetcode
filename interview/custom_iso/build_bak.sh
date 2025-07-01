#!/bin/bash
set -euo pipefail

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

#清理挂载和临时目录
#   cleanup() {
#       log "[!] 清理操作..."
#       if mountpoint -q "$MNT"; then
#           sudo umount "$MNT"
#       fi
#       rm -rf "$MNT" "$ISO_NEW" "$ROOTFS" "$INITFS_DIR"
#       log "[!] 清理完成"
#   }
#
# trap cleanup EXIT ERR

log() { echo "[$(date '+%T')] $*" | tee -a "$LOG"; }

require_cmd() {
    command -v "$1" >/dev/null 2>&1 || { echo "缺少依赖: $1"; exit 1; }
}

log "[*] 检查依赖..."
for cmd in unsquashfs mksquashfs xorriso cpio zcat gzip sha256sum; do
    require_cmd "$cmd"
done

# log "[*] 清除旧文件..."
# sudo rm -rf "$MNT" "$ISO_NEW" "$ROOTFS" "$OUT" "$INITFS_DIR"

log "[*] 创建目录..."
mkdir -p "$MNT" "$ISO_NEW" "$ROOTFS" "$OUT" "$INITFS_DIR"

LOG="$OUT/build.log"
touch "$LOG"
log "[*] 初始化日志文件完成"

log "[*] 校验原始ISO..."
sha256sum "$ISO_ORIG" | tee "$OUT/iso.sha256"

log "[*] 挂载 ISO..."
sudo mount -o loop "$ISO_ORIG" "$MNT"

log "[*] 解包 squashfs..."
unsquashfs -d "$ROOTFS" "$MNT/arch/x86_64/airootfs.sfs"

log "[*] 备份并劫持 dmidecode..."
echo "[*] 劫持 dmidecode"
mv "$ROOTFS/usr/bin/dmidecode" "$ROOTFS/usr/bin/dmidecode.real"
cp "$WORKDIR/overlay/usr/bin/dmidecode" "$ROOTFS/usr/bin/dmidecode"
chmod +x "$ROOTFS/usr/bin/dmidecode"

log "[*] 合并 overlay 配置..."
rsync -a "$OVERLAY/etc/" "$ROOTFS/etc/" || log "[!] overlay/etc 为空或不存在"

log "[*] 解包 initramfs..."
cd "$INITFS_DIR"
cpio -idmv < "$MNT/arch/boot/x86_64/initramfs-linux.img"

log "[*] 替换 init 并注入授权..."
cp "$WORKDIR/initfs-init" "$INITFS_DIR/init"
chmod +x "$INITFS_DIR/init"
mkdir -p "$INITFS_DIR/authorized"
if [ -d "$OVERLAY/authorized" ] && [ "$(ls -A "$OVERLAY/authorized")" ]; then
    cp "$OVERLAY/authorized/"* "$INITFS_DIR/authorized/"
else
    log "[!] overlay/authorized 为空或不存在"
fi

mkdir -p "$(dirname "$NEW_INITRAMFS")"

log "[*] 重新打包 initramfs..."
cd "$INITFS_DIR"
find . | cpio -o -H newc > "$NEW_INITRAMFS"
cd "$WORKDIR"

log "[*] 重新打包 squashfs..."
mksquashfs "$ROOTFS" "$OUT/airootfs.sfs" -noappend -comp xz

 log "[*] 准备 ISO 文件结构..."
 rsync -a "$MNT/" "$ISO_NEW/"
 cp "$OUT/airootfs.sfs" "$ISO_NEW/arch/x86_64/"


log "[*] 生成新ISO..."
xorriso -as mkisofs \
  -iso-level 3 -full-iso9660-filenames \
  -volid "ARCH_CUSTOM_GPG" \
  -output "$OUT/arch-custom-gpg.iso" \
  -isohybrid-mbr /usr/lib/syslinux/bios/isohdpfx.bin  \
  -c boot/syslinux/boot.cat \
  -b boot/syslinux/isolinux.bin \
    -no-emul-boot -boot-load-size 4 -boot-info-table \
  -eltorito-alt-boot \
    -no-emul-boot \
  "$ISO_NEW"
