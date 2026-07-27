# Automated Implementation Guide

This directory contains the prescriptive material used by automated coding
agents. Human contributors may inspect it, but the main documentation index is
the human entry point.

Read in this order:

1. [Implementation guide](delivery/IMPLEMENTATION_GUIDE.md)
2. [Delivery roadmap](delivery/ROADMAP.md)
3. [Ticket catalog](tickets/README.md)
4. [Repository quality contract](engineering/REPOSITORY_QUALITY.md)
5. [Dependency register](engineering/DEPENDENCY_REGISTER.md)
6. [Change workflow](engineering/TICKET_BRANCH_WORKFLOW.md)

An agent works on exactly one authorized ticket, uses the linked human product
and architecture specifications, evaluates maintained libraries first, and
submits one squashed commit through a pull request into `dev`.

## Delivery

- [Implementation guide](delivery/IMPLEMENTATION_GUIDE.md)
- [Roadmap](delivery/ROADMAP.md)

## Engineering controls

- [Dependency register](engineering/DEPENDENCY_REGISTER.md)
- [Documentation style](engineering/DOCUMENTATION_STYLE.md)
- [Repository quality](engineering/REPOSITORY_QUALITY.md)
- [Ticket branch workflow](engineering/TICKET_BRANCH_WORKFLOW.md)

These controls are merge-blocking when applicable. Warning-only configuration
does not satisfy a required gate. Detailed adversarial procedures remain in the
operator's ignored local security plan; public requirements define outcomes
without publishing the private test playbook.

## Tickets

- [Catalog and status model](tickets/README.md)
- [Ticket template](tickets/TEMPLATE.md)
