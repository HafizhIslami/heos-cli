# WORKFLOW.md

# Hynexis Engineering Operating System (HEOS)

Version: 1.0.0

---

# Purpose

This document defines the mandatory execution workflow for every software engineering task.

Every AI must follow this workflow.

No phase may be skipped unless explicitly instructed by the user.

---

# Engineering Lifecycle

Every task follows this lifecycle.

Request

↓

Analysis

↓

Planning

↓

Architecture

↓

Implementation

↓

Self Review

↓

Independent Review

↓

Revision

↓

Testing

↓

Documentation

↓

Acceptance

↓

Completed

---

# Phase 1 — Request Analysis

Assigned Model:

Qwen3

Objectives:

* Understand the business goal.
* Identify expected outcome.
* Determine scope.
* Detect ambiguity.
* Detect missing information.

Deliverables:

* Requirement Summary
* Scope
* Non-Scope
* Assumptions
* Questions (if required)

If critical information is missing:

Stop.

Request clarification.

Never guess.

---

# Phase 2 — Product Planning

Assigned Model:

Qwen3

Objectives:

Break the work into logical milestones.

Deliverables:

* Milestone List
* Task Breakdown
* Dependencies
* Risks
* Estimated Complexity

---

# Phase 3 — Architecture

Assigned Model:

Qwen3

Objectives:

Design implementation.

Review:

* Existing architecture
* Database
* API
* Security
* Performance
* Scalability

Deliverables:

Architecture Proposal

Only redesign existing architecture when there is a strong technical reason.

---

# Phase 4 — Implementation

Assigned Model:

Qwen2.5-Coder

Objectives:

Implement the approved plan.

Implementation must:

* Follow project conventions.
* Preserve compatibility.
* Minimize changes.
* Avoid duplicated logic.
* Keep code readable.

Deliverables:

Production-ready implementation.

---

# Phase 5 — Self Review

Assigned Model:

Qwen2.5-Coder

Before requesting review, verify:

* Build
* Formatting
* Lint
* Naming
* Error handling
* Logging
* Validation
* Authorization
* Transactions

If any issue exists:

Fix it before continuing.

---

# Phase 6 — Independent Review

Assigned Model:

Qwen3

Objectives:

Review implementation independently.

Search for:

* Incorrect logic
* Security vulnerabilities
* Architecture violations
* Maintainability issues
* Performance issues
* Missing edge cases

Possible decisions:

APPROVED

REVISION REQUIRED

REJECTED

---

# Phase 7 — Revision

Assigned Model:

Qwen2.5-Coder

Objectives:

Resolve reviewer findings.

Do not introduce unrelated changes.

Repeat review until approved.

---

# Phase 8 — Testing

Assigned Model:

Qwen2.5-Coder

Generate or update tests covering:

* Happy Path
* Negative Path
* Boundary Cases
* Edge Cases
* Regression

Fix failing tests before continuing.

---

# Phase 9 — Documentation

Assigned Model:

Qwen3

Update documentation when required.

Examples:

* README
* API Documentation
* ADR
* Architecture Diagram
* Migration Notes
* Changelog

Documentation should explain why, not only what.

---

# Phase 10 — Acceptance

Assigned Model:

Qwen3

Verify:

✓ Business objective achieved

✓ Architecture respected

✓ Security maintained

✓ Tests passed

✓ Documentation updated

✓ No known critical issues

If all conditions are satisfied:

TASK COMPLETED

Otherwise:

Return to the appropriate phase.

---

# Automatic Iteration

The AI team must continue working without waiting for additional instructions when:

* The next step is obvious.
* The task has already been approved.
* Only engineering work remains.
* No business decision is required.

Continue until:

* The task is completed.
* A blocker is encountered.
* User intervention is required.

---

# Approval Required

Execution must pause only when one or more of the following occurs:

* Business requirements change.
* Multiple valid architectural options exist.
* Destructive database migration.
* Data loss is possible.
* Public API breaking changes.
* Security-sensitive changes.
* Missing business information.

Present available options with advantages, disadvantages, risks, and recommendations.

Wait for user decision.

---

# Failure Handling

If an implementation fails:

1. Identify root cause.
2. Explain the failure.
3. Propose a solution.
4. Retry implementation.
5. Continue workflow.

Never abandon a task without explanation.

---

# Completion Policy

A task is not complete because code exists.

A task is complete only when the requested business outcome has been achieved and verified.

Always optimize for business value rather than code volume.
