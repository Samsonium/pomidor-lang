#!/usr/bin/env python3
"""
🍅 Pomidor Language Interpreter v0.1
Executes files with the .pmd extension.
Author: https://github.com/vldanch/
"""

import sys
import re

VERSION = "0.1"

class PomidorError(Exception):
    """Custom exception for Pomidor interpreter errors."""
    pass

def panic(message, line=0):
    """Terminate the interpreter with an error message."""
    print("🍅 POMIDOR PANIC")
    if line:
        print(f"Line {line}: {message}")
    else:
        print(message)
    sys.exit(1)

def run(filename):
    if not filename.endswith(".pmd"):
        raise PomidorError("Pomidor files must have the .pmd extension")

    try:
        with open(filename, "r", encoding="utf-8") as f:
            lines = f.readlines()
    except FileNotFoundError:
        panic(f"File not found: {filename}")

    garden_declared = False
    in_pluck = False

    print(f"🍅 Pomidor Language Interpreter v{VERSION}")
    print(f"Running file: {filename}\n")

    for lineno, raw in enumerate(lines, start=1):
        line = raw.strip()

        if not line or line.startswith("//"):
            continue

        # garden HelloPomidor
        if line.startswith("garden"):
            if garden_declared:
                panic("garden already declared", lineno)
            garden_declared = True
            continue

        # pluck <name>() {
        if line.startswith("pluck ") and line.endswith("() {"):
            if in_pluck:
                panic("nested pluck not allowed", lineno)
            in_pluck = True
            continue

        # closing brace
        if line == "}":
            if not in_pluck:
                panic("unexpected closing brace", lineno)
            in_pluck = False
            continue

        # inside pluck
        if in_pluck:
            # see "text"
            if line.startswith("see"):
                match = re.match(r'see\s+"(.*)"', line)
                if not match:
                    panic("invalid see syntax", lineno)
                print(match.group(1))
            else:
                panic(f"unknown statement: {line}", lineno)

    if not garden_declared:
        panic("missing garden declaration")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(f"🍅 Pomidor Interpreter v{VERSION}")
        print("Usage: python3 pomidor_interpreter.py <file.pmd>")
        sys.exit(1)

    run(sys.argv[1])
