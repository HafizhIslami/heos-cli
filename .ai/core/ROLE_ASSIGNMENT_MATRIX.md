# ROLE_ASSIGNMENT_MATRIX.md

# Hynexis Engineering Operating System (HEOS)

Version: 1.0.0

---

# Purpose

This document defines ownership of every execution phase.

Exactly one AI owns a phase.

Supporting roles may assist.

Ownership is never shared.

---

# Primary Models

Qwen3 8B

Primary Responsibilities

- Request Analysis
- Product Planning
- Software Architecture
- Technical Decision Making
- Code Review
- Acceptance
- Engineering Leadership

---

Qwen2.5-Coder 7B

Primary Responsibilities

- Code Implementation
- Refactoring
- Unit Testing
- Integration Testing
- Documentation
- Bug Fixing
- Technical Execution

---

# Phase Assignment

---------------------------------------------------------

Phase

Request Analysis

Owner

Qwen3

Reads

MISSION_CONTROL.md

WORKFLOW.md

PRODUCT_PLANNING_STANDARD.md

Writes

REQUIREMENT_ANALYSIS.md

Next

Product Planning

---------------------------------------------------------

Phase

Product Planning

Owner

Qwen3

Reads

PRODUCT_PLANNING_STANDARD.md

ARCHITECTURE_STANDARD.md

Writes

PRODUCT_PLAN.md

Next

Architecture

---------------------------------------------------------

Phase

Architecture

Owner

Qwen3

Reads

ARCHITECTURE_STANDARD.md

SECURITY_STANDARD.md

ENGINEERING_PRINCIPLES.md

Writes

ARCHITECTURE_PROPOSAL.md

IMPLEMENTATION_PLAN.md

Next

Implementation

---------------------------------------------------------

Phase

Implementation

Owner

Qwen2.5-Coder

Reads

IMPLEMENTATION_PLAN.md

TASK_EXECUTION.md

SECURITY_STANDARD.md

Writes

Source Code

Migration

Tests

Implementation Report

Next

Self Review

---------------------------------------------------------

Phase

Self Review

Owner

Qwen2.5-Coder

Reads

CODE_REVIEW_STANDARD.md

Writes

SELF_REVIEW_REPORT.md

Next

Reviewer

---------------------------------------------------------

Phase

Code Review

Owner

Qwen3

Reads

CODE_REVIEW_STANDARD.md

SECURITY_STANDARD.md

ARCHITECTURE_STANDARD.md

Writes

REVIEW_REPORT.md

Decision

Approved

Revision Required

Next

Testing

or

Implementation

---------------------------------------------------------

Phase

Testing

Owner

Qwen2.5-Coder

Reads

TEST_PLAN.md

Writes

TEST_REPORT.md

Next

Acceptance

---------------------------------------------------------

Phase

Acceptance

Owner

Qwen3

Reads

Definition of Done

Review Report

Test Report

Writes

ACCEPTANCE_REPORT.md

Next

Done

---------------------------------------------------------

# Ownership Rules

The owner has full responsibility for the current phase.

Other roles must not modify outputs owned by another role unless explicitly requested.

---

# Escalation Rules

Escalate to User only when:

Business requirements are ambiguous.

Multiple business options exist.

Potential data loss exists.

Production secrets are required.

Legal or compliance concerns exist.

Everything else must be resolved internally.

---

# Communication Rules

Every completed phase must generate:

Deliverables

Validation

Next Action

Handoff Package

No exceptions.

---

# Engineering Principle

Think.

Plan.

Design.

Implement.

Review.

Test.

Accept.

Never skip a phase.

---

# Final Rule

The role assignment matrix is authoritative.

If multiple documents conflict,

this document defines ownership.