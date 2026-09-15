import re
from datetime import datetime
from pathlib import Path

script_dir = Path(__file__).resolve().parent
project_root = script_dir.parent.parent

today = datetime.now().strftime("%d%m%Y")
NEW_VERSION_STR = f"v0.1.{today}"

TARGET_FILE = project_root / "module" / "version" / "show_version.go"

KEYWORD = "version = "
VAL_WRAPPER = f'"{NEW_VERSION_STR}"'

try:
    content = TARGET_FILE.read_text(encoding="utf-8")

    pattern = rf"^(\s*){re.escape(KEYWORD)}\s*.*$"
    replacement = rf"\1{KEYWORD}{VAL_WRAPPER}"

    updated_content = re.sub(
        pattern, replacement,
        content, flags=re.MULTILINE,
    )

    TARGET_FILE.write_text(
        updated_content,
        encoding="utf-8",
    )

    print(f"\x1b[0;32m[+] \x1b[0mUpdating version: \x1b[0;32m{TARGET_FILE} \x1b[1;90m(\x1b[0;36m{KEYWORD}{VAL_WRAPPER}\x1b[1;90m)\x1b[0m")

except FileNotFoundError:
    print(f"\x1b[1;31m[!] \x1b[0mFile: \x1b[0;32m{TARGET_FILE} \x1b[0mnot found!")