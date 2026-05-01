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


def extract_tables(pdf_path: str) -> list[dict]:
    """Extract all tables from the PDF with page numbers."""
    tables = []
    with pdfplumber.open(pdf_path) as pdf:
        for i, page in enumerate(pdf.pages):
            page_tables = page.extract_tables() or []
            for table in page_tables:
                tables.append({"page": i + 1, "data": table})
    return tables


def write_full_text(text: str, output_dir: Path) -> None:
    """Write raw extracted text to full-text.md."""
    path = output_dir / "full-text.md"
    path.write_text(f"# Gaslands Refuelled — Full Text Extraction\n\n{text}\n")
    print(f"  Wrote {path}")


def parse_vehicles(text: str, tables: list[dict]) -> str:
    """Parse vehicle type data from extracted text and tables.

    TODO: Implement structured parsing. For now, outputs a template
    that needs manual population from the full-text extraction.
    """
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
        "<!-- TODO: Populate from extracted data or manually from rulebook -->",
        "",
    ]
    return "\n".join(lines)


def parse_weapons(text: str, tables: list[dict]) -> str:
    """Parse weapon data from extracted text and tables."""
    lines = [
        "# Gaslands Refuelled — Weapons",
        "",
        "Each weapon has: Name, Cost (Cans), Attack Dice, Range, "
        "Build Slots, Special Rules.",
        "",
        "## Weapon Table",
        "",
        "| Name | Cost | Dice | Range | Slots | Special |",
        "|------|------|------|-------|-------|---------|",
        "<!-- TODO: Populate from extracted data or manually from rulebook -->",
        "",
    ]
    return "\n".join(lines)


def parse_upgrades(text: str, tables: list[dict]) -> str:
    """Parse upgrade data from extracted text and tables."""
    lines = [
        "# Gaslands Refuelled — Upgrades",
        "",
        "Each upgrade has: Name, Cost (Cans), Build Slots, "
        "Description, Restrictions.",
        "",
        "## Upgrade Table",
        "",
        "| Name | Cost | Slots | Description | Restrictions |",
        "|------|------|-------|-------------|--------------|",
        "<!-- TODO: Populate from extracted data or manually from rulebook -->",
        "",
    ]
    return "\n".join(lines)


def parse_perks(text: str, tables: list[dict]) -> str:
    """Parse perk data from extracted text and tables."""
    lines = [
        "# Gaslands Refuelled — Perks",
        "",
        "Each perk has: Name, Cost (Cans), Class, Description.",
        "",
        "Perk classes: Aggression, Badass, Speed, Precision, "
        "Military, Technology, etc.",
        "",
        "## Perk Table",
        "",
        "| Name | Cost | Class | Description |",
        "|------|------|-------|-------------|",
        "<!-- TODO: Populate from extracted data or manually from rulebook -->",
        "",
    ]
    return "\n".join(lines)


def parse_sponsors(text: str, tables: list[dict]) -> str:
    """Parse sponsor data from extracted text and tables."""
    lines = [
        "# Gaslands Refuelled — Sponsors",
        "",
        "Each sponsor has: Name, Flavor Text, Perk Classes (allowed), "
        "Sponsor Perks (unique abilities).",
        "",
        "## Sponsors",
        "",
        "<!-- TODO: Populate each sponsor section from extracted data -->",
        "",
    ]
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

    # Step 2: Table extraction
    print("Step 2: Extracting tables...")
    tables = extract_tables(str(pdf_path))
    print(f"  Found {len(tables)} tables")

    # Step 3: Parse structured data
    print("Step 3: Parsing structured data...")

    vehicles_md = parse_vehicles(full_text, tables)
    (output_dir / "vehicles.md").write_text(vehicles_md)
    print(f"  Wrote {output_dir / 'vehicles.md'}")

    weapons_md = parse_weapons(full_text, tables)
    (output_dir / "weapons.md").write_text(weapons_md)
    print(f"  Wrote {output_dir / 'weapons.md'}")

    upgrades_md = parse_upgrades(full_text, tables)
    (output_dir / "upgrades.md").write_text(upgrades_md)
    print(f"  Wrote {output_dir / 'upgrades.md'}")

    perks_md = parse_perks(full_text, tables)
    (output_dir / "perks.md").write_text(perks_md)
    print(f"  Wrote {output_dir / 'perks.md'}")

    sponsors_md = parse_sponsors(full_text, tables)
    (output_dir / "sponsors.md").write_text(sponsors_md)
    print(f"  Wrote {output_dir / 'sponsors.md'}")

    print("\nExtraction complete.")
    print(
        "IMPORTANT: Review and correct all files in the output directory "
        "against the rulebook before using them to populate Go structs."
    )


if __name__ == "__main__":
    main()
