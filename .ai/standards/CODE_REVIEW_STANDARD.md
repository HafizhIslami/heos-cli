# CODE_REVIEW_STANDARD.md

# Hynexis Engineering Operating System (HEOS)

Version: 1.0.0

---

# Purpose

This document defines the mandatory code review process.

Every implementation must pass independent review before it is considered complete.

Review exists to improve software quality.

Not to criticize developers.

---

# Review Principles

A reviewer evaluates:

- correctness
- maintainability
- security
- scalability
- readability
- consistency
- business alignment

Never review personal coding style.

Review engineering quality.

---

# Review Order

Always review in this sequence.

1.

Business Objective

↓

2.

Architecture

↓

3.

Security

↓

4.

Correctness

↓

5.

Performance

↓

6.

Maintainability

↓

7.

Readability

↓

8.

Testing

↓

9.

Documentation

Never skip the order.

---

# Business Review

Questions

Does this implementation solve the requested problem?

Does it introduce unnecessary functionality?

Is the solution proportional to the problem?

---

# Architecture Review

Verify:

□ Layer boundaries respected

□ No circular dependency

□ No business logic inside controllers

□ Proper dependency direction

□ Existing architecture preserved

---

# Security Review

Verify:

□ Authentication

□ Authorization

□ Input validation

□ Output sanitization

□ Secret management

□ Sensitive logging

□ SQL Injection

□ XSS

□ CSRF

□ SSRF

□ Path Traversal

□ IDOR

---

# Correctness Review

Verify:

□ Logic correct

□ Edge cases

□ Error handling

□ Null safety

□ Transactions

□ Concurrency

□ Resource cleanup

---

# Maintainability Review

Review:

Naming

Complexity

Function length

Class responsibility

Duplication

Dead code

Comments

Folder organization

Configuration

Dependencies

---

# Readability Review

Code should explain itself.

Avoid:

Nested logic

Magic values

Hidden behavior

Long functions

Ambiguous names

---

# Performance Review

Only review actual bottlenecks.

Avoid premature optimization.

Review:

Database queries

Memory allocation

Loops

Network calls

Caching

Concurrency

---

# Testing Review

Verify:

Unit Test

Integration Test

Regression Risk

Negative Cases

Boundary Cases

Test readability

---

# Documentation Review

Verify:

README

API

Migration

Architecture

Configuration

Changelog

---

# Severity Levels

CRITICAL

Must fix.

HIGH

Should fix before merge.

MEDIUM

Recommended.

LOW

Optional improvement.

---

# Approval Rules

APPROVED

No critical findings.

REVISION REQUIRED

One or more findings exist.

REJECTED

Implementation fundamentally incorrect.

---

# Reviewer Behavior

Be objective.

Explain every finding.

Suggest improvements.

Never rewrite code unless requested.

Never reject without explanation.

Every finding must include:

Problem

Reason

Recommendation

Expected outcome

---

# Final Checklist

□ Business objective achieved

□ Engineering standards followed

□ Security maintained

□ Tests adequate

□ Documentation complete

□ No critical issues

Only then:

APPROVED