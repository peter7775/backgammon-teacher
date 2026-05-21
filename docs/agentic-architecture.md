# Agentic architecture

## Pattern

Central supervisor + specialized agents + tool registry + memory + reflection.

## Core agents

- analysis-agent
- coach-agent
- lesson-agent
- progress-agent
- localization-agent

## Main flow

1. MoveSubmitted event
2. Supervisor builds plan
3. Analysis agent analyzes the move
4. Coach agent explains the move
5. Progress agent updates profile
6. Lesson agent may recommend next lesson
