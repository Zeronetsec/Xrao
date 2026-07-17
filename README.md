<!-- https://github.com/Zeronetsec/Xrao -->

<div style="display: flex; gap: 8px; justify-content: left;">
    <img src="https://img.shields.io/badge/Xrao-Version%200.1-blue.svg?style=square&logo=go" />
    <img src="https://img.shields.io/badge/Supported%20OS-Android-blue.svg?style=square&logo=android" />
    <img src="https://img.shields.io/badge/License-MIT-blue.svg?style=square&logo=github" />
</div>

# Xrao
Xrao is a declarative, indentation-based configuration DSL designed for structured Android system tweaking and automation via ADB.

## Features
- Parse declarative indentation-based DSL configuration rules.
- Apply full configuration rules to the Android system over ADB.
- Deploy isolated global-only rules from the configuration file.
- Revert all applied modifications and reset system states completely.
- And more features.

## Disclaimer
Please read [.docs/disclaimer.md](.docs/disclaimer.md) before using this tool. </br>
Use this software at your own risk. </br>
The author is not responsible for any damage, data loss, or issues that may result from its use.

## Installation
Quick install:
```bash
git clone https://github.com/Zeronetsec/Xrao
cd Xrao
chmod +x install.sh
./install.sh
```
For more detailed installation and uninstallation instructions, see [.docs/install_and_uninstall.md](.docs/install_and_uninstall.md).

## Usage Example
```bash
xrao --pair 192.168.x.x:5555
xrao --connect 192.168.x.x:5555
xrao --apply
xrao --generate --mode apply-global-only --config ~/mycustom_config.xr --out ~/mytweak.sh
xrao --status
```
And more commands.

## License
This project is licensed under the MIT License.

<!-- Copyright (c) 2026 Zeronetsec -->