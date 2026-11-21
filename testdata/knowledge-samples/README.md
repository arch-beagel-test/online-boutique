# Knowledge-base samples: manifest

Thirty files in `docs/knowledge-samples/` for measuring ArchBeagle's knowledge base and MarkItDown
(architecture/0104, issue ArchBeagle/archbeagle#1666). This manifest sits under `testdata/`, which the docs read
skips, so its answers are never indexed.

`manifest.json` lists every file with its format, origin, licence, what the read should do with it, and
questions a search should answer from it (55 in all).

| Format | Files |
|---|---|
| adoc | 1 |
| docx | 5 |
| html | 1 |
| md | 5 |
| pdf (born-digital) | 3 |
| pdf (born-digital, long) | 2 |
| pdf (corrupt) | 1 |
| pdf (scanned) | 3 |
| pptx | 4 |
| rst | 2 |
| txt | 2 |
| xlsx | 1 |

- **Authored for this set** (Apache-2.0, as this repository): ADRs, specs, runbooks, notes and slides about this
  repository's real services. Decisions, numbers and names in them are invented for testing.
- **Scanned:** authored pages rendered to images and wrapped in a PDF with no text layer (one rotated, one low
  resolution). MarkItDown gives no text for them without OCR; the expected answer is an empty doc that says so.
- **External:** NIST SP 800-204 and SP 800-207, works of the US government (public domain), from nvlpubs.nist.gov.
- **Edge cases:** `data/inventory.xlsx` (a format 0104 leaves out: must be skipped and reported) and
  `broken/truncated.pdf` (a valid PDF cut to a third: must fail on its own without stopping the run).

Checked with MarkItDown 0.1.9 (`pdf`, `docx`, `pptx` extras) on 2026-11-21: every authored file converts; the three
scans give 0 characters; the NIST PDFs give about 127,000 and 158,000 characters; the Excel file and the broken PDF
raise a conversion error.
