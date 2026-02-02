<!-- Xrao Project -->

[![version](https://img.shields.io/badge/Xrao-Version%201.0-blue.svg?maxAge=259200)]()
[![os](https://img.shields.io/badge/Supported%20OS-Android-blue.svg)]()
[![license](https://img.shields.io/badge/License-MIT-blue.svg)]()

# Xrao
Xrao is a simple tool to help stabilize Android performance. <br>
It supports both rooted and non-rooted Android devices. <br>
It is recommended to run this tool using
[Brevent](https://play.google.com/store/apps/details?id=me.piebridge.brevent).

**Note:** <br>
Xrao behavior depends on the device and system environment. <br>
Performance improvements may vary — some devices may see noticeable changes, while others may experience minimal or no difference.

## Requirements
- Android 14 or higher (rooted or non-rooted)
- Wireless ADB support
- Toybox or BusyBox

## Features
- Thermal throttling stabilization
- Background app termination
- Resolution downscaling
- Forced FPS locking
- Force enable ANGLE
- Enable dynamic updatable graphics driver
- Touch response booster
- Low latency mode
- App whitelisting and more...

## Installation
1. Download the ZIP file.
2. Extract the archive.
3. Run the tool.

## Usage
```bash
sh xrao.sh --path="/path/to/Xrao" <options>
```

<strong>Available options:</strong>
<table>
    <tr>
        <th>Option</th>
        <th>Description</th>
    </tr>
    <tr>
        <td>--inject</td>
        <td>Apply performance tweaks</td>
    </tr>
    <tr>
        <td>--eject</td>
        <td>Remove applied tweaks and restore defaults</td>
    </tr>
    <tr>
        <td>--status</td>
        <td>Show current tweak status</td>
    </tr>
    <tr>
        <td>--show-config</td>
        <td>Display current configuration settings</td>
    </tr>
    <tr>
        <td>--help</td>
        <td>Show available commands and usage</td>
    </tr>
    <tr>
        <td>--version</td>
        <td>Show tool version</td>
    </tr>
</table>

**Example:**
```bash
sh /sdcard/Download/Xrao/xrao.sh --path="/sdcard/Download/Xrao" --inject
```

## Configuration
Xrao uses 5 configuration files inside the `/config/` directory to control its behavior.

<table>
    <tr>
        <th>File</th>
        <th>Description</th>
    </tr>
    <tr>
        <td>xrao_settings.sh</td>
        <td>
            Basic settings such as FPS, downscaling, animations, and more.<br>
            These settings are not global and only apply to apps listed in affected_apps.sh.
        </td>
    </tr>
    <tr>
        <td>affected_apps.sh</td>
        <td>Configure which applications are affected by Xrao.</td>
    </tr>
    <tr>
        <td>kill_apps.sh</td>
        <td>Define which applications will be terminated.</td>
    </tr>
    <tr>
        <td>inject_settings.sh</td>
        <td>Configure actions that will be applied during injection.</td>
    </tr>
    <tr>
        <td>eject_settings.sh</td>
        <td>Configure actions that will be applied during ejection.</td>
    </tr>
</table>

## Warning
Xrao interacts with system-level components and may affect device behavior. <br>
Improper usage may cause instability or unexpected results. <br>
Use at your own risk. <br>
For more detailed information, please read the `DISCLAIMER` file.

<!-- Copyright (c) 2026 Zeronetsec -->