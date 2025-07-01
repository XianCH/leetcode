#!/bin/bash
set -euo pipefail

WORKDIR="${1:-$(pwd)}"
ISO_NEW="$WORKDIR/iso-new"
MNT="$WORKDIR/iso"
ROOTFS="$WORKDIR/airootfs"
OUT="$WORKDIR/output"
INITFS_DIR="$WORKDIR/initfs"
ROOT_IMG="$WORKDIR/root.img"
NEW_INITRAMFS="$ISO_NEW/arch/boot/x86_64/initramfs-linux.img"

mkdir -p "$ISO_NEW" "$OUT"

log() { echo "[$(date '+%T')] $*" | tee -a "$OUT/build.log"; }



log "[*] 替换 init"
cp "$WORKDIR/initfs-init" "$INITFS_DIR/init"
chmod +x "$INITFS_DIR/init"

log "[*] 注入解密用文件"
cp "$WORKDIR/authorized/"* "$INITFS_DIR/authorized/" 2>/dev/null || true

mkdir -p "$(dirname "$NEW_INITRAMFS")"

log "[*] 重新打包 initramfs"
cd "$INITFS_DIR"
find . | cpio -o -H newc | gzip > "$NEW_INITRAMFS"
cd "$WORKDIR"

# log "[*] 重新打包 squashfs..."
# mksquashfs "$ROOTFS" "$OUT/airootfs.sfs" -noappend -comp xz
#

log "[*] 组装 ISO 文件结构..."
rsync -a "$MNT/" "$ISO_NEW/"
# cp "$OUT/airootfs.sfs" "$ISO_NEW/arch/x86_64/"
cp "$ROOT_IMG" "$ISO_NEW/arch/root.img"

log "[*] 生成 ISO..."
xorriso -as mkisofs \
  -iso-level 3 -full-iso9660-filenames \
  -volid "ARCH_CUSTOM_ENC" \
  -output "$OUT/arch-custom-luks.iso" \
  -isohybrid-mbr /usr/lib/syslinux/bios/isohdpfx.bin \
  -c boot/syslinux/boot.cat \
  -b boot/syslinux/isolinux.bin \
    -no-emul-boot -boot-load-size 4 -boot-info-table \
  -eltorito-alt-boot \
    -no-emul-boot \
  "$ISO_NEW"


echo "[*] 验证生成的 root.img 是否是 LUKS 加密卷..."
cryptsetup isLuks "$ROOT_IMG" && echo "✅ 是有效的 LUKS 镜像"
