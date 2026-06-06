#!/usr/bin/env python3
"""Tests for Sprint 9 CORE-RC-DOC-CONSISTENCY-A validation."""

from __future__ import annotations

import tempfile
import unittest
from pathlib import Path

import validate_sprint9_core_doc_consistency as doc_consistency


class Sprint9CoreDocConsistencyValidationTest(unittest.TestCase):
    def test_default_docs_and_makefile_are_sprint9_aligned(self) -> None:
        doc_consistency.validate_workspace(doc_consistency.ROOT)

    def test_validation_accepts_historical_sprint9_markers_after_later_sprints(self) -> None:
        with tempfile.TemporaryDirectory() as tmpdir:
            root = Path(tmpdir)
            (root / "repo" / "development-records").mkdir(parents=True)
            (root / "ANI-DOCS-INDEX.md").write_text(
                "Sprint 6-10 完成 contract/local/release-prep scaffold\n", encoding="utf-8"
            )
            (root / "ANI-06-开发计划.md").write_text("Sprint 9 ⭐ | ✅ Core-only 已完成\n", encoding="utf-8")
            (root / "repo" / "CURRENT-SPRINT.md").write_text(
                "Sprint 9 RC readiness gates\n", encoding="utf-8"
            )
            (root / "repo" / "Makefile").write_text(
                "\n".join(f"{target}:" for target in doc_consistency.REQUIRED_MAKE_TARGETS),
                encoding="utf-8",
            )
            (root / "repo" / "development-records" / "README.md").write_text(
                "\n".join(doc_consistency.REQUIRED_RECORDS), encoding="utf-8"
            )

            doc_consistency.validate_workspace(root)

    def test_validation_rejects_stale_sprint8_current_marker(self) -> None:
        with tempfile.TemporaryDirectory() as tmpdir:
            root = Path(tmpdir)
            (root / "repo" / "development-records").mkdir(parents=True)
            (root / "ANI-DOCS-INDEX.md").write_text("Sprint 9 Core-only 代码开发已完成\n", encoding="utf-8")
            (root / "ANI-06-开发计划.md").write_text("Sprint 9 Core-only 代码开发已完成\n", encoding="utf-8")
            (root / "repo" / "CURRENT-SPRINT.md").write_text(
                "当前重心：Sprint 8 Core-only 代码开发已完成\n", encoding="utf-8"
            )
            (root / "repo" / "Makefile").write_text("validate-sprint9-rc:\n", encoding="utf-8")
            (root / "repo" / "development-records" / "README.md").write_text(
                "Sprint 9 Core-only 代码开发已完成\nSPRINT9-CLOSURE-A\n", encoding="utf-8"
            )

            with self.assertRaises(SystemExit) as raised:
                doc_consistency.validate_workspace(root)

        self.assertIn("stale Sprint 8 current marker", str(raised.exception))


if __name__ == "__main__":
    unittest.main()
