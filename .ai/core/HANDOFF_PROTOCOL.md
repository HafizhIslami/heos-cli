# HANDOFF_PROTOCOL.md

# Hynexis Engineering Operating System (HEOS)

Version: 1.0.0

---

# Purpose

This document defines the mandatory protocol for transferring work between AI Employees.

A handoff must provide complete context so the receiving AI can continue the task immediately without requesting previously known information.

Every completed phase must end with a handoff package.

No work may be transferred without a valid handoff.

---

# Principles

A handoff must be:

- Complete
- Accurate
- Actionable
- Verifiable
- Self-contained

The receiving AI should never need to infer missing information.

---

# Handoff Lifecycle

Current Owner

↓

Complete Assigned Phase

↓

Generate Handoff Package

↓

Transfer Ownership

↓

Receiving AI Validates Package

↓

Continue Execution

---

# Mandatory Handoff Package

Every handoff must include the following sections.

## 1. Task Information

Task ID

Task Name

Current State

Current Owner

Next Owner

Priority

Status

---

## 2. Business Context

Business Objective

Expected Business Outcome

Success Criteria

Business Constraints

---

## 3. Current Progress

Completed Work

Remaining Work

Blocked Work

Completion Percentage

---

## 4. Deliverables

Documents Created

Source Files Modified

Database Changes

API Changes

Configuration Changes

Infrastructure Changes

---

## 5. Decisions

List every engineering decision.

Each decision must contain:

Decision

Reason

Alternative Considered

Impact

---

## 6. Risks

Known Risks

Potential Risks

Security Risks

Migration Risks

Performance Risks

---

## 7. Assumptions

Explicitly list every assumption.

Never hide assumptions.

---

## 8. Open Questions

Question

Reason

Blocking?

Priority

---

## 9. Validation

Build Status

Lint Status

Tests Status

Migration Status

Documentation Status

---

## 10. Next Action

Exactly one next action.

Example:

Implement Membership Repository.

Not:

Continue development.

---

# Ownership Rules

The current owner retains responsibility until:

- Handoff Package completed.
- Receiving AI validates package.

Only then ownership transfers.

---

# Receiving AI Validation

The receiving AI must verify:

□ Task understood

□ Business goal understood

□ Architecture understood

□ Risks understood

□ Files identified

□ Next action clear

If validation fails:

Reject handoff.

Request missing information.

Never continue with incomplete context.

---

# Handoff Types

## Planning → Architecture

Includes:

Requirements

Business Goals

Scope

Constraints

Risks

---

## Architecture → Engineer

Includes:

Architecture Proposal

Folder Structure

API Design

Database Design

Implementation Plan

Definition of Done

---

## Engineer → Reviewer

Includes:

Implementation Summary

Files Changed

Tests Added

Known Limitations

Self Review Result

---

## Reviewer → Engineer

Includes:

Review Findings

Severity

Reason

Recommendation

Expected Fix

---

## Engineer → Tester

Includes:

Features Implemented

Test Scope

Known Edge Cases

Expected Results

---

## Tester → Acceptance

Includes:

Test Results

Coverage Summary

Outstanding Issues

Recommendation

---

## Acceptance → Done

Includes:

Acceptance Decision

Release Readiness

Final Summary

Lessons Learned

---

# Rejected Handoff

A receiving AI must reject the handoff if:

Business objective unclear.

Architecture missing.

Implementation incomplete.

Tests missing.

Critical information missing.

Never continue from an incomplete handoff.

---

# Communication Style

Every handoff must be:

Objective

Concise

Structured

Evidence-based

Avoid:

Opinions

Speculation

Redundant explanations

---

# Standard Handoff Template

Task ID:

Task Name:

Current State:

Current Owner:

Next Owner:

---

Business Objective

...

---

Completed Work

...

---

Remaining Work

...

---

Files Modified

...

---

Decisions

...

---

Risks

...

---

Validation

Build:

Lint:

Tests:

Documentation:

---

Next Action

...

---

# Completion Rule

A phase is considered complete only when:

□ Deliverables finished.

□ Validation completed.

□ Handoff package generated.

□ Ownership transferred.

Without a valid handoff, the phase is NOT complete.

---

# Final Principle

A task never belongs to an AI forever.

It belongs only until a successful handoff is completed.

Clear handoffs create continuous execution.

Poor handoffs create project failure.