#!/bin/bash
set -euo pipefail

WORKDIR="${1:-$(pwd)}"

MNT="$WORKDIR/iso"
ISO_NEW="$WORKDIR/iso-new"
ROOTFS="$WORKDIR/airootfs"
INITFS_DIR="$WORKDIR/initfs"
OUT="$WORKDIR/output"
LOG="$OUT/build.log"

info()    { echo -e "\033[1;32m[INFO]\033[0m $*"; }
warn()    { echo -e "\033[1;33m[WARN]\033[0m $*"; }
error()   { echo -e "\033[1;31m[ERR ]\033[0m $*"; }
log()     { echo "[$(date '+%T')] $*" | tee -a "$LOG"; }

cleanup() {
    info "开始清理构建中间文件..."

    if mountpoint -q "$MNT"; then
        warn "$MNT 当前已挂载，尝试卸载..."
        sudo umount "$MNT" && info "卸载成功: $MNT" || error "卸载失败: $MNT"
    else
        info "$MNT 未挂载，无需卸载"
    fi

    for path in "$MNT" "$ISO_NEW" "$ROOTFS" "$INITFS_DIR" "$OUT"; do
        if [ -e "$path" ]; then
            rm -rf "$path"
            info "已删除目录: $path"
        else
            warn "目录不存在: $path"
        fi
    done

    info "清理完成"
}

cleanup
