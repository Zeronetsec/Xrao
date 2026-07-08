<!-- https://github.com/Zeronetsec/Xrao -->

# Installation
`install.sh` optional option:
- `--backup`

Use `--backup` to create a backup of the existing Xrao installation before replacing it.

## Termux
```bash
git clone https://github.com/Zeronetsec/Xrao
cd Xrao
chmod +x install.sh
./install.sh
```

## Uninstallation
```bash
export prefix="${PREFIX:-/usr}"
rm -f "${prefix}/bin/xrao"
rm -rf "${prefix}/opt/xrao"
rm -rf "${HOME}/.xrao"
```

<!-- Copyright (c) 2026 Zeronetsec -->