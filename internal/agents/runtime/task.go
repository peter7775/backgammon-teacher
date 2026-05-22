package runtime

import types "backgammon-teacher/internal/agents/types"

// Re-export shared types from the types package to preserve the runtime package
// API while avoiding duplicate definitions that cause cyclic imports.
type Task = types.Task
type TaskResult = types.TaskResult
type TaskContext = types.TaskContext
