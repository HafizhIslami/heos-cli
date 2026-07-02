# EXECUTION_ORCHESTRATOR.md

Version: 1.0.0

Purpose

Coordinate all AI Employees during task execution.

The orchestrator decides:

- who works
- when they work
- when they stop
- when work is handed over
- when user approval is required
- when the task is complete

No AI may skip the orchestrator.

---

## Execution Pipeline

User Request

↓

Request Handler (Qwen3)

↓

Requirement Analysis (Qwen3)

↓

Product Planning (Qwen3)

↓

Architecture Design (Qwen3)

↓

Implementation Planning (Qwen3)

↓

============================

HANDOFF TO QWEN2.5

============================

↓

Implementation

↓

Self Validation

↓

Unit Test

↓

Implementation Report

↓

============================

HANDOFF TO QWEN3

============================

↓

Independent Review

↓

Approved?

├── NO

│

▼

Revision Request

↓

HANDOFF TO QWEN2.5

↓

Fix

↓

Self Review

↓

Return to Reviewer

└── YES

↓

Acceptance

↓

Done

---

## Role Ownership

Request Analysis

Owner

Qwen3

Product Planning

Owner

Qwen3

Architecture

Owner

Qwen3

Implementation

Owner

Qwen2.5

Testing

Owner

Qwen2.5

Review

Owner

Qwen3

Acceptance

Owner

Qwen3

---

## Automatic Continuation

Every AI continues automatically.

Never stop after finishing a phase.

Immediately transfer work to the next owner.

Stop only when:

- User decision required
- High-risk action detected
- Task completed

---

## Approval Gates

Only stop for:

Business ambiguity

Breaking API

Destructive migration

Production secrets

Legal concerns

Multiple valid business choices

Everything else continues automatically.

---

## Final Rule

The orchestrator owns the task.

Individual AI owns only the current phase.