# STATE_MACHINE.md

# Hynexis Engineering Operating System (HEOS)

Version: 1.0.0

---

# Purpose

This document defines the lifecycle of every task executed by AI Employees.

Every task must exist in exactly one state.

Every state has:

- Owner
- Objective
- Allowed Actions
- Exit Criteria
- Next State

No AI may skip, invent, or reorder states.

---

# State Diagram

NEW
↓

ANALYZING
↓

PLANNING
↓

ARCHITECTING
↓

IMPLEMENTATION_READY
↓

IMPLEMENTING
↓

SELF_REVIEW
↓

REVIEWING
↓

REVISION_REQUIRED (optional)
↓

IMPLEMENTING
↓

SELF_REVIEW
↓

REVIEWING
↓

APPROVED
↓

TESTING
↓

DOCUMENTING
↓

ACCEPTANCE
↓

DONE

---

# State Definitions

## NEW

Owner

System

Purpose

Task has been received but not processed.

Allowed Actions

- Register task
- Assign task ID

Exit Criteria

Task registered.

Next State

ANALYZING

---

## ANALYZING

Owner

Request Handler

Purpose

Understand user intent.

Deliverables

- Requirement Analysis
- Clarified objectives
- Questions (if any)

Exit Criteria

Requirements understood.

Next State

PLANNANNING

---

## PLANNING

Owner

Product Planner

Purpose

Transform requirements into executable work.

Deliverables

- Product Plan
- Scope
- Milestones

Exit Criteria

Plan approved internally.

Next State

ARCHITECTING

---

## ARCHITECTING

Owner

Software Architect

Purpose

Design the implementation.

Deliverables

- Architecture Proposal
- API Plan
- Database Plan
- Folder Structure
- Technical Decisions

Exit Criteria

Architecture complete.

Next State

IMPLEMENTATION_READY

---

## IMPLEMENTATION_READY

Owner

Orchestrator

Purpose

Prepare implementation package.

Deliverables

- Work Order
- Handoff Package

Exit Criteria

Engineer receives task.

Next State

IMPLEMENTING

---

## IMPLEMENTING

Owner

Software Engineer

Purpose

Implement approved design.

Allowed Actions

- Create code
- Modify code
- Refactor
- Create migrations
- Create tests

Not Allowed

- Change architecture
- Change business requirements

Exit Criteria

Implementation complete.

Next State

SELF_REVIEW

---

## SELF_REVIEW

Owner

Software Engineer

Purpose

Review own implementation.

Checklist

□ Build succeeds

□ Lint passes

□ Tests pass

□ No obvious bug

□ Documentation updated

Exit Criteria

Self-review complete.

Next State

REVIEWING

---

## REVIEWING

Owner

Senior Reviewer

Purpose

Independent verification.

Possible Results

APPROVED

REVISION_REQUIRED

---

## REVISION_REQUIRED

Owner

Software Engineer

Purpose

Fix reviewer findings.

Allowed Actions

Only address review findings.

Do not introduce unrelated changes.

Exit Criteria

All findings resolved.

Next State

SELF_REVIEW

---

## APPROVED

Owner

Reviewer

Purpose

Code quality accepted.

Exit Criteria

Ready for testing.

Next State

TESTING

---

## TESTING

Owner

QA Engineer

Purpose

Validate implementation.

Testing Types

- Unit
- Integration
- Regression
- Boundary
- Negative
- Security (when applicable)

Exit Criteria

All mandatory tests passed.

Next State

DOCUMENTING

---

## DOCUMENTING

Owner

Engineer

Purpose

Update documentation.

Required

README

API

Migration

Release Notes

Architecture (if changed)

Exit Criteria

Documentation complete.

Next State

ACCEPTANCE

---

## ACCEPTANCE

Owner

Architect

Purpose

Final engineering acceptance.

Verify

Business objective achieved.

Architecture preserved.

Security maintained.

Documentation complete.

Tests passed.

Possible Results

DONE

REVISION_REQUIRED

---

## DONE

Owner

System

Purpose

Task completed.

No further action required.

---

# State Transition Rules

Allowed

NEW → ANALYZING

ANALYZING → PLANNING

PLANNING → ARCHITECTING

ARCHITECTING → IMPLEMENTATION_READY

IMPLEMENTATION_READY → IMPLEMENTING

IMPLEMENTING → SELF_REVIEW

SELF_REVIEW → REVIEWING

REVIEWING → APPROVED

REVIEWING → REVISION_REQUIRED

REVISION_REQUIRED → IMPLEMENTING

APPROVED → TESTING

TESTING → DOCUMENTING

DOCUMENTING → ACCEPTANCE

ACCEPTANCE → DONE

ACCEPTANCE → REVISION_REQUIRED

---

# Forbidden Transitions

Never allow:

IMPLEMENTING → DONE

IMPLEMENTING → TESTING

ANALYZING → IMPLEMENTING

PLANNING → TESTING

ARCHITECTING → DONE

SELF_REVIEW → DONE

REVIEWING → IMPLEMENTING

DOCUMENTING → IMPLEMENTING

---

# Automatic Continuation

Every AI must automatically continue to the next state.

Never wait for user input unless:

- Business ambiguity exists.
- User approval is explicitly required.
- High-risk action is detected.

Otherwise continue automatically.

---

# Failure Handling

If execution fails:

Remain in current state.

Record:

- Failure reason
- Evidence
- Suggested resolution

Retry only after issue is resolved.

Never silently ignore failures.

---

# Retry Policy

Implementation

Unlimited until approved.

Testing

Unlimited until passed.

Review

Unlimited until approved.

Architecture

Unlimited until accepted.

Every retry must document changes made.

---

# Completion Criteria

A task reaches DONE only if:

□ Requirements satisfied.

□ Architecture followed.

□ Security standards met.

□ Review approved.

□ Tests passed.

□ Documentation updated.

□ Acceptance completed.

Only then may the state become DONE.

---

# Final Principle

The state machine defines the only valid lifecycle of work.

Every AI Employee must obey it without exception.
