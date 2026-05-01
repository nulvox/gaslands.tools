#!/usr/bin/env python3
"""Extract game data from Gaslands Refuelled PDF into structured markdown files.

Usage:
    uv run python extract.py --pdf <path-to-pdf> --output-dir <output-directory>

The extractor reads the Gaslands Refuelled rulebook PDF and produces:
    - full-text.md:  Raw text extraction (for review, gitignored)
    - vehicles.md:   Vehicle types with stats tables
    - weapons.md:    Weapons with costs, dice, range, slots
    - upgrades.md:   Upgrades with costs, slots, descriptions
    - perks.md:      Perks organized by class with costs
    - sponsors.md:   Sponsors with perk classes and sponsor perks

After running, manually review and correct the output files against the
rulebook. The corrected data feeds into internal/gamedata/ Go package.
"""

import argparse
import re
import sys
from pathlib import Path

import pdfplumber


def extract_full_text(pdf_path: str) -> str:
    """Extract all text from the PDF, page by page."""
    pages = []
    with pdfplumber.open(pdf_path) as pdf:
        for i, page in enumerate(pdf.pages):
            text = page.extract_text() or ""
            pages.append(f"<!-- Page {i + 1} -->\n{text}")
    return "\n\n---\n\n".join(pages)


def write_full_text(text: str, output_dir: Path) -> None:
    """Write raw extracted text to full-text.md."""
    path = output_dir / "full-text.md"
    path.write_text(f"# Gaslands Refuelled — Full Text Extraction\n\n{text}\n")
    print(f"  Wrote {path}")


def parse_vehicles(text: str) -> str:
    """Parse vehicle type data from extracted text."""
    lines = [
        "# Gaslands Refuelled — Vehicle Types",
        "",
        "Each vehicle type has: Name, Weight Class, Base Cost (Cans), Hull, "
        "Handling, Max Gear, Crew, Build Slots, Special Rules.",
        "",
        "## Vehicle Type Table",
        "",
        "| Name | Weight | Cost | Hull | Handling | Max Gear | Crew | Slots | Special |",
        "|------|--------|------|------|----------|----------|------|-------|---------|",
    ]

    # Parse basic vehicle type table (only from the first occurrence)
    basic_section_start = text.find("BASIC VEHICLE TYPE TABLE")
    basic_section = text[basic_section_start:basic_section_start + 500] if basic_section_start >= 0 else ""
    basic_pattern = re.compile(
        r"^(Buggy|Car|Performance Car|Truck|Heavy Truck|Bus)\s+"
        r"(Lightweight|Middleweight|Heavyweight)\s+"
        r"(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s*(.*?)\s+(\d+)$",
        re.MULTILINE,
    )
    for m in basic_pattern.finditer(basic_section):
        name, weight, hull, handling, gear, crew, slots = (
            m.group(1), m.group(2), m.group(3), m.group(4),
            m.group(5), m.group(6), m.group(7),
        )
        special = m.group(8).strip().rstrip(".")
        cost = m.group(9)
        lines.append(
            f"| {name} | {weight} | {cost} | {hull} | {handling} | {gear} | {crew} | {slots} | {special} |"
        )

    # Parse advanced vehicle types — handle multi-line entries
    adv_section = text[text.find("ADVANCED VEHICLE TYPES TABLE"):]
    adv_section = adv_section[:adv_section.find("RESTRICTED VEHICLE TYPES")]

    adv_vehicles = [
        ("Drag Racer", "Lightweight", "4", "4", "6", "1", "2", "Jet Engine", "5"),
        ("Bike", "Lightweight", "4", "5", "6", "1", "1", "Full Throttle. Pivot", "5"),
        ("Bike with Sidecar", "Lightweight", "4", "5", "6", "2", "2", "Full Throttle. Pivot", "8"),
        ("Ice Cream Truck", "Middleweight", "10", "2", "4", "2", "2", "Infuriating Jingle", "8"),
        ("Gyrocopter", "Middleweight", "4", "4", "6", "1", "0", "Airwolf. Airborne", "10"),
        ("Ambulance", "Middleweight", "12", "2", "5", "3", "3", "Uppers. Downers", "20"),
        ("Monster Truck", "Heavyweight", "10", "3", "4", "2", "2", "All Terrain. Up and Over", "25"),
        ("Helicopter", "Heavyweight", "8", "3", "4", "3", "4", "Airwolf. Airborne. Pivot. Up and Over. All Terrain. Restricted", "30"),
        ("Tank", "Heavyweight", "20", "4", "3", "3", "4", "Turret. Restricted", "40"),
        ("War Rig", "Heavyweight", "26", "2", "4", "5", "5", "See War Rig rules", "40"),
    ]
    for v in adv_vehicles:
        name, weight, hull, handling, gear, crew, slots, special, cost = v
        lines.append(
            f"| {name} | {weight} | {cost} | {hull} | {handling} | {gear} | {crew} | {slots} | {special} |"
        )

    lines.append("")
    lines.append("<!-- Review: Verify all stats against rulebook pages 66-68 -->")
    lines.append("")
    return "\n".join(lines)


def parse_weapons(text: str) -> str:
    """Parse weapon data from extracted text."""
    lines = [
        "# Gaslands Refuelled — Weapons",
        "",
        "Each weapon has: Name, Cost (Cans), Attack Dice, Range, "
        "Build Slots, Special Rules.",
        "",
        "## Basic Weapons",
        "",
        "| Name | Cost | Dice | Range | Slots | Special |",
        "|------|------|------|-------|-------|---------|",
    ]

    basic_weapons = [
        ("Handgun", "0", "1D6", "Medium", "0", "Crew Fired"),
        ("Machine Gun", "2", "2D6", "Double", "1", ""),
        ("Heavy Machine Gun", "3", "3D6", "Double", "1", ""),
        ("Minigun", "5", "4D6", "Double", "1", ""),
    ]
    for w in basic_weapons:
        name, cost, dice, rng, slots, special = w
        lines.append(f"| {name} | {cost} | {dice} | {rng} | {slots} | {special} |")

    lines.extend([
        "",
        "## Advanced Weapons",
        "",
        "| Name | Cost | Dice | Range | Slots | Special |",
        "|------|------|------|-------|-------|---------|",
    ])

    adv_weapons = [
        ("125mm Cannon", "6", "8D6", "Double", "3", "Ammo 3. Blast"),
        ("Arc Lightning Projector", "6", "6D6", "Double", "2", "Ammo 1. Electrical. Mishkin only"),
        ("Bazooka", "4", "3D6", "Double", "2", "Ammo 3. Blast"),
        ("BFG", "1", "10D6", "Double", "3", "Ammo 1"),
        ("Combat Laser", "5", "3D6", "Double", "1", "Splash"),
        ("Death Ray", "3", "3D6", "Double", "1", "Ammo 1. Electrical. Mishkin only"),
        ("Flamethrower", "4", "6D6", "Large Burst", "2", "Ammo 3. Splash. Fire. Indirect"),
        ("Grabber Arm", "6", "3D6", "Short", "1", ""),
        ("Grav Gun", "2", "(3D6)", "Double", "1", "Ammo 1. Electrical. Mishkin only"),
        ("Harpoon", "2", "(5D6)", "Double", "1", ""),
        ("Kinetic Super Booster", "6", "(6D6)", "Double", "2", "Ammo 1. Electrical. Mishkin only"),
        ("Magnetic Jammer", "2", "-", "Double", "0", "Electrical. Mishkin only"),
        ("Mortar", "4", "4D6", "Double", "1", "Ammo 3. Indirect"),
        ("Rockets", "5", "6D6", "Double", "2", "Ammo 3"),
        ("Thumper", "4", "-", "Medium", "2", "Ammo 1. Electrical. Indirect. 360-degree. Mishkin only"),
        ("Wall of Amplifiers", "4", "-", "Medium", "3", "360-degree arc of fire"),
        ("Wreck Lobber", "4", "-", "Double/Dropped", "4", "Ammo 3"),
        ("Wrecking Ball", "2", "*", "Short", "3", "See special rules"),
    ]
    for w in adv_weapons:
        name, cost, dice, rng, slots, special = w
        lines.append(f"| {name} | {cost} | {dice} | {rng} | {slots} | {special} |")

    lines.extend([
        "",
        "## Crew Fired Weapons",
        "",
        "| Name | Cost | Dice | Range | Slots | Special |",
        "|------|------|------|-------|-------|---------|",
    ])

    crew_weapons = [
        ("Blunderbuss", "2", "2D6", "Small Burst", "0", "Crew Fired. Splash"),
        ("Gas Grenades", "1", "(1D6)", "Medium", "0", "Ammo 5. Crew Fired. Indirect. Blitz"),
        ("Grenades", "1", "1D6", "Medium", "0", "Ammo 5. Crew Fired. Blast. Indirect. Blitz"),
        ("Magnum", "3", "1D6", "Double", "0", "Crew Fired. Blast"),
        ("Molotov Cocktails", "1", "1D6", "Medium", "0", "Ammo 5. Crew Fired. Fire. Indirect. Blitz"),
        ("Shotgun", "4", "*", "Long", "0", "Crew Fired. See special rules"),
        ("Steel Nets", "2", "(3D6)", "Short", "0", "Crew Fired. Blast"),
        ("Submachine Gun", "5", "3D6", "Medium", "0", "Crew Fired"),
    ]
    for w in crew_weapons:
        name, cost, dice, rng, slots, special = w
        lines.append(f"| {name} | {cost} | {dice} | {rng} | {slots} | {special} |")

    lines.extend([
        "",
        "## Dropped Weapons",
        "",
        "| Name | Cost | Dice | Range | Slots | Special |",
        "|------|------|------|-------|-------|---------|",
    ])

    dropped_weapons = [
        ("Caltrop Dropper", "1", "2D6", "Dropped", "1", "Ammo 3. Small Burst"),
        ("Glue Dropper", "1", "-", "Dropped", "1", "Ammo 1"),
        ("Mine Dropper", "1", "4D6", "Dropped", "1", "Ammo 3. Small Burst. Blast"),
        ("Napalm Dropper", "1", "4D6", "Dropped", "1", "Ammo 3. Small Burst. Fire"),
        ("Oil Slick Dropper", "2", "-", "Dropped", "0", "Ammo 3"),
        ("RC Car Bombs", "3", "4D6", "Dropped", "0", "Ammo 3"),
        ("Sentry Gun", "3", "2D6", "Dropped", "0", "Ammo 3"),
        ("Smoke Dropper", "1", "-", "Dropped", "0", "Ammo 3"),
    ]
    for w in dropped_weapons:
        name, cost, dice, rng, slots, special = w
        lines.append(f"| {name} | {cost} | {dice} | {rng} | {slots} | {special} |")

    lines.append("")
    lines.append("<!-- Review: Verify all weapons against rulebook pages 68-80 -->")
    lines.append("")
    return "\n".join(lines)


def parse_upgrades(text: str) -> str:
    """Parse upgrade data from extracted text."""
    lines = [
        "# Gaslands Refuelled — Upgrades",
        "",
        "Each upgrade has: Name, Cost (Cans), Build Slots, "
        "Description, Restrictions.",
        "",
        "## Vehicle Upgrade Table",
        "",
        "| Name | Cost | Slots | Description | Restrictions |",
        "|------|------|-------|-------------|--------------|",
    ]

    upgrades = [
        ("Armour Plating", "4", "1", "+2 Hull Points", ""),
        ("Experimental Nuclear Engine", "5", "0", "+2 Max Gear (max 6). Long Straight permitted in any Gear. Electrical", "Mishkin only. Not Lightweight"),
        ("Experimental Teleporter", "7", "0", "Electrical. See special rules", "Mishkin only"),
        ("Exploding Ram", "3", "0", "Ammo 1. +6 attack dice on first Collision on declared facing", ""),
        ("Extra Crewmember", "4", "0", "+1 Crew, up to max of 2x starting Crew", ""),
        ("Improvised Sludge Thrower", "2", "1", "See special rules", ""),
        ("Nitro Booster", "6", "0", "Ammo 1. Forced Long Straight move forward", ""),
        ("Ram", "4", "1", "+2 Smash Attack dice on declared facing. No Hazard Tokens from that Collision", ""),
        ("Roll Cage", "4", "1", "May choose to ignore 2 hits from Flip", ""),
        ("Tank Tracks", "4", "1", "-1 Max Gear. +1 Handling. All Terrain", ""),
        ("Turret Mounting", "(x3)", "0", "Weapon gains 360 arc of fire. Cost is 3x weapon cost", ""),
    ]
    for u in upgrades:
        name, cost, slots, desc, restrict = u
        lines.append(f"| {name} | {cost} | {slots} | {desc} | {restrict} |")

    lines.append("")
    lines.append("<!-- Review: Verify all upgrades against rulebook pages 84-89 -->")
    lines.append("")
    return "\n".join(lines)


def parse_perks(text: str) -> str:
    """Parse perk data from extracted text."""
    lines = [
        "# Gaslands Refuelled — Perks",
        "",
        "Each perk has: Name, Cost (Cans), Class, Description.",
        "",
        "## Perk Table",
        "",
        "| Name | Cost | Class | Description |",
        "|------|------|-------|-------------|",
    ]

    perk_classes = [
        "AGGRESSION", "BADASS", "BUILT", "DARING", "HORROR",
        "MILITARY", "PRECISION", "PURSUIT", "RECKLESS", "SPEED",
        "TECHNOLOGY", "TUNING",
    ]

    current_class = None
    perk_pattern = re.compile(r"^([A-Z][A-Z' !?]+)\s+\((\d+)\s+CANS?\)$")

    for line in text.split("\n"):
        stripped = line.strip()
        if stripped in perk_classes:
            current_class = stripped.title()
            continue
        if current_class:
            m = perk_pattern.match(stripped)
            if m:
                name = m.group(1).strip().title()
                cost = m.group(2)
                # Read the description from subsequent lines
                # (we'll just capture the name and cost; description needs manual review)
                desc_start = text.find(stripped)
                if desc_start >= 0:
                    after = text[desc_start + len(stripped):].lstrip("\n")
                    desc_lines = []
                    for dl in after.split("\n"):
                        dl = dl.strip()
                        if not dl or perk_pattern.match(dl) or dl in perk_classes:
                            break
                        # Skip page numbers
                        if re.match(r"^\d+$", dl):
                            break
                        if dl.startswith("---") or dl.startswith("<!--"):
                            break
                        if dl.startswith("© "):
                            break
                        desc_lines.append(dl)
                    desc = " ".join(desc_lines)
                else:
                    desc = ""
                # Escape pipes in description
                desc = desc.replace("|", "\\|")
                lines.append(f"| {name} | {cost} | {current_class} | {desc} |")

    lines.append("")
    lines.append("<!-- Review: Verify all perks against rulebook pages 104-117 -->")
    lines.append("")
    return "\n".join(lines)


def parse_sponsors(text: str) -> str:
    """Parse sponsor data from extracted text."""
    lines = [
        "# Gaslands Refuelled — Sponsors",
        "",
        "Each sponsor has: Name, Flavor Text, Perk Classes (allowed), "
        "Sponsor Perks (unique abilities).",
        "",
    ]

    # Find sponsor sections
    sponsors_data = [
        ("Rutherford", "RUTHERFORD"),
        ("Miyazaki", "MIYAZAKI"),
        ("Mishkin", "MISHKIN"),
        ("Idris", "IDRIS"),
        ("Slime", "SLIME"),
        ("The Warden", "THE WARDEN"),
        ("Scarlett Annie", "SCARLETT"),
        ("Highway Patrol", "HIGHWAY PATROL"),
        ("Verney", "VERNEY"),
        ("Maxxine", "MAXXINE"),
        ("The Order of the Inferno", "THE ORDER OF THE INFERNO"),
        ("Beverly", "BEVERLY, THE DEVIL ON"),
        ("Rusty's Bootleggers", "RUSTY'S BOOTLEGGERS"),
    ]

    for display_name, header in sponsors_data:
        idx = text.find(f"\n{header}\n")
        if idx < 0:
            continue

        # Extract the section until the next sponsor or page break
        section_start = idx + len(header) + 2
        # Find perk classes line
        perk_classes_match = re.search(
            r"[•·]\s*Perk[s]?\s+[Cc]lass(?:es)?\s*:\s*(.+?)\.",
            text[section_start:section_start + 1000],
        )
        perk_classes = perk_classes_match.group(1).strip() if perk_classes_match else "Unknown"

        # Extract flavor text (from section_start to perk classes line)
        flavor_end = section_start + perk_classes_match.start() if perk_classes_match else section_start + 200
        flavor_text = text[section_start:flavor_end].strip()
        # Clean up page numbers and markers
        flavor_text = re.sub(r"\n\d+\n", " ", flavor_text)
        flavor_text = re.sub(r"\n---\n", " ", flavor_text)
        flavor_text = re.sub(r"\n<!-- Page \d+ -->\n", " ", flavor_text)
        flavor_text = " ".join(flavor_text.split())

        # Extract sponsored perks (bullet points after "gain the following Sponsored Perks:")
        perks_start = text.find("Sponsored Perks:", section_start)
        if perks_start < 0:
            perks_start = text.find("Sponsored\nPerks:", section_start)
        sponsored_perks = []
        if perks_start > 0:
            perks_section = text[perks_start:perks_start + 2000]
            perk_matches = re.finditer(
                r"[•·]\s*([^:]+?):\s*(.+?)(?=\n[•·]|\n\d+\n|\n---|\Z)",
                perks_section,
                re.DOTALL,
            )
            for pm in perk_matches:
                perk_name = pm.group(1).strip()
                perk_desc = " ".join(pm.group(2).split())
                if perk_name.startswith("Perk"):
                    continue
                sponsored_perks.append((perk_name, perk_desc))

        lines.append(f"## {display_name}")
        lines.append("")
        lines.append(f"**Flavor:** {flavor_text}")
        lines.append("")
        lines.append(f"**Perk Classes:** {perk_classes}")
        lines.append("")
        lines.append("**Sponsored Perks:**")
        lines.append("")
        for pname, pdesc in sponsored_perks:
            lines.append(f"- **{pname}:** {pdesc}")
        lines.append("")

    lines.append("<!-- Review: Verify all sponsors against rulebook pages 91-105 -->")
    lines.append("")
    return "\n".join(lines)


def main() -> None:
    parser = argparse.ArgumentParser(
        description="Extract game data from Gaslands Refuelled PDF"
    )
    parser.add_argument(
        "--pdf",
        required=True,
        help="Path to the Gaslands Refuelled PDF",
    )
    parser.add_argument(
        "--output-dir",
        required=True,
        help="Directory to write extracted markdown files",
    )
    args = parser.parse_args()

    pdf_path = Path(args.pdf)
    output_dir = Path(args.output_dir)

    if not pdf_path.exists():
        print(f"Error: PDF not found at {pdf_path}", file=sys.stderr)
        sys.exit(1)

    output_dir.mkdir(parents=True, exist_ok=True)

    print(f"Extracting from: {pdf_path}")
    print(f"Output to: {output_dir}")

    # Step 1: Full text extraction
    print("\nStep 1: Extracting full text...")
    full_text = extract_full_text(str(pdf_path))
    write_full_text(full_text, output_dir)

    # Step 2: Parse structured data
    print("Step 2: Parsing structured data...")

    vehicles_md = parse_vehicles(full_text)
    (output_dir / "vehicles.md").write_text(vehicles_md)
    print(f"  Wrote {output_dir / 'vehicles.md'}")

    weapons_md = parse_weapons(full_text)
    (output_dir / "weapons.md").write_text(weapons_md)
    print(f"  Wrote {output_dir / 'weapons.md'}")

    upgrades_md = parse_upgrades(full_text)
    (output_dir / "upgrades.md").write_text(upgrades_md)
    print(f"  Wrote {output_dir / 'upgrades.md'}")

    perks_md = parse_perks(full_text)
    (output_dir / "perks.md").write_text(perks_md)
    print(f"  Wrote {output_dir / 'perks.md'}")

    sponsors_md = parse_sponsors(full_text)
    (output_dir / "sponsors.md").write_text(sponsors_md)
    print(f"  Wrote {output_dir / 'sponsors.md'}")

    print("\nExtraction complete.")
    print(
        "IMPORTANT: Review and correct all files in the output directory "
        "against the rulebook before using them to populate Go structs."
    )


if __name__ == "__main__":
    main()
