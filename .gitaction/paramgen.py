#!/usr/bin/env python3

import re
from pathlib import Path

def generate_uiflag_params():
    script_dir = Path(__file__).resolve().parent
    project_root = script_dir.parent

    target_md = project_root / ".docs" / "install_and_uninstall.md"
    output_flg = script_dir / "uiflag_params.flg"

    if not target_md.exists():
        print(f"[!] File: {target_md} not found!")
        return

    with open(target_md, "r", encoding="utf-8") as f:
        lines = f.readlines()

    sections = {"installation": [], "uninstallation": []}
    current_section = None

    for line in lines:
        stripped = line.strip()
        if re.match(
            r"^#\s+Installation",
            stripped,
            re.IGNORECASE,
        ):
            current_section = "installation"
            continue
        elif re.match(
            r"^#\s+Uninstallation",
            stripped,
            re.IGNORECASE,
        ):
            current_section = "uninstallation"
            continue
        elif stripped.startswith("#"):
            current_section = None
            continue

        if current_section and stripped:
            sections[current_section].append(stripped)

    def format_section(raw_lines):
        formatted = []
        for line in raw_lines:
            flag_match = re.match(
                r"^-\s*`([^`]+)`", line,
            )
            if flag_match:
                flag_name = flag_match.group(1)
                formatted.append(
                    f"    \\x1b[1;90m* \\x1b[0;32m{flag_name}\\x1b[0m",
                )
                continue

            tree_match = re.match(
                r"^-\s*└──\s*(.*)", line,
            )
            if tree_match:
                desc = tree_match.group(1).replace("`", "")
                formatted.append(
                    f"    \\x1b[1;90m└── \\x1b[0;37m{desc}\\x1b[0m",
                )
                continue

            if line.startswith("`"):
                formatted_line = re.sub(
                    r"`([^`]+)`\s*",
                    r"\\x1b[0;32m\1 \\x1b[0m",
                    line,
                )
                formatted.append(formatted_line)
                continue
        return "\n".join(formatted)

    iflags = format_section(sections["installation"])
    uflags = format_section(sections["uninstallation"])

    output_content = f'echo -e "Iflags:\n{iflags}\n\nuflags:\n{uflags}"\n'

    output_flg.parent.mkdir(parents=True, exist_ok=True)
    with open(output_flg, "w", encoding="utf-8") as f:
        f.write(output_content)

    print(f"[+] Generated: {output_flg}")

if __name__ == "__main__":
    generate_uiflag_params()